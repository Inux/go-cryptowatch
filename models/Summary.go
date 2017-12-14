package models

import (
	"encoding/json"
	"go-cryptowatch/common"
	"go-cryptowatch/cryptowatchmodels"
	"net/http"
	"strings"
	"time"
)

//Summary model
type Summary struct {
	Actual       float64 `json:"actual"`
	High         float64 `json:"high"`
	Low          float64 `json:"low"`
	Percentage   float64 `json:"percentage"`
	Absolute     float64 `json:"absolute"`
	TimeStamp    int64   `json:"timestamp"`
	Unit         string  `json:"unit"`
	APICost      float64 `json:"apicost"`
	APIRemaining float64 `json:"apiremaining"`
}

// FetchSummary from single Market
func FetchSummary(market Market) *Summary {
	t := time.Now()
	resp, err := http.Get(market.SummaryURL)
	t = common.GetCorrectedTime(t, time.Now())
	common.CheckErr(err)
	var summary cryptowatchmodels.Summary
	err = json.Unmarshal(common.GetBytes(resp.Body), &summary)
	common.CheckErr(err)

	//convert two own Summary
	return &Summary{
		Actual:       summary.Result.Price.Last,
		High:         summary.Result.Price.High,
		Low:          summary.Result.Price.Low,
		Percentage:   summary.Result.Price.Change.Percentage,
		Absolute:     summary.Result.Price.Change.Absolute,
		TimeStamp:    int64(t.Nanosecond()),
		Unit:         strings.ToLower(market.PairType.String()),
		APICost:      summary.Allowance.Cost,
		APIRemaining: summary.Allowance.Remaining,
	}
}
