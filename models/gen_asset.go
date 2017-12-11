package models

type Asset struct {
	Allowance Asset_sub1 `json:"allowance"`
	Result    Asset_sub4 `json:"result"`
}

type Asset_sub2 struct {
	Active   bool   `json:"active"`
	Exchange string `json:"exchange"`
	ID       int64  `json:"id"`
	Pair     string `json:"pair"`
	Route    string `json:"route"`
}

type Asset_sub3 struct {
	Base  []Asset_sub2 `json:"base"`
	Quote []Asset_sub2 `json:"quote"`
}

type Asset_sub1 struct {
	Cost      int64 `json:"cost"`
	Remaining int64 `json:"remaining"`
}

type Asset_sub4 struct {
	Fiat    bool       `json:"fiat"`
	ID      int64      `json:"id"`
	Markets Asset_sub3 `json:"markets"`
	Name    string     `json:"name"`
	Symbol  string     `json:"symbol"`
}
