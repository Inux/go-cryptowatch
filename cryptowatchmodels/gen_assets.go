package cryptowatchmodels

type Assets struct {
	Allowance Assets_sub1   `json:"allowance"`
	Result    []Assets_sub2 `json:"result"`
}

type Assets_sub1 struct {
	Cost      float64 `json:"cost"`
	Remaining float64 `json:"remaining"`
}

type Assets_sub2 struct {
	Fiat   bool   `json:"fiat"`
	ID     float64  `json:"id"`
	Name   string `json:"name"`
	Route  string `json:"route"`
	Symbol string `json:"symbol"`
}
