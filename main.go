package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/inux/go-cryptowatch/endpoint"
	"github.com/inux/go-cryptowatch/m/market"
	"github.com/inux/go-cryptowatch/m/types"
	"github.com/inux/go-cryptowatch/ratelimiter"
)

var markets = [...]market.Market{
	//ETH
	*market.NewMarket(types.Bittrex, types.ETH, types.Ethusd),
	//IOTA
	*market.NewMarket(types.Bittrex, types.IOT, types.Iotusd),
	//Litecoin
	*market.NewMarket(types.Bitfinex, types.LTC, types.Ltcusd),
	//Monero
	*market.NewMarket(types.Bitfinex, types.XMR, types.Xmrusd),
	//NEO
	*market.NewMarket(types.Bitfinex, types.NEO, types.Neousd),
	//XRP
	*market.NewMarket(types.Bitfinex, types.XRP, types.Xrpusd),
}

var ep = []endpoint.Endpoint{
	&endpoint.SummaryEndpoint{},
}

var rlURL = "https://api.cryptowat.ch/"
var rlMaxCost = int(8 * time.Second)
var rl = ratelimiter.NewRateLimiter(rlURL, rlMaxCost, len(markets))

var sigs = make(chan os.Signal, 1)

func main() {
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ma := make([]market.Market, len(markets))
	for _, v := range markets {
		ma = append(ma, v)
	}

	for _, e := range ep {
		e.Run(rl, ma)
	}

	<-sigs
}
