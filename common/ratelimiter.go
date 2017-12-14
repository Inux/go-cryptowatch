package common

import (
	"encoding/json"
	"fmt"
	"go-cryptowatch/cryptowatchmodels"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var url = "https://api.cryptowat.ch/"

// TotalRateLimiter RateLimiter of api in nanosecond
const TotalRateLimiter = 8 * 1000 * 1000 * 1000

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
func NewRateLimiter(initalCostFactor int) *RateLimiter {
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	resp, err := http.Get(url)
	CheckErr(err)
	var index cryptowatchmodels.Index
	err = json.Unmarshal(GetBytes(resp.Body), &index)
	CheckErr(err)

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
	if t.CanSchedule() {
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

func initAverageStats(costs, costfactor int) []int {
	stats := make([]int, costfactor)
	for x := 0; x < len(stats); x++ {
		stats[x] = costs
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
	return sum
}
