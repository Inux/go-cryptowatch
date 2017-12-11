package models

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var out io.Writer = os.Stdout

func TestAssets(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/assets")
	checkErr(err)
	var assets Assets
	err = json.Unmarshal(getBytes(resp.Body), &assets)
	checkErr(err)
	assert.True(t, len(assets.Result) > 0, "assets.Result is empty")
	for _, v := range assets.Result {
		assert.True(t, v.Symbol != "", "Symbol is empty")
		assert.True(t, v.Name != "", "Name is empty")
		assert.True(t, v.Fiat == false || v.Fiat == true, "Fiat is empty")
		assert.True(t, v.Route != "", "Route is empty")
	}
}

func TestAsset(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/assets/btc")
	checkErr(err)
	var asset Asset
	err = json.Unmarshal(getBytes(resp.Body), &asset)
	checkErr(err)
	assert.True(t, asset.Result.ID > 0, "ID is empty")
	assert.True(t, asset.Result.Symbol != "", "Symbol is empty")
	assert.True(t, asset.Result.Name != "", "Name is empty")
	assert.True(t, asset.Result.Fiat == false || asset.Result.Fiat == true, "Fiat is empty")
	for _, v := range asset.Result.Markets.Base {
		assert.True(t, v.ID > 0, "ID is empty")
		assert.True(t, v.Exchange != "", "Exchange is empty")
		assert.True(t, v.Pair != "", "Pair is empty")
		assert.True(t, v.Active == false || v.Active == true, "Active is empty")
		assert.True(t, v.Route != "", "Route is empty")
	}
	for _, v := range asset.Result.Markets.Quote {
		assert.True(t, v.ID > 0, "ID is empty")
		assert.True(t, v.Exchange != "", "Exchange is empty")
		assert.True(t, v.Pair != "", "Pair is empty")
		assert.True(t, v.Active == false || v.Active == true, "Active is empty")
		assert.True(t, v.Route != "", "Route is empty")
	}
}

func TestPairs(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/pairs")
	checkErr(err)
	var pairs Pairs
	err = json.Unmarshal(getBytes(resp.Body), &pairs)
	checkErr(err)
	for _, v := range pairs.Result {
		assert.True(t, v.ID > 0, "ID is empty")
		assert.True(t, v.Symbol != "", "Symbol is empty")
		assert.True(t, v.Route != "", "Route is empty")
		//can be empty -> assert.True(t, v.FuturesContractPeriod != "", "FuturesContractPeriod is empty: "+v.FuturesContractPeriod)
		//Base
		assert.True(t, v.Base.ID > 0, "ID is empty")
		assert.True(t, v.Base.Symbol != "", "Symbol is empty")
		assert.True(t, v.Base.Name != "", "Name is empty")
		assert.True(t, v.Base.Fiat == false || v.Base.Fiat == true, "Fiat is empty")
		assert.True(t, v.Base.Route != "", "Route is empty")
		//Quote
		assert.True(t, v.Quote.ID > 0, "ID is empty")
		assert.True(t, v.Quote.Symbol != "", "Symbol is empty")
		assert.True(t, v.Quote.Name != "", "Name is empty")
		assert.True(t, v.Quote.Fiat == false || v.Quote.Fiat == true, "Fiat is empty")
		assert.True(t, v.Quote.Route != "", "Route is empty")
	}
}

func TestPair(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/pairs/ethbtc")
	checkErr(err)
	var pair Pair
	err = json.Unmarshal(getBytes(resp.Body), &pair)
	checkErr(err)

	assert.True(t, pair.Result.ID > 0, "ID is empty")
	assert.True(t, pair.Result.Symbol != "", "Symbol is empty")
	assert.True(t, pair.Result.Route != "", "Route is empty")
	//Base
	assert.True(t, pair.Result.Base.ID > 0, "ID is empty")
	assert.True(t, pair.Result.Base.Symbol != "", "Symbol is empty")
	assert.True(t, pair.Result.Base.Name != "", "Name is empty")
	assert.True(t, pair.Result.Base.Fiat == false || pair.Result.Base.Fiat == true, "Fiat is empty")
	assert.True(t, pair.Result.Base.Route != "", "Route is empty")
	//Quote
	assert.True(t, pair.Result.Quote.ID > 0, "ID is empty")
	assert.True(t, pair.Result.Quote.Symbol != "", "Symbol is empty")
	assert.True(t, pair.Result.Quote.Name != "", "Name is empty")
	assert.True(t, pair.Result.Quote.Fiat == false || pair.Result.Quote.Fiat == true, "Fiat is empty")
	assert.True(t, pair.Result.Quote.Route != "", "Route is empty")

	for _, v := range pair.Result.Markets {
		assert.True(t, v.ID > 0, "ID is empty")
		assert.True(t, v.Exchange != "", "Exchange is empty")
		assert.True(t, v.Pair != "", "Pair is empty")
		assert.True(t, v.Active == false || v.Active == true, "Active is empty")
		assert.True(t, v.Route != "", "Route is empty")
	}
}

func TestExchanges(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/exchanges")
	checkErr(err)
	var exchanges Exchanges
	err = json.Unmarshal(getBytes(resp.Body), &exchanges)
	checkErr(err)

	for _, v := range exchanges.Result {
		assert.True(t, v.Symbol != "", "Symbol is empty")
		assert.True(t, v.Name != "", "Name is empty")
		assert.True(t, v.Route != "", "Route is empty")
		assert.True(t, v.Active == false || v.Active == true, "Active is empty")
	}
}

func TestExchange(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/exchanges/kraken")
	checkErr(err)
	var exchange Exchange
	err = json.Unmarshal(getBytes(resp.Body), &exchange)
	checkErr(err)

	assert.True(t, exchange.Result.Symbol != "", "Symbol is empty")
	assert.True(t, exchange.Result.Name != "", "Name is empty")
	assert.True(t, exchange.Result.Active == false || exchange.Result.Active == true, "Active is empty")
	assert.True(t, exchange.Result.Routes.Markets != "", "Markets is empty")
}

func TestMarkets(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/markets")
	checkErr(err)
	var markets Markets
	err = json.Unmarshal(getBytes(resp.Body), &markets)
	checkErr(err)

	for _, v := range markets.Result {
		assert.True(t, v.ID > 0, "ID is empty")
		assert.True(t, v.Exchange != "", "Exchange is empty")
		assert.True(t, v.Pair != "", "Pair is empty")
		assert.True(t, v.Active == false || v.Active == true, "Active is empty")
		assert.True(t, v.Route != "", "Route is empty")
	}
}

func TestMarket(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/markets/gdax/btcusd")
	checkErr(err)
	var market Market
	err = json.Unmarshal(getBytes(resp.Body), &market)
	checkErr(err)

	assert.True(t, market.Result.Exchange != "", "Exchange is empty")
	assert.True(t, market.Result.Pair != "", "Pair is empty")
	assert.True(t, market.Result.Active == false || market.Result.Active == true, "Active is empty")
	assert.True(t, market.Result.Routes.Price != "", "Price is empty")
	assert.True(t, market.Result.Routes.Summary != "", "Summary is empty")
	assert.True(t, market.Result.Routes.Orderbook != "", "Orderbook is empty")
	assert.True(t, market.Result.Routes.Trades != "", "Trades is empty")
	assert.True(t, market.Result.Routes.Ohlc != "", "Ohlc is empty")
}

func TestPrice(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/markets/gdax/btcusd/price")
	checkErr(err)
	var price Price
	err = json.Unmarshal(getBytes(resp.Body), &price)
	checkErr(err)

	assert.True(t, price.Result.Price > 0.0, "Price is empty")
}

func TestSummary(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/markets/gdax/btcusd/summary")
	checkErr(err)
	var summary Summary
	err = json.Unmarshal(getBytes(resp.Body), &summary)
	checkErr(err)

	assert.True(t, summary.Result.Price.Last > 0.0, "Last is empty")
	assert.True(t, summary.Result.Price.High > 0.0, "High is empty")
	assert.True(t, summary.Result.Price.Low > 0.0, "Low is empty")
	assert.True(t, summary.Result.Price.Change.Percentage >= 0.0 || summary.Result.Price.Change.Percentage < 0.0, "Percentage is empty")
	assert.True(t, summary.Result.Price.Change.Absolute >= 0.0 || summary.Result.Price.Change.Absolute < 0.0, "Absolute is empty")
	assert.True(t, summary.Result.Volume > 0.0, "Volume is empty")
}

func checkErr(err error) {
	if err != nil {
		panic(err)

	}
}

func getBytes(reader io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	return buf.Bytes()
}
