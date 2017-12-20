package currency

import "github.com/inux/go-cryptowatch/m/types"

//Currency model
type Currency struct {
	Type types.CurrencyType `json:"type"`
	Name string             `json:"name"`
}

//NewCurrency create new Currency
func NewCurrency(currencyType types.CurrencyType) *Currency {
	currency := &Currency{
		Type: currencyType,
		Name: currencyType.String(),
	}
	return currency
}
