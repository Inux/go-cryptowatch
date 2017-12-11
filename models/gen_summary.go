package models

type Summary struct {
	Allowance Summary_sub1 `json:"allowance"`
	Result    Summary_sub4 `json:"result"`
}

type Summary_sub2 struct {
	Absolute   float64 `json:"absolute"`
	Percentage float64 `json:"percentage"`
}

type Summary_sub3 struct {
	Change Summary_sub2 `json:"change"`
	High   int64        `json:"high"`
	Last   float64      `json:"last"`
	Low    float64      `json:"low"`
}

type Summary_sub1 struct {
	Cost      int64 `json:"cost"`
	Remaining int64 `json:"remaining"`
}

type Summary_sub4 struct {
	Price  Summary_sub3 `json:"price"`
	Volume float64      `json:"volume"`
}
