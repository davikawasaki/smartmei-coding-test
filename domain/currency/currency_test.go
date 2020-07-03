package currency

import (
	dateformat "searchprice/utils/date"
	"strings"
	"testing"
)

type TestDataItem struct {
	price         float64
	base          string
	symbols       []string
	qtyCurrencies int
	hasError      bool
}

func TestConvertToCurrencies(t *testing.T) {
	dataItems := []TestDataItem{
		{12, "", []string{"EUR", "USD"}, 2, false},
		{12, "BRL", []string{"EUR", "USD"}, 2, false},
		{1, "BRL", []string{"EUR"}, 1, false},
		{1, "BRL", []string{"EUR", "EUR"}, 1, false},
		{15000.46846446484646, "BRL", []string{"USD", "EUR"}, 2, false},
		{46846446484646.46846446484646, "BRL", []string{}, 33, false},
		{0, "", []string{"EUR", "USD"}, 0, true},
		{-5000.00440, "BRL", []string{"EUR"}, 0, true},
	}

	for _, item := range dataItems {
		err, currencies, date := ConvertToCurrencies(item.price, item.symbols, item.base)

		var convertedCurrenciesMap []string
		for key, _ := range currencies {
			convertedCurrenciesMap = append(convertedCurrenciesMap, key)
		}

		expectedDate := dateformat.GetNowFormattedDate("")

		if err != nil {
			if item.hasError {
				t.Logf("ConvertToCurrencies() with args (price: %v, symbols: %v, base: %v) PASSED, expected error and got error '%v'", item.price, item.symbols, item.base, err)
			} else {
				t.Errorf("ConvertToCurrencies() with args (price: %v, symbols: %v, base: %v) FAILED, expected total of %v currencies and date %v but got error '%v'", item.price, item.symbols, item.base, item.qtyCurrencies, expectedDate, err)
			}
		} else if len(convertedCurrenciesMap) == item.qtyCurrencies {
			t.Logf("ConvertToCurrencies() with args (price: %v, symbols: %v, base: %v) PASSED, expected output (total currencies: '%v', today date: '%v') and got output (currencies: '%v', latest exchange date: '%v')", item.price, item.symbols, item.base, item.qtyCurrencies, expectedDate, strings.Join(convertedCurrenciesMap, ","), date)
		} else {
			t.Errorf("ConvertToCurrencies() with args (price: %v, symbols: %v, base: %v) FAILED, expected output (total currencies: '%v', today date: '%v') but got output (currencies: '%v', latest exchange date: '%v')", item.price, item.symbols, item.base, item.qtyCurrencies, expectedDate, strings.Join(convertedCurrenciesMap, ","), date)
		}
	}
}
