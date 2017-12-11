package models

type Assets struct {
	Allowance Assets_sub1   `json:"allowance"`
	Result    []Assets_sub2 `json:"result"`
}

type Assets_sub1 struct {
	Cost      int64 `json:"cost"`
	Remaining int64 `json:"remaining"`
}

type Assets_sub2 struct {
	Fiat   bool   `json:"fiat"`
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Route  string `json:"route"`
	Symbol string `json:"symbol"`
}
