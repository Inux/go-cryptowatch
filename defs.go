package main

import (
	"go-cryptowatch/models"
	"sync"
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

func getCostFactor() int {
	costfactor := 0
	for _, c := range currencies {
		costfactor += len(c.Markets)
	}
	return costfactor
}
