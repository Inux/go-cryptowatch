package models

type Market struct {
	Allowance Market_sub1 `json:"allowance"`
	Result    Market_sub3 `json:"result"`
}

type Market_sub3 struct {
	Active   bool        `json:"active"`
	Exchange string      `json:"exchange"`
	Pair     string      `json:"pair"`
	Routes   Market_sub2 `json:"routes"`
}

type Market_sub1 struct {
	Cost      int64 `json:"cost"`
	Remaining int64 `json:"remaining"`
}

type Market_sub2 struct {
	Ohlc      string `json:"ohlc"`
	Orderbook string `json:"orderbook"`
	Price     string `json:"price"`
	Summary   string `json:"summary"`
	Trades    string `json:"trades"`
}
