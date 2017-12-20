package ratelimiter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/inux/go-cryptowatch/common"
	"github.com/inux/go-cryptowatch/cryptowatchmodels"
)

// totalRateLimiter RateLimiter of api in nanosecond
var totalRateLimiter int

// RateLimiter model
type RateLimiter struct {
	mutex             *sync.Mutex
	RemainingCosts    int
	AverageCostsStats []int
	averageStatsPos   int
	costFactor        int
}

var sigs = make(chan os.Signal, 1)

// NewRateLimiter - create a new RateLimiter handler
func NewRateLimiter(url string, maxCosts, initalCostFactor int) *RateLimiter {
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	totalRateLimiter = maxCosts
	resp, err := http.Get(url)
	common.CheckErr(err)
	var index cryptowatchmodels.Index
	err = json.Unmarshal(common.GetBytes(resp.Body), &index)
	common.CheckErr(err)

	rl := &RateLimiter{
		mutex:             &sync.Mutex{},
		RemainingCosts:    int(index.Allowance.Remaining),
		AverageCostsStats: initAverageStats(int(index.Allowance.Cost), initalCostFactor),
		costFactor:        initalCostFactor,
	}

	fmt.Println("New RateLimiter: ")
	fmt.Println("Cost          : " + strconv.Itoa((int(index.Allowance.Cost))))
	fmt.Println("Average       : " + strconv.Itoa(rl.GetAverage()))
	fmt.Println("Remaining Cost: " + strconv.Itoa(rl.RemainingCosts))

	return rl
}

// Schedule executes Task - channel sends true if success
func (t *RateLimiter) Schedule(task func()) <-chan (bool) {
	c := make(chan (bool), 1)
	executed := false
	var tmr *time.Timer
	if t.canSchedule() {
		task()
		executed = true
		waittime := caluclateWaitTime(t)
		tmr = time.NewTimer(waittime)
		fmt.Printf("ratelimiter.Schedule: waiting: %.6fs\n", waittime.Seconds())
	} else {
		waittime := time.Hour - time.Duration(time.Now().Minute())
		tmr = time.NewTimer(waittime)
		fmt.Printf("ratelimiter.Schedule: out of costs, waiting: %.6fs\n", waittime.Seconds())
	}
	go func() {
		select {
		case <-tmr.C:
			if !executed {
				task()
			}
			c <- true
			break
		case <-sigs:
			c <- false
			break
		}
	}()
	return c
}

// SetRemainingCosts set RateLimiter.RemainingCosts
func (t *RateLimiter) SetRemainingCosts(remainingCost int) {
	t.mutex.Lock()
	cost := t.RemainingCosts - remainingCost
	if cost > 0 {
		if t.averageStatsPos < len(t.AverageCostsStats) {
			t.AverageCostsStats[t.averageStatsPos] = cost
			t.averageStatsPos++
		} else {
			t.averageStatsPos = 0
			t.AverageCostsStats[t.averageStatsPos] = cost
		}
	}
	t.RemainingCosts = remainingCost
	t.mutex.Unlock()
}

// GetAverage gets average cost
func (t *RateLimiter) GetAverage() int {
	return getAverage(t.AverageCostsStats)
}

// canSchedule return true if there is enough cost to execeute
func (t *RateLimiter) canSchedule() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	retval := t.RemainingCosts > getAverage(t.AverageCostsStats)
	return retval
}

// getNextTimer for next timer so the RateLimiter is never reached
func (t *RateLimiter) getNextTimer() (*time.Timer, time.Duration) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	dur := caluclateWaitTime(t)
	return time.NewTimer(dur), dur
}

func caluclateWaitTime(t *RateLimiter) time.Duration {
	rc := totalRateLimiter / getAverage(t.AverageCostsStats)
	return time.Duration(int(time.Hour) / rc)
}

func initAverageStats(costs, costfactor int) []int {
	stats := make([]int, costfactor)
	for x := 0; x < len(stats); x++ {
		stats[x] = costs
	}
	return stats
}

func getAverage(values []int) int {
	var sum = 0
	for _, v := range values {
		sum += v
	}
	return sum
}
