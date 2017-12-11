package models

//Assets - https://api.cryptowat.ch/assets
type Assets struct {
	Result []struct {
		ID     int    `json:"id"`
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
		Fiat   bool   `json:"fiat"`
		Route  string `json:"route"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

//Asset - https://api.cryptowat.ch/assets/"Assets.Symbol"
type Asset struct {
	Result struct {
		ID      int    `json:"id"`
		Symbol  string `json:"symbol"`
		Name    string `json:"name"`
		Fiat    bool   `json:"fiat"`
		Markets struct {
			Base []struct {
				ID       int    `json:"id"`
				Exchange string `json:"exchange"`
				Pair     string `json:"pair"`
				Active   bool   `json:"active"`
				Route    string `json:"route"`
			} `json:"base"`
			Quote []struct {
				ID       int    `json:"id"`
				Exchange string `json:"exchange"`
				Pair     string `json:"pair"`
				Active   bool   `json:"active"`
				Route    string `json:"route"`
			} `json:"quote"`
		} `json:"markets"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

//Pairs - https://api.cryptowat.ch/pairs
type Pairs struct {
	Result []struct {
		Symbol string `json:"symbol"`
		ID     int    `json:"id"`
		Base   struct {
			ID     int    `json:"id"`
			Symbol string `json:"symbol"`
			Name   string `json:"name"`
			Fiat   bool   `json:"fiat"`
			Route  string `json:"route"`
		} `json:"base"`
		Quote struct {
			ID     int    `json:"id"`
			Symbol string `json:"symbol"`
			Name   string `json:"name"`
			Fiat   bool   `json:"fiat"`
			Route  string `json:"route"`
		} `json:"quote"`
		Route                 string `json:"route"`
		FuturesContractPeriod string `json:"futuresContractPeriod,omitempty"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

//Pair - https://api.cryptowat.ch/pairs/"Pairs.Symbol"
type Pair struct {
	Result struct {
		Symbol string `json:"symbol"`
		ID     int    `json:"id"`
		Base   struct {
			ID     int    `json:"id"`
			Symbol string `json:"symbol"`
			Name   string `json:"name"`
			Fiat   bool   `json:"fiat"`
			Route  string `json:"route"`
		} `json:"base"`
		Quote struct {
			ID     int    `json:"id"`
			Symbol string `json:"symbol"`
			Name   string `json:"name"`
			Fiat   bool   `json:"fiat"`
			Route  string `json:"route"`
		} `json:"quote"`
		Route   string `json:"route"`
		Markets []struct {
			ID       int    `json:"id"`
			Exchange string `json:"exchange"`
			Pair     string `json:"pair"`
			Active   bool   `json:"active"`
			Route    string `json:"route"`
		} `json:"markets"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

//Exchanges - https://api.cryptowat.ch/exchanges
type Exchanges struct {
	Result []struct {
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
		Route  string `json:"route"`
		Active bool   `json:"active"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

//Exchange - https://api.cryptowat.ch/exchanges/"Exchanges.Symbol"
type Exchange struct {
	Result struct {
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
		Routes struct {
			Markets string `json:"markets"`
		} `json:"routes"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

//Markets - https://api.cryptowat.ch/markets
type Markets struct {
	Result []struct {
		ID       int    `json:"id"`
		Exchange string `json:"exchange"`
		Pair     string `json:"pair"`
		Active   bool   `json:"active"`
		Route    string `json:"route"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

//Market - https://api.cryptowat.ch/markets/"Exchange"/"Pair"
type Market struct {
	Result struct {
		Exchange string `json:"exchange"`
		Pair     string `json:"pair"`
		Active   bool   `json:"active"`
		Routes   struct {
			Price     string `json:"price"`
			Summary   string `json:"summary"`
			Orderbook string `json:"orderbook"`
			Trades    string `json:"trades"`
			Ohlc      string `json:"ohlc"`
		} `json:"routes"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

//Price - https://api.cryptowat.ch/markets/"Exchange"/"Pair"/price
type Price struct {
	Result struct {
		Price float64 `json:"price"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

//Summary - https://api.cryptowat.ch/markets/"Exchange"/"Pair"/summary
type Summary struct {
	Result struct {
		Price struct {
			Last   float64 `json:"last"`
			High   float64 `json:"high"`
			Low    float64 `json:"low"`
			Change struct {
				Percentage float64 `json:"percentage"`
				Absolute   float64 `json:"absolute"`
			} `json:"change"`
		} `json:"price"`
		Volume float64 `json:"volume"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

//Trades - https://api.cryptowat.ch/markets/"Exchange"/"Pair"/trades
//ID, Timestamp, Price, Amount
type Trades struct {
	Result    [][]float64 `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

//OrderBook - https://api.cryptowat.ch/markets/"Exchange"/"Pair"/orderbook
//Price, Amount
type OrderBook struct {
	Result struct {
		Asks [][]int     `json:"asks"`
		Bids [][]float64 `json:"bids"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

//OHLC - https://api.cryptowat.ch/markets/"Exchange"/"Pair"/ohlc
//CloseTime, OpenPrice, HighPrice, LowPrice, ClosePrice, Volume
type OHLC struct {
	Result []struct {
		TimeValue [][]float64 `json:"string"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

//Prices - https://api.cryptowat.ch/markets/prices
type Prices struct {
	Result []struct {
		MarketPair float64 `json:"string"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

//Summaries - https://api.cryptowat.ch/markets/summaries
type Summaries struct {
	Result []struct {
		MarketPair struct {
			Price struct {
				Last   float64 `json:"last"`
				High   float64 `json:"high"`
				Low    float64 `json:"low"`
				Change struct {
					Percentage float64 `json:"percentage"`
					Absolute   float64 `json:"absolute"`
				} `json:"change"`
			} `json:"price"`
			Volume float64 `json:"volume"`
		} `json:"string"`
	} `json:"result"`
}
