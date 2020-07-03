package crawler

import (
	"errors"
	"fmt"
	"net/http"
	stringformat "searchprice/utils/string"
	"time"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
)

func ExtractProfessionalPrice(url string, class string) (error, string) {
	if url == "" {
		url = "https://www.smartmei.com.br/"
	}

	if class == "" {
		class = "tarifas-2-2-2"
	}

	// Make HTTP GET request
	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		fmt.Printf("%v\n", stringformat.LogMessage(err.Error(), "ERROR"))
		return err, ""
	}
	defer res.Body.Close()

	// Parse the page
	root, err := html.Parse(res.Body)
	if err != nil {
		fmt.Printf("%v\n", stringformat.LogMessage(err.Error(), "ERROR"))
		return err, ""
	}

	// Search for value
	value, hasMatch := scrape.Find(root, scrape.ByClass(class))
	if hasMatch {
		return nil, scrape.Text(value)
	}

	errMsg := fmt.Sprintf("No element class '%v' found on URL '%v'!", class, url)
	fmt.Printf("%v\n", stringformat.LogMessage(errMsg, "ERROR"))
	return errors.New(errMsg), ""
}
