package main

import (
	"encoding/json"
	"fmt"
	"go-cryptowatch/common"
	"go-cryptowatch/cryptowatchmodels"
	"go-cryptowatch/models"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

var currenciesMutex = &sync.Mutex{}

// currencies - Use currenciesMutex to modify
var currencies = [...]models.Currency{
	//Etherum
	*models.NewCurrency(
		models.ETH,
		[]models.Market{
			*models.NewMarket("Bitfinex", models.Ethusd),
			*models.NewMarket("Bitfinex", models.Ethbtc),
			*models.NewMarket("Bittrex", models.Ethbtc),
		},
	),
	//IOTA
	*models.NewCurrency(
		models.IOT,
		[]models.Market{
			*models.NewMarket("Bitfinex", models.Iotusd),
			*models.NewMarket("Bitfinex", models.Iotbtc),
			*models.NewMarket("Bittrex", models.Iotusd),
			*models.NewMarket("Bittrex", models.Iotbtc),
		},
	),
	//Litecoin
	*models.NewCurrency(
		models.LTC,
		[]models.Market{
			*models.NewMarket("Bitfinex", models.Ltcusd),
			*models.NewMarket("Bitfinex", models.Ltcbtc),
		},
	),
	//Monero
	*models.NewCurrency(
		models.XMR,
		[]models.Market{
			*models.NewMarket("Bitfinex", models.Xmrusd),
			*models.NewMarket("Bitfinex", models.Xmrbtc),
		},
	),
	//NEO
	*models.NewCurrency(
		models.NEO,
		[]models.Market{
			*models.NewMarket("Bitfinex", models.Neousd),
			*models.NewMarket("Bitfinex", models.Neobtc),
			*models.NewMarket("Bittrex", models.Neobtc),
		},
	),
	//XRP
	*models.NewCurrency(
		models.XRP,
		[]models.Market{
			*models.NewMarket("Bitfinex", models.Xrpusd),
			*models.NewMarket("Bitfinex", models.Xrpbtc),
		},
	),
}

var sigs = make(chan os.Signal, 1)
var done = make(chan bool, 1)

var genSummaryPool = &sync.WaitGroup{}
var addSummariesPool = &sync.WaitGroup{}

var ratelimiter = models.NewRateLimiter(getCostFactor())

func main() {
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	scheduledSummaries()
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

func scheduledSummaries() {
	if ratelimiter.CanSchedule() {
		generateSummaries()
		printSummaries()

		//fakeTask()
	}
	timer, dur := ratelimiter.GetNextTimer()
	fmt.Printf("Timer duration: %.6f\n", dur.Seconds())
	select {
	case <-timer.C:
		break
	case <-sigs:
		os.Exit(0)
	}
	scheduledSummaries()
}

func fakeTask() {
	fmt.Println("Start fake task, time: " + time.Now().String())
	ratelimiter.SetRemainingCosts(ratelimiter.RemainingCosts - ratelimiter.GetAverage()/getCostFactor())
	fmt.Println("AverageCosts: " + strconv.Itoa(ratelimiter.GetAverage()))
	fmt.Println("RemainingCost: " + strconv.Itoa(ratelimiter.RemainingCosts))
}

func generateSummaries() {
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
	summary := fetchSummary(market)

	currency.Summaries[market.GetSummaryKey()] = append(currency.Summaries[market.GetSummaryKey()], *summary)
	ratelimiter.SetRemainingCosts(int(summary.APIRemaining))
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

func getCostFactor() int {
	costfactor := 0
	for _, c := range currencies {
		costfactor += len(c.Markets)
	}
	return costfactor
}
