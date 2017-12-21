package summary

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/inux/go-cryptowatch/common"
	"github.com/inux/go-cryptowatch/cryptowatchmodels"
	"github.com/inux/go-cryptowatch/m/types"
)

const (
	baseurlSummary = "https://api.cryptowat.ch/markets" //"market"/"pairname"/summary
)

//Summary model
type Summary struct {
	types.CurrencyType `json:"currencytype"`
	types.MarketType   `json:"markettype"`
	types.PairType     `json:"pairtype"`

	CurrencyName string  `json:"currencyname"`
	MarketeName  string  `json:"marketname"`
	PairName     string  `json:"pairname"`
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
func FetchSummary(ct types.CurrencyType, mt types.MarketType, pt types.PairType) *Summary {
	defer handleRecover()
	t := time.Now()
	resp, err := http.Get(getSummaryURL(mt, pt))
	t = common.GetCorrectedTime(t, time.Now())
	common.CheckErr(err)
	var summary cryptowatchmodels.Summary
	err = json.Unmarshal(common.GetBytes(resp.Body), &summary)
	common.CheckErr(err)

	//convert two own Summary
	return &Summary{
		CurrencyType: ct,
		MarketType:   mt,
		PairType:     pt,
		CurrencyName: ct.String(),
		MarketeName:  mt.String(),
		PairName:     pt.String(),
		Actual:       summary.Result.Price.Last,
		High:         summary.Result.Price.High,
		Low:          summary.Result.Price.Low,
		Percentage:   summary.Result.Price.Change.Percentage,
		Absolute:     summary.Result.Price.Change.Absolute,
		TimeStamp:    int64(t.Nanosecond()),
		Unit:         strings.ToLower(pt.String()),
		APICost:      summary.Allowance.Cost,
		APIRemaining: summary.Allowance.Remaining,
	}
}

func getSummaryURL(mt types.MarketType, pt types.PairType) string {
	return baseurlSummary + "/" + strings.ToLower(mt.String()) + "/" + strings.ToLower(pt.String()) + "/summary"
}

func handleRecover() *Summary {
	return nil
}
