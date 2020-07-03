package currency

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	stringformat "searchprice/utils/string"
	"strings"
)

type CurrencyExchangeRates struct {
	Rates map[string]float64
	Base  string
	Date  string
}

func ConvertToCurrencies(price float64, currencies []string, base string) (error, map[string]float64, string) {
	if price <= 0 {
		errMsg := fmt.Sprintf("Set a valid price to be converted instead of '%v'!", price)
		fmt.Printf("%v\n", stringformat.LogMessage(errMsg, "ERROR"))
		return errors.New(errMsg), nil, ""
	}

	if base == "" {
		base = "BRL"
	}

	baseURL := "https://api.exchangeratesapi.io/latest"
	argsURL := fmt.Sprintf("%s?base=%s", baseURL, base)

	if len(currencies) > 0 {
		argsURL = fmt.Sprintf("%s&symbols=%s", argsURL, strings.Join(currencies, ","))
	}

	// Make HTTP GET request
	res, err := http.Get(argsURL)
	if err != nil {
		fmt.Printf("%v\n", stringformat.LogMessage(err.Error(), "ERROR"))
		return err, nil, ""
	}
	defer res.Body.Close()

	// Parse the object
	var cer CurrencyExchangeRates
	err = json.NewDecoder(res.Body).Decode(&cer)
	if err != nil {
		fmt.Printf("%v\n", stringformat.LogMessage(err.Error(), "ERROR"))
		return err, nil, ""
	}

	return nil, cer.Rates, cer.Date
}
