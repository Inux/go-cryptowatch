package cryptowatchmodels

type Markets struct {
	Allowance Markets_sub1   `json:"allowance"`
	Result    []Markets_sub2 `json:"result"`
}

type Markets_sub2 struct {
	Active   bool   `json:"active"`
	Exchange string `json:"exchange"`
	ID       float64  `json:"id"`
	Pair     string `json:"pair"`
	Route    string `json:"route"`
}

type Markets_sub1 struct {
	Cost      float64 `json:"cost"`
	Remaining float64 `json:"remaining"`
}
