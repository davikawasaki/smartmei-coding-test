package crawler

import (
	"testing"
)

type TestDataItem struct {
	url      string
	class    string
	output   string
	hasError bool
}

func TestExtractProfessionalPrice(t *testing.T) {
	dataItems := []TestDataItem{
		{"", "tarifas-2-1-2", "R$ 5,00 por boleto pago.", false},
		{"https://www.smartmei.com.br/", "tarifas-2-1-2", "R$ 5,00 por boleto pago.", false},
		{"https://www.smartmei.com.br/", "tarifas-2-2-2", "R$ 7,00", false},
		{"https://www.smartmei.com.br/", "", "R$ 7,00", false},
		{"https://www.smartmei.com.br/", "tarifas-2-3-2", "R$ 15,00 *pagando R$45,00/trimestre", false},
		{"https://www.smartmei.com.br/", "tarifaxs-2-3-2", "", true},
	}

	for _, item := range dataItems {
		err, result := ExtractProfessionalPrice(item.url, item.class)

		if err != nil {
			if item.hasError {
				t.Logf("ExtractProfessionalPrice() with args %v, %v PASSED, expected error and got error '%v'", item.url, item.class, err)
			} else {
				t.Errorf("ExtractProfessionalPrice() with args %v, %v FAILED, expected %v but got error '%v'", item.url, item.class, item.output, err)
			}
		} else if result == item.output {
			t.Logf("ExtractProfessionalPrice() with args %v, %v PASSED, expected output '%v' and got output '%v'", item.url, item.class, item.output, result)
		} else {
			t.Errorf("ExtractProfessionalPrice() with args %v, %v FAILED, expected output '%v' but got output '%v'", item.url, item.class, item.output, result)
		}
	}
}
