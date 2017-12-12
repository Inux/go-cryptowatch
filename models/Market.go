package models

import "strings"

const (
	baseurlSummary = "https://api.cryptowat.ch/markets" //"market"/"pairname"/summary
)

//SummaryKey used for Summaries access
type SummaryKey string

//Market model
type Market struct {
	Name       string   `json:"name"`
	SummaryURL string   `json:"summaryurl"`
	PairType   PairType `json:"pairtype"`
}

//GetSummaryKey returns the key to access Summaries of Currency
func (t *Market) GetSummaryKey() SummaryKey {
	return SummaryKey(t.Name + t.PairType.String())
}

//NewMarket creates a new Market
func NewMarket(Name string, PairType PairType) *Market {
	return &Market{
		Name:       Name,
		SummaryURL: getSummaryURL(Name, PairType),
		PairType:   PairType,
	}
}

func getSummaryURL(Name string, PairType PairType) string {
	return baseurlSummary + "/" + strings.ToLower(Name) + "/" + strings.ToLower(PairType.String()) + "/summary"
}

//PairType type
type PairType int

const (
	Ethbtc PairType = iota
	Ethusd
	Miotabtc
	Miotausd
	Ltcbtc
	Ltcusd
	Xmrbtc
	Xmrusd
	Xembtc
	Xemusd
	Adabtc
	Adausd
	Neobtc
	Neousd
	Nxtbtc
	Nxtusd
	Xlmbtc
	Xlmusd
	Xrpbtc
	Xrpusd
)
