package cryptowatchmodels

type Exchanges struct {
	Allowance Exchanges_sub1   `json:"allowance"`
	Result    []Exchanges_sub2 `json:"result"`
}

type Exchanges_sub2 struct {
	Active bool   `json:"active"`
	Name   string `json:"name"`
	Route  string `json:"route"`
	Symbol string `json:"symbol"`
}

type Exchanges_sub1 struct {
	Cost      float64 `json:"cost"`
	Remaining float64 `json:"remaining"`
}
