package models

//Currency model
type Currency struct {
	CurrencyType CurrencyType             `json:"type"`
	Name         string                   `json:"name"`
	Markets      []Market                 `json:"markets"`
	Summaries    map[SummaryKey][]Summary `json:"summaries"`
}

//NewCurrency create new Currency
func NewCurrency(currencyType CurrencyType, markets []Market) *Currency {
	currency := &Currency{
		CurrencyType: currencyType,
		Name:         currencyType.String(),
		Markets:      markets,
	}
	currency.Summaries = make(map[SummaryKey][]Summary)
	for _, m := range markets {
		currency.Summaries[m.GetSummaryKey()] = make([]Summary, 0)
	}
	return currency
}

//CurrencyType type
type CurrencyType int

const (
	ETH CurrencyType = iota
	MIOTA
	LTC
	XMR
	XEM
	ADA
	NEO
	NXT
	XLM
	XRP
)
