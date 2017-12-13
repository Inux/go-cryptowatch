package models

import (
	"encoding/json"
	"fmt"
	"go-cryptowatch/common"
	"go-cryptowatch/cryptowatchmodels"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var url = "https://api.cryptowat.ch/"

// TotalRateLimiter RateLimiter of api in nanosecond
const TotalRateLimiter = 8 * 1000 * 1000 * 1000

const averageStatsSize = 20

var averageStatsPos = 0
var averageStats [averageStatsSize]int

// RateLimiter model
type RateLimiter struct {
	mutex             *sync.Mutex
	RemainingCosts    int
	AverageCostsStats [averageStatsSize]int
	costFactor        int
}

// NewRateLimiter - create a new RateLimiter handler
func NewRateLimiter(costFactor int) *RateLimiter {
	resp, err := http.Get(url)
	common.CheckErr(err)
	var index cryptowatchmodels.Index
	err = json.Unmarshal(common.GetBytes(resp.Body), &index)
	common.CheckErr(err)

	rl := &RateLimiter{
		mutex:             &sync.Mutex{},
		RemainingCosts:    int(index.Allowance.Remaining),
		AverageCostsStats: initAverageStats(int(index.Allowance.Cost), costFactor),
		costFactor:        costFactor,
	}

	fmt.Println("New RateLimiter: ")
	fmt.Println("Cost          : " + strconv.Itoa((int(index.Allowance.Cost))))
	fmt.Println("Average       : " + strconv.Itoa(rl.GetAverage()))
	fmt.Println("Remaining Cost: " + strconv.Itoa(rl.RemainingCosts))

	return rl
}

// SetRemainingCosts set RateLimiter.RemainingCosts
func (t *RateLimiter) SetRemainingCosts(remainingCost int) {
	t.mutex.Lock()
	cost := t.RemainingCosts - remainingCost
	if cost > 0 {
		averageStats[averageStatsPos] = cost * t.costFactor
		averageStatsPos++
		if averageStatsPos >= averageStatsSize {
			averageStatsPos = 0
		}
	}
	t.RemainingCosts = remainingCost
	t.mutex.Unlock()
}

// CanSchedule return true if there is enough cost to execeute
func (t *RateLimiter) CanSchedule() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	retval := t.RemainingCosts > t.getAverage()
	return retval
}

// GetNextTimer for next timer so the RateLimiter is never reached
func (t *RateLimiter) GetNextTimer() (*time.Timer, time.Duration) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	dur := caluclateWaitTime(t)
	return time.NewTimer(dur), dur
}

func caluclateWaitTime(t *RateLimiter) time.Duration {
	rc := TotalRateLimiter / t.getAverage()
	return time.Duration(int(time.Hour) / rc)
}

func initAverageStats(costs, costfactor int) [averageStatsSize]int {
	var stats [averageStatsSize]int
	for x := 0; x < len(stats); x++ {
		stats[x] = costs * costfactor
	}
	return stats
}

//GetAverage return the Average Costs
func (t *RateLimiter) GetAverage() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.getAverage()
}

func (t *RateLimiter) getAverage() int {
	var sum = 0
	for _, v := range t.AverageCostsStats {
		sum += v
	}
	return sum / averageStatsSize
}
