package cryptowatchmodels

type Price struct {
	Allowance Price_sub1 `json:"allowance"`
	Result    Price_sub2 `json:"result"`
}

type Price_sub1 struct {
	Cost      float64 `json:"cost"`
	Remaining float64 `json:"remaining"`
}

type Price_sub2 struct {
	Price float64 `json:"price"`
}
