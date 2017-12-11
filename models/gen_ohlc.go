package models

type OHLC struct {
	Allowance OHLC_sub1 `json:"allowance"`
	Result    OHLC_sub2 `json:"result"`
}

type OHLC_sub1 struct {
	Cost      float64 `json:"cost"`
	Remaining float64 `json:"remaining"`
}

type OHLC_sub2 struct {
	One4400   [][]float64 `json:"14400"`
	One80     [][]float64 `json:"180"`
	One800    [][]float64 `json:"1800"`
	Two1600   [][]float64 `json:"21600"`
	Two59200  [][]float64 `json:"259200"`
	Three00   [][]float64 `json:"300"`
	Three600  [][]float64 `json:"3600"`
	Four3200  [][]float64 `json:"43200"`
	Six0      [][]float64 `json:"60"`
	Six04800  [][]float64 `json:"604800"`
	Seven200  [][]float64 `json:"7200"`
	Eight6400 [][]float64 `json:"86400"`
	Nine00    [][]float64 `json:"900"`
}
