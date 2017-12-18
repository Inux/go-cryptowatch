package endpoint

import (
	"go-cryptowatch/m/market"
	"go-cryptowatch/ratelimiter"
)

// Endpoint interface
type Endpoint interface {
	Run(rl *ratelimiter.RateLimiter, market []market.Market)
	Get(m market.Market) interface{}
	GetRange(m market.Market, count int) []interface{}
}
