package cryptowatchmodels

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

func TestTrades(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/markets/gdax/btcusd/trades")
	checkErr(err)
	var trades Trades
	err = json.Unmarshal(getBytes(resp.Body), &trades)
	checkErr(err)

	//[ ID, Timestamp, Price, Amount ]
	for _, v := range trades.Result {
		assert.True(t, len(v) == 4, "Results length not 4")
	}
}

func TestOrderBook(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/markets/gdax/btcusd/orderbook")
	checkErr(err)
	var ob OrderBook
	err = json.Unmarshal(getBytes(resp.Body), &ob)
	checkErr(err)

	//[ Price, Amount ]
	for _, v := range ob.Result.Asks {
		assert.True(t, len(v) == 2, "Asks length not 2")
	}
	for _, v := range ob.Result.Bids {
		assert.True(t, len(v) == 2, "Bids length not 2")
	}
}

func TestOHLC(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/markets/gdax/btcusd/ohlc")
	checkErr(err)
	var ohlc OHLC
	err = json.Unmarshal(getBytes(resp.Body), &ohlc)
	checkErr(err)

	//[ CloseTime, OpenPrice, HighPrice, LowPrice, ClosePrice, Volume ]
	for _, v := range ohlc.Result.One4400 {
		assert.True(t, len(v) > 0, "One4400 length is 0")
	}

	for _, v := range ohlc.Result.One80 {
		assert.True(t, len(v) > 0, "One80 length is 0")
	}

	for _, v := range ohlc.Result.One800 {
		assert.True(t, len(v) > 0, "One800 length is 0")
	}

	for _, v := range ohlc.Result.Two1600 {
		assert.True(t, len(v) > 0, "Two1600 length is 0")
	}

	for _, v := range ohlc.Result.Two59200 {
		assert.True(t, len(v) > 0, "Two59200 length is 0")
	}

	for _, v := range ohlc.Result.Three00 {
		assert.True(t, len(v) > 0, "Three00 length is 0")
	}

	for _, v := range ohlc.Result.Three600 {
		assert.True(t, len(v) > 0, "Three600 length is 0")
	}

	for _, v := range ohlc.Result.Four3200 {
		assert.True(t, len(v) > 0, "Four3200 length is 0")
	}

	for _, v := range ohlc.Result.Six0 {
		assert.True(t, len(v) > 0, "Six0 length is 0")
	}

	for _, v := range ohlc.Result.Six04800 {
		assert.True(t, len(v) > 0, "Six04800 length is 0")
	}

	for _, v := range ohlc.Result.Seven200 {
		assert.True(t, len(v) > 0, "Seven200 length is 0")
	}

	for _, v := range ohlc.Result.Eight6400 {
		assert.True(t, len(v) > 0, "Eight6400 length is 0")
	}

	for _, v := range ohlc.Result.Nine00 {
		assert.True(t, len(v) > 0, "Nine00 length is 0")
	}
}

func TestPrices(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/markets/prices")
	checkErr(err)
	var prices Prices
	err = json.Unmarshal(getBytes(resp.Body), &prices)
	checkErr(err)

	assert.True(t, prices.Result.Bitfinex_btcusd > 0.0)
}

func TestSummaries(t *testing.T) {
	resp, err := http.Get("https://api.cryptowat.ch/markets/summaries")
	checkErr(err)
	var summaries Summaries
	err = json.Unmarshal(getBytes(resp.Body), &summaries)
	checkErr(err)

	assert.True(t, summaries.Result.Bitfinex_btcusd.Price.Low > 0.0)
	assert.True(t, summaries.Result.Bitfinex_btcusd.Price.Last > 0.0)
	assert.True(t, summaries.Result.Bitfinex_btcusd.Price.High > 0.0)
	assert.True(t, summaries.Result.Bitfinex_btcusd.Price.Change.Absolute > 0.0)
	assert.True(t, summaries.Result.Bitfinex_btcusd.Price.Change.Percentage > 0.0)
	assert.True(t, summaries.Result.Bitfinex_btcusd.Volume > 0.0)
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
