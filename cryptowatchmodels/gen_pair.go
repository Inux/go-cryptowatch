package cryptowatchmodels

type Pair struct {
	Allowance Pair_sub1 `json:"allowance"`
	Result    Pair_sub4 `json:"result"`
}

type Pair_sub3 struct {
	Active   bool   `json:"active"`
	Exchange string `json:"exchange"`
	ID       float64  `json:"id"`
	Pair     string `json:"pair"`
	Route    string `json:"route"`
}

type Pair_sub4 struct {
	Base    Pair_sub2   `json:"base"`
	ID      float64       `json:"id"`
	Markets []Pair_sub3 `json:"markets"`
	Quote   Pair_sub2   `json:"quote"`
	Route   string      `json:"route"`
	Symbol  string      `json:"symbol"`
}

type Pair_sub1 struct {
	Cost      float64 `json:"cost"`
	Remaining float64 `json:"remaining"`
}

type Pair_sub2 struct {
	Fiat   bool   `json:"fiat"`
	ID     float64  `json:"id"`
	Name   string `json:"name"`
	Route  string `json:"route"`
	Symbol string `json:"symbol"`
}
