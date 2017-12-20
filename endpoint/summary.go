package endpoint

import (
	"fmt"
	"sync"
	"time"

	"github.com/inux/go-cryptowatch/m/market"
	"github.com/inux/go-cryptowatch/m/summary"
	"github.com/inux/go-cryptowatch/m/types"
	"github.com/inux/go-cryptowatch/ratelimiter"
)

// SummaryEndpoint summaryendpoint
type SummaryEndpoint struct {
	sync.Mutex
	rl        *ratelimiter.RateLimiter
	markets   []market.Market
	summaries map[types.CurrencyType]map[types.MarketType]map[types.PairType][]summary.Summary
	gopool    sync.WaitGroup
}

// Run is starting Endpoint
func (t *SummaryEndpoint) Run(rl *ratelimiter.RateLimiter, markets []market.Market) {
	t.rl = rl
	t.markets = markets
	t.summaries = initSummaries(markets)
	go t.scheduledSummaries()
}

// Get gets last summary
func (t *SummaryEndpoint) Get(m market.Market) interface{} {
	return t.GetRange(m, 1)[0]
}

// GetRange returns range of summaries
func (t *SummaryEndpoint) GetRange(m market.Market, count int) []interface{} {
	t.Lock()
	defer func() {
		recover()
		t.Unlock()
	}()
	length := len(t.summaries[m.CurrencyType][m.Type][m.PairType])
	if length < count {
		return getSlice(t.summaries[m.CurrencyType][m.Type][m.PairType][:length])
	}
	return getSlice(t.summaries[m.CurrencyType][m.Type][m.PairType][:count])
}

// started in Run
func (t *SummaryEndpoint) scheduledSummaries() {
	c := t.rl.Schedule(t.fetchSummariesTask)
	select {
	case msg := <-c:
		if msg {
			t.scheduledSummaries()
		} else {
			break
		}
	}
}

func (t *SummaryEndpoint) fetchSummariesTask() {
	for _, m := range t.markets {
		go t.addSummary(m)
		t.gopool.Add(1)
	}
	t.gopool.Wait()
}

// async
func (t *SummaryEndpoint) addSummary(m market.Market) {
	t.Lock()
	summary := summary.FetchSummary(m.CurrencyType, m.Type, m.PairType)
	if summary != nil {
		t.summaries[m.CurrencyType][m.Type][m.PairType] =
			append(t.summaries[m.CurrencyType][m.Type][m.PairType], *summary)
		t.rl.SetRemainingCosts(int(summary.APIRemaining))
		printSummary(*summary)
	}
	t.Unlock()
	t.gopool.Done()
}

func getSlice(summaries []summary.Summary) []interface{} {
	var is = make([]interface{}, len(summaries))
	for i, d := range summaries {
		is[i] = d
	}
	return is
}

func initSummaries(markets []market.Market) map[types.CurrencyType]map[types.MarketType]map[types.PairType][]summary.Summary {
	m := make(map[types.CurrencyType]map[types.MarketType]map[types.PairType][]summary.Summary)
	for _, v := range markets {
		if m[v.CurrencyType] == nil {
			m[v.CurrencyType] = make(map[types.MarketType]map[types.PairType][]summary.Summary)
		}
		if m[v.CurrencyType][v.Type] == nil {
			m[v.CurrencyType][v.Type] = make(map[types.PairType][]summary.Summary)
		}
		if m[v.CurrencyType][v.Type][v.PairType] == nil {
			m[v.CurrencyType][v.Type][v.PairType] = make([]summary.Summary, 100)
		}
	}
	return m
}

func printSummary(s summary.Summary) {
	fmt.Println("Currency: " + s.CurrencyType.String())
	fmt.Println("Market  : " + s.MarketType.String() + " / " + s.PairType.String() + " :")
	fmt.Printf("Actual  : %.6f, High: %.6f, Low: %.6f\n", s.Actual, s.High, s.Low)
	fmt.Printf("Growth %%: %.6f, Growth $: %.6f\n", s.Percentage, s.Absolute)
	fmt.Printf("Time    : %.6fs\n", time.Duration(s.TimeStamp).Seconds())
	fmt.Printf("Remaining Costs [s]: %.6f\n", time.Duration(int(s.APIRemaining)).Seconds())
	fmt.Printf("Costs [s]  : %.6f\n", time.Duration(int(s.APICost)).Seconds())
	fmt.Println("-------------------------------------------------------")
}
