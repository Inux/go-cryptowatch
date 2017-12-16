package main

import (
	"fmt"
	"go-cryptowatch/common"
	"go-cryptowatch/models"
	"strconv"
	"sync"
	"time"
)

var URLratelimiter = "https://api.cryptowat.ch/"

var done = make(chan bool, 1)

var genSummaryPool = &sync.WaitGroup{}
var addSummariesPool = &sync.WaitGroup{}

var ratelimiter = common.NewRateLimiter(URLratelimiter, getCostFactor())

func main() {
	scheduledSummaries()
}

func scheduledSummaries() {
	c := ratelimiter.Schedule(genSummariesTotal)
	select {
	case msg := <-c:
		if msg {
			scheduledSummaries()
		} else {
			break
		}
	}
}

func genSummariesTotal() {
	genSummaries()
	printSummaries()
}

func genSummaries() {
	for _, c := range currencies {
		go addSummaries(c)
		genSummaryPool.Add(1)
	}
	genSummaryPool.Wait()
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
	currenciesMutex.Lock()
	summary := models.FetchSummary(market)

	currency.Summaries[market.GetSummaryKey()] = append(currency.Summaries[market.GetSummaryKey()], *summary)
	ratelimiter.SetRemainingCosts(int(summary.APIRemaining))
	currenciesMutex.Unlock()
	addSummariesPool.Done()
}

func printSummaries() {
	for _, c := range currencies {
		fmt.Println("Currency: " + c.CurrencyType.String())
		for i, m := range c.Markets {
			fmt.Println("   Market " + strconv.Itoa(i) + " - " + m.Name + " / " + m.PairType.String() + " :")
			for si, s := range c.Summaries[m.GetSummaryKey()] {
				fmt.Println("Summary " + strconv.Itoa(si) + ":")
				fmt.Printf("      Actual: %.6f, High: %.6f, Low: %.6f\n", s.Actual, s.High, s.Low)
				fmt.Printf("      Growth %%: %.6f, Growth $: %.6f\n", s.Percentage, s.Absolute)
				fmt.Println("      Time: " + time.Now().Truncate(time.Duration(s.TimeStamp)).String() + ", Unit: " + s.Unit)
				fmt.Printf("      Costs [s]: %.6fs, Remaining [s]: %.6fs\n", s.APICost/(1000*1000*1000), s.APIRemaining/(1000*1000*1000))
			}
		}
		fmt.Println()
	}
	fmt.Println("Time: " + time.Now().String())
	fmt.Printf("Remaining Costs [s]: %.6f\n", time.Duration(ratelimiter.RemainingCosts).Seconds())
	fmt.Printf("Average Costs [s]  : %.6f\n", time.Duration(ratelimiter.GetAverage()).Seconds())
	fmt.Println("-------------------------------------------------------")
}
