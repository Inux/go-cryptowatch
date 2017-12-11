package models

type Markets struct {
	Allowance Markets_sub1   `json:"allowance"`
	Result    []Markets_sub2 `json:"result"`
}

type Markets_sub2 struct {
	Active   bool   `json:"active"`
	Exchange string `json:"exchange"`
	ID       int64  `json:"id"`
	Pair     string `json:"pair"`
	Route    string `json:"route"`
}

type Markets_sub1 struct {
	Cost      int64 `json:"cost"`
	Remaining int64 `json:"remaining"`
}
