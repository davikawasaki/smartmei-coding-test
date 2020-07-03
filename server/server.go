package main

import (
	"fmt"
	"net/http"
	"searchprice/domain/crawler"
	"searchprice/domain/currency"
	stringformat "searchprice/utils/string"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Rates struct {
	Date            string  `json:"date"`
	RateDescription string  `json:"description"`
	PriceBRL        float64 `json:"priceBRL"`
	PriceUSD        float64 `json:"priceUSD"`
	PriceEUR        float64 `json:"priceEUR"`
}

var ratesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Rates",
	Fields: graphql.Fields{
		"date": &graphql.Field{
			Type: graphql.String,
		},
		"rateDescription": &graphql.Field{
			Type: graphql.String,
		},
		"priceBRL": &graphql.Field{
			Type: graphql.Float,
		},
		"priceUSD": &graphql.Field{
			Type: graphql.Float,
		},
		"priceEUR": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var ratesQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProfessionalRatesQuery",
	Fields: graphql.Fields{
		"rate": &graphql.Field{
			Type:        ratesType,
			Description: "Get professional rates",
			Args: graphql.FieldConfigArgument{
				"url": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				url, _ := params.Args["url"].(string)

				err, rateDescription := crawler.ExtractProfessionalPrice(url, "")
				if err != nil {
					return nil, err
				}

				var price float64

				rateDescriptionStripped := stringformat.ReplaceStringSubvals(rateDescription, map[string]string{" ": "", ",": ".", "R$": ""})
				price, err = strconv.ParseFloat(rateDescriptionStripped, 32)
				if err != nil {
					return nil, err
				}

				err, currencies, date := currency.ConvertToCurrencies(price, []string{"EUR", "USD"}, "BRL")
				if err != nil {
					return nil, err
				}

				return Rates{
					Date:            date,
					RateDescription: rateDescription,
					PriceBRL:        price,
					PriceUSD:        currencies["USD"] * price,
					PriceEUR:        currencies["EUR"] * price,
				}, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: ratesQueries,
})

var h = handler.New(&handler.Config{
	Schema: &schema,
	Pretty: true,
})

func main() {

	// Static file server to serve Playground-style Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Use GraphQL handler for this route
	http.Handle("/graphql", h)

	fmt.Println("Playground server has started at: http://localhost:8080/")
	fmt.Println("GraphQL API server has started at: http://localhost:8080/graphql")
	fmt.Println("[EXAMPLE] Get rate: curl -g 'http://localhost:8080/graphql?query={rate(url:%22https://www.smartmei.com.br%22){date,%20rateDescription,priceBRL,priceEUR,priceUSD}}'")
	http.ListenAndServe(":8080", nil)
}
