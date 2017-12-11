package cryptowatchmodels

type Trades struct {
	Allowance Trades_sub1 `json:"allowance"`
	Result    [][]float64   `json:"result"`
}

type Trades_sub1 struct {
	Cost      float64 `json:"cost"`
	Remaining float64 `json:"remaining"`
}
