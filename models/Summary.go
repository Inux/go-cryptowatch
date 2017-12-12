package models

//Summary model
type Summary struct {
	Actual     float64 `json:"actual"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Percentage float64 `json:"percentage"`
	Absolute   float64 `json:"absolute"`
	TimeStamp  int64   `json:"timestamp"`
	Unit       string  `json:"unit"`
}