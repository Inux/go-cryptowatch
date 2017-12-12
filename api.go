package main

import (
	"encoding/json"
	"fmt"
	"go-cryptowatch/common"
	"go-cryptowatch/cryptowatchmodels"
	"go-cryptowatch/models"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var currenciesMutex = &sync.Mutex{}
var genSummaryPool = &sync.WaitGroup{}
var addSummariesPool = &sync.WaitGroup{}

var currencies = [...]models.Currency{
	*models.NewCurrency(
		models.ETH,
		[]models.Market{
			*models.NewMarket("Bitfinex", models.Ethusd),
			*models.NewMarket("Bitfinex", models.Ethbtc),
			*models.NewMarket("Bittrex", models.Ethbtc),
		},
	),
}

func main() {
	generateSummaries()
	printSummaries()
}

func printSummaries() {
	for _, c := range currencies {
		fmt.Println("Currency: " + c.CurrencyType.String())
		for i, m := range c.Markets {
			fmt.Println("   Market " + strconv.Itoa(i) + " - " + string(m.Name) + " / " + m.PairType.String() + " :")
			for si, s := range c.Summaries[m.GetSummaryKey()] {
				fmt.Println("      Summary " + strconv.Itoa(si) + ":")
				fmt.Printf("      Actual: %.6f, High: %.6f, Low: %.6f\n", s.Actual, s.High, s.Low)
				fmt.Printf("      Growth %%: %.6f, Growth $: %.6f\n", s.Percentage, s.Absolute)
				fmt.Println("      Time: " + time.Now().Truncate(time.Duration(s.TimeStamp)).String() + ", Unit: " + s.Unit)
			}
		}
		fmt.Println()
	}
}

func generateSummaries() int64 {
	start := time.Now()
	for _, c := range currencies {
		go addSummaries(c)
		genSummaryPool.Add(1)
	}
	genSummaryPool.Wait()
	return time.Since(start).Nanoseconds()
}

func addSummaries(currency models.Currency) {
	for _, v := range currency.Markets {
		go addSummary(currency, v)
		addSummariesPool.Add(1)
	}
	addSummariesPool.Wait()
	genSummaryPool.Done()
}

func addSummary(currency models.Currency, market models.Market) {
	summary := fetchSummary(market)
	currenciesMutex.Lock()
	currency.Summaries[market.GetSummaryKey()] = append(currency.Summaries[market.GetSummaryKey()], *summary)
	currenciesMutex.Unlock()
	addSummariesPool.Done()
}

func fetchSummary(market models.Market) *models.Summary {
	t := time.Now()
	resp, err := http.Get(market.SummaryURL)
	t = common.GetCorrectedTime(t, time.Now())
	common.CheckErr(err)
	var summary cryptowatchmodels.Summary
	err = json.Unmarshal(common.GetBytes(resp.Body), &summary)
	common.CheckErr(err)

	//convert two own Summary
	return &models.Summary{
		Actual:     summary.Result.Price.Last,
		High:       summary.Result.Price.High,
		Low:        summary.Result.Price.Low,
		Percentage: summary.Result.Price.Change.Percentage,
		Absolute:   summary.Result.Price.Change.Absolute,
		TimeStamp:  int64(t.Nanosecond()),
		Unit:       strings.ToLower(market.PairType.String()),
	}
}
