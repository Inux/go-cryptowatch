package cryptowatchmodels

type Pairs struct {
	Allowance Pairs_sub1   `json:"allowance"`
	Result    []Pairs_sub3 `json:"result"`
}

type Pairs_sub3 struct {
	Base                  Pairs_sub2 `json:"base"`
	FuturesContractPeriod string     `json:"futuresContractPeriod"`
	ID                    float64      `json:"id"`
	Quote                 Pairs_sub2 `json:"quote"`
	Route                 string     `json:"route"`
	Symbol                string     `json:"symbol"`
}

type Pairs_sub1 struct {
	Cost      float64 `json:"cost"`
	Remaining float64 `json:"remaining"`
}

type Pairs_sub2 struct {
	Fiat   bool   `json:"fiat"`
	ID     float64  `json:"id"`
	Name   string `json:"name"`
	Route  string `json:"route"`
	Symbol string `json:"symbol"`
}
