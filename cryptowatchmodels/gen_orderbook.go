package cryptowatchmodels

type OrderBook struct {
	Allowance OrderBook_sub1 `json:"allowance"`
	Result    OrderBook_sub2 `json:"result"`
}

type OrderBook_sub2 struct {
	Asks [][]float64   `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

type OrderBook_sub1 struct {
	Cost      float64 `json:"cost"`
	Remaining float64 `json:"remaining"`
}
