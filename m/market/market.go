package market

import "go-cryptowatch/m/types"

//Market model
type Market struct {
	Type         types.MarketType   `json:"type"`
	Name         string             `json:"name"`
	PairType     types.PairType     `json:"pairtype"`
	PairName     string             `json:"pairname"`
	CurrencyType types.CurrencyType `json:"currencytype"`
	CurrencyName string             `json:"currencyname"`
}

//NewMarket creates a new Market
func NewMarket(marketType types.MarketType, currencyType types.CurrencyType, pairType types.PairType) *Market {
	return &Market{
		Type:         marketType,
		Name:         marketType.String(),
		PairType:     pairType,
		PairName:     pairType.String(),
		CurrencyType: currencyType,
		CurrencyName: currencyType.String(),
	}
}
