package models

type Exchange struct {
	Allowance Exchange_sub1 `json:"allowance"`
	Result    Exchange_sub3 `json:"result"`
}

type Exchange_sub3 struct {
	Active bool          `json:"active"`
	Name   string        `json:"name"`
	Routes Exchange_sub2 `json:"routes"`
	Symbol string        `json:"symbol"`
}

type Exchange_sub1 struct {
	Cost      int64 `json:"cost"`
	Remaining int64 `json:"remaining"`
}

type Exchange_sub2 struct {
	Markets string `json:"markets"`
}
