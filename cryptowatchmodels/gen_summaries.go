package cryptowatchmodels

type Summaries struct {
	Allowance Summaries_sub1  `json:"allowance"`
	Result    Summaries_sub46 `json:"result"`
}

type Summaries_sub5 struct {
	Absolute   float64 `json:"absolute"`
	Percentage float64 `json:"percentage"`
}

type Summaries_sub2 struct {
	Absolute   float64   `json:"absolute"`
	Percentage float64 `json:"percentage"`
}

type Summaries_sub8 struct {
	Absolute   float64 `json:"absolute"`
	Percentage float64 `json:"percentage"`
}

type Summaries_sub46 struct {
	Bitfinex_avtbtc                   Summaries_sub4  `json:"bitfinex:avtbtc"`
	Bitfinex_avteth                   Summaries_sub7  `json:"bitfinex:avteth"`
	Bitfinex_avtusd                   Summaries_sub7  `json:"bitfinex:avtusd"`
	Bitfinex_bccbtc                   Summaries_sub10 `json:"bitfinex:bccbtc"`
	Bitfinex_bccusd                   Summaries_sub12 `json:"bitfinex:bccusd"`
	Bitfinex_bchbtc                   Summaries_sub7  `json:"bitfinex:bchbtc"`
	Bitfinex_bcheth                   Summaries_sub7  `json:"bitfinex:bcheth"`
	Bitfinex_bchusd                   Summaries_sub14 `json:"bitfinex:bchusd"`
	Bitfinex_bcubtc                   Summaries_sub4  `json:"bitfinex:bcubtc"`
	Bitfinex_bcuusd                   Summaries_sub16 `json:"bitfinex:bcuusd"`
	Bitfinex_bt1btc                   Summaries_sub7  `json:"bitfinex:bt1btc"`
	Bitfinex_bt1usd                   Summaries_sub12 `json:"bitfinex:bt1usd"`
	Bitfinex_bt2btc                   Summaries_sub7  `json:"bitfinex:bt2btc"`
	Bitfinex_bt2usd                   Summaries_sub18 `json:"bitfinex:bt2usd"`
	Bitfinex_btceur                   Summaries_sub20 `json:"bitfinex:btceur"`
	Bitfinex_btcusd                   Summaries_sub22 `json:"bitfinex:btcusd"`
	Bitfinex_btgbtc                   Summaries_sub7  `json:"bitfinex:btgbtc"`
	Bitfinex_btgusd                   Summaries_sub18 `json:"bitfinex:btgusd"`
	Bitfinex_dashbtc                  Summaries_sub7  `json:"bitfinex:dashbtc"`
	Bitfinex_dashusd                  Summaries_sub7  `json:"bitfinex:dashusd"`
	Bitfinex_databtc                  Summaries_sub24 `json:"bitfinex:databtc"`
	Bitfinex_dataeth                  Summaries_sub24 `json:"bitfinex:dataeth"`
	Bitfinex_datausd                  Summaries_sub24 `json:"bitfinex:datausd"`
	Bitfinex_datbtc                   Summaries_sub12 `json:"bitfinex:datbtc"`
	Bitfinex_dateth                   Summaries_sub25 `json:"bitfinex:dateth"`
	Bitfinex_datusd                   Summaries_sub26 `json:"bitfinex:datusd"`
	Bitfinex_edobtc                   Summaries_sub4  `json:"bitfinex:edobtc"`
	Bitfinex_edoeth                   Summaries_sub7  `json:"bitfinex:edoeth"`
	Bitfinex_edousd                   Summaries_sub7  `json:"bitfinex:edousd"`
	Bitfinex_eosbtc                   Summaries_sub4  `json:"bitfinex:eosbtc"`
	Bitfinex_eoseth                   Summaries_sub7  `json:"bitfinex:eoseth"`
	Bitfinex_eosusd                   Summaries_sub26 `json:"bitfinex:eosusd"`
	Bitfinex_etcbtc                   Summaries_sub26 `json:"bitfinex:etcbtc"`
	Bitfinex_etcusd                   Summaries_sub26 `json:"bitfinex:etcusd"`
	Bitfinex_ethbtc                   Summaries_sub7  `json:"bitfinex:ethbtc"`
	Bitfinex_ethusd                   Summaries_sub28 `json:"bitfinex:ethusd"`
	Bitfinex_etpbtc                   Summaries_sub4  `json:"bitfinex:etpbtc"`
	Bitfinex_etpeth                   Summaries_sub7  `json:"bitfinex:etpeth"`
	Bitfinex_etpusd                   Summaries_sub7  `json:"bitfinex:etpusd"`
	Bitfinex_iotbtc                   Summaries_sub25 `json:"bitfinex:iotbtc"`
	Bitfinex_ioteth                   Summaries_sub7  `json:"bitfinex:ioteth"`
	Bitfinex_iotusd                   Summaries_sub26 `json:"bitfinex:iotusd"`
	Bitfinex_ltcbtc                   Summaries_sub7  `json:"bitfinex:ltcbtc"`
	Bitfinex_ltcusd                   Summaries_sub30 `json:"bitfinex:ltcusd"`
	Bitfinex_neobtc                   Summaries_sub7  `json:"bitfinex:neobtc"`
	Bitfinex_neoeth                   Summaries_sub7  `json:"bitfinex:neoeth"`
	Bitfinex_neousd                   Summaries_sub7  `json:"bitfinex:neousd"`
	Bitfinex_omgbtc                   Summaries_sub4  `json:"bitfinex:omgbtc"`
	Bitfinex_omgeth                   Summaries_sub7  `json:"bitfinex:omgeth"`
	Bitfinex_omgusd                   Summaries_sub30 `json:"bitfinex:omgusd"`
	Bitfinex_qtumbtc                  Summaries_sub4  `json:"bitfinex:qtumbtc"`
	Bitfinex_qtumeth                  Summaries_sub7  `json:"bitfinex:qtumeth"`
	Bitfinex_qtumusd                  Summaries_sub7  `json:"bitfinex:qtumusd"`
	Bitfinex_sanbtc                   Summaries_sub4  `json:"bitfinex:sanbtc"`
	Bitfinex_saneth                   Summaries_sub7  `json:"bitfinex:saneth"`
	Bitfinex_sanusd                   Summaries_sub7  `json:"bitfinex:sanusd"`
	Bitfinex_xmrbtc                   Summaries_sub7  `json:"bitfinex:xmrbtc"`
	Bitfinex_xmrusd                   Summaries_sub30 `json:"bitfinex:xmrusd"`
	Bitfinex_xrpbtc                   Summaries_sub31 `json:"bitfinex:xrpbtc"`
	Bitfinex_xrpusd                   Summaries_sub26 `json:"bitfinex:xrpusd"`
	Bitfinex_zecbtc                   Summaries_sub7  `json:"bitfinex:zecbtc"`
	Bitfinex_zecusd                   Summaries_sub18 `json:"bitfinex:zecusd"`
	Bitflyer_bchbtc                   Summaries_sub7  `json:"bitflyer:bchbtc"`
	Bitflyer_btcfxjpy                 Summaries_sub12 `json:"bitflyer:btcfxjpy"`
	Bitflyer_btcjpy                   Summaries_sub12 `json:"bitflyer:btcjpy"`
	Bitflyer_btcjpy_biweekly_futures  Summaries_sub12 `json:"bitflyer:btcjpy-biweekly-futures"`
	Bitflyer_btcjpy_weekly_futures    Summaries_sub12 `json:"bitflyer:btcjpy-weekly-futures"`
	Bitflyer_btcusd                   Summaries_sub7  `json:"bitflyer:btcusd"`
	Bitflyer_ethbtc                   Summaries_sub7  `json:"bitflyer:ethbtc"`
	Bithumb_bchkrw                    Summaries_sub12 `json:"bithumb:bchkrw"`
	Bithumb_btckrw                    Summaries_sub12 `json:"bithumb:btckrw"`
	Bithumb_btgkrw                    Summaries_sub12 `json:"bithumb:btgkrw"`
	Bithumb_dashkrw                   Summaries_sub12 `json:"bithumb:dashkrw"`
	Bithumb_etckrw                    Summaries_sub31 `json:"bithumb:etckrw"`
	Bithumb_ethkrw                    Summaries_sub12 `json:"bithumb:ethkrw"`
	Bithumb_ltckrw                    Summaries_sub12 `json:"bithumb:ltckrw"`
	Bithumb_qtumbtc                   Summaries_sub31 `json:"bithumb:qtumbtc"`
	Bithumb_qtumeth                   Summaries_sub31 `json:"bithumb:qtumeth"`
	Bithumb_qtumkrw                   Summaries_sub12 `json:"bithumb:qtumkrw"`
	Bithumb_xmrkrw                    Summaries_sub12 `json:"bithumb:xmrkrw"`
	Bithumb_xrpkrw                    Summaries_sub31 `json:"bithumb:xrpkrw"`
	Bithumb_zeckrw                    Summaries_sub12 `json:"bithumb:zeckrw"`
	Bitmex_bchbtc_monthly_futures     Summaries_sub26 `json:"bitmex:bchbtc-monthly-futures"`
	Bitmex_bchbtc_quarterly_futures   Summaries_sub24 `json:"bitmex:bchbtc-quarterly-futures"`
	Bitmex_btcjpy_monthly_futures     Summaries_sub33 `json:"bitmex:btcjpy-monthly-futures"`
	Bitmex_btcjpy_quarterly_futures   Summaries_sub34 `json:"bitmex:btcjpy-quarterly-futures"`
	Bitmex_btcusd_monthly_futures     Summaries_sub33 `json:"bitmex:btcusd-monthly-futures"`
	Bitmex_btcusd_perpetual_futures   Summaries_sub36 `json:"bitmex:btcusd-perpetual-futures"`
	Bitmex_btcusd_quarterly_futures   Summaries_sub34 `json:"bitmex:btcusd-quarterly-futures"`
	Bitmex_dashbtc_monthly_futures    Summaries_sub26 `json:"bitmex:dashbtc-monthly-futures"`
	Bitmex_dashbtc_quarterly_futures  Summaries_sub26 `json:"bitmex:dashbtc-quarterly-futures"`
	Bitmex_etcbtc_weekly_futures      Summaries_sub25 `json:"bitmex:etcbtc-weekly-futures"`
	Bitmex_ethbtc_monthly_futures     Summaries_sub26 `json:"bitmex:ethbtc-monthly-futures"`
	Bitmex_ethbtc_quarterly_futures   Summaries_sub26 `json:"bitmex:ethbtc-quarterly-futures"`
	Bitmex_ltcbtc_monthly_futures     Summaries_sub26 `json:"bitmex:ltcbtc-monthly-futures"`
	Bitmex_ltcbtc_quarterly_futures   Summaries_sub26 `json:"bitmex:ltcbtc-quarterly-futures"`
	Bitmex_xmrbtc_monthly_futures     Summaries_sub26 `json:"bitmex:xmrbtc-monthly-futures"`
	Bitmex_xmrbtc_quarterly_futures   Summaries_sub26 `json:"bitmex:xmrbtc-quarterly-futures"`
	Bitmex_xrpbtc_monthly_futures     Summaries_sub31 `json:"bitmex:xrpbtc-monthly-futures"`
	Bitmex_xrpbtc_quarterly_futures   Summaries_sub31 `json:"bitmex:xrpbtc-quarterly-futures"`
	Bitmex_xtzbtc_biquarterly_futures Summaries_sub25 `json:"bitmex:xtzbtc-biquarterly-futures"`
	Bitmex_xtzbtc_monthly_futures     Summaries_sub25 `json:"bitmex:xtzbtc-monthly-futures"`
	Bitmex_xtzbtc_quarterly_futures   Summaries_sub24 `json:"bitmex:xtzbtc-quarterly-futures"`
	Bitmex_zecbtc_monthly_futures     Summaries_sub26 `json:"bitmex:zecbtc-monthly-futures"`
	Bitmex_zecbtc_quarterly_futures   Summaries_sub26 `json:"bitmex:zecbtc-quarterly-futures"`
	Bitsquare_btcaud                  Summaries_sub24 `json:"bitsquare:btcaud"`
	Bitsquare_btceur                  Summaries_sub7  `json:"bitsquare:btceur"`
	Bitsquare_btcusd                  Summaries_sub14 `json:"bitsquare:btcusd"`
	Bitsquare_etcbtc                  Summaries_sub24 `json:"bitsquare:etcbtc"`
	Bitsquare_ethbtc                  Summaries_sub24 `json:"bitsquare:ethbtc"`
	Bitsquare_ltcbtc                  Summaries_sub24 `json:"bitsquare:ltcbtc"`
	Bitstamp_bchbtc                   Summaries_sub7  `json:"bitstamp:bchbtc"`
	Bitstamp_bcheur                   Summaries_sub20 `json:"bitstamp:bcheur"`
	Bitstamp_bchusd                   Summaries_sub38 `json:"bitstamp:bchusd"`
	Bitstamp_btceur                   Summaries_sub30 `json:"bitstamp:btceur"`
	Bitstamp_btcusd                   Summaries_sub7  `json:"bitstamp:btcusd"`
	Bitstamp_ethbtc                   Summaries_sub7  `json:"bitstamp:ethbtc"`
	Bitstamp_etheur                   Summaries_sub7  `json:"bitstamp:etheur"`
	Bitstamp_ethusd                   Summaries_sub16 `json:"bitstamp:ethusd"`
	Bitstamp_eurusd                   Summaries_sub26 `json:"bitstamp:eurusd"`
	Bitstamp_ltcbtc                   Summaries_sub7  `json:"bitstamp:ltcbtc"`
	Bitstamp_ltceur                   Summaries_sub30 `json:"bitstamp:ltceur"`
	Bitstamp_ltcusd                   Summaries_sub7  `json:"bitstamp:ltcusd"`
	Bitstamp_xrpbtc                   Summaries_sub31 `json:"bitstamp:xrpbtc"`
	Bitstamp_xrpeur                   Summaries_sub26 `json:"bitstamp:xrpeur"`
	Bitstamp_xrpusd                   Summaries_sub26 `json:"bitstamp:xrpusd"`
	Bittrex_batbtc                    Summaries_sub31 `json:"bittrex:batbtc"`
	Bittrex_bccbtc                    Summaries_sub7  `json:"bittrex:bccbtc"`
	Bittrex_bccusdt                   Summaries_sub28 `json:"bittrex:bccusdt"`
	Bittrex_blkbtc                    Summaries_sub12 `json:"bittrex:blkbtc"`
	Bittrex_btcusdt                   Summaries_sub18 `json:"bittrex:btcusdt"`
	Bittrex_dcrbtc                    Summaries_sub4  `json:"bittrex:dcrbtc"`
	Bittrex_edgbtc                    Summaries_sub31 `json:"bittrex:edgbtc"`
	Bittrex_etcbtc                    Summaries_sub26 `json:"bittrex:etcbtc"`
	Bittrex_etceth                    Summaries_sub7  `json:"bittrex:etceth"`
	Bittrex_etcusdt                   Summaries_sub7  `json:"bittrex:etcusdt"`
	Bittrex_ethbtc                    Summaries_sub7  `json:"bittrex:ethbtc"`
	Bittrex_ethusdt                   Summaries_sub28 `json:"bittrex:ethusdt"`
	Bittrex_neobtc                    Summaries_sub7  `json:"bittrex:neobtc"`
	Bittrex_neoeth                    Summaries_sub7  `json:"bittrex:neoeth"`
	Bittrex_neousdt                   Summaries_sub7  `json:"bittrex:neousdt"`
	Bittrex_omgbtc                    Summaries_sub4  `json:"bittrex:omgbtc"`
	Bittrex_omgeth                    Summaries_sub7  `json:"bittrex:omgeth"`
	Bittrex_qtumbtc                   Summaries_sub4  `json:"bittrex:qtumbtc"`
	Bittrex_stratbtc                  Summaries_sub4  `json:"bittrex:stratbtc"`
	Bittrex_vrcbtc                    Summaries_sub31 `json:"bittrex:vrcbtc"`
	Bittrex_vrmbtc                    Summaries_sub4  `json:"bittrex:vrmbtc"`
	Bittrex_vtcbtc                    Summaries_sub4  `json:"bittrex:vtcbtc"`
	Bittrex_xrpusdt                   Summaries_sub26 `json:"bittrex:xrpusdt"`
	Bittrex_zecusdt                   Summaries_sub7  `json:"bittrex:zecusdt"`
	Bittrex_zenbtc                    Summaries_sub4  `json:"bittrex:zenbtc"`
	Btce_bchbtc                       Summaries_sub7  `json:"btce:bchbtc"`
	Btce_bchusd                       Summaries_sub7  `json:"btce:bchusd"`
	Btce_btceur                       Summaries_sub7  `json:"btce:btceur"`
	Btce_btcrur                       Summaries_sub16 `json:"btce:btcrur"`
	Btce_btcusd                       Summaries_sub18 `json:"btce:btcusd"`
	Btce_ethbtc                       Summaries_sub7  `json:"btce:ethbtc"`
	Btce_etheur                       Summaries_sub7  `json:"btce:etheur"`
	Btce_ethltc                       Summaries_sub7  `json:"btce:ethltc"`
	Btce_ethrur                       Summaries_sub7  `json:"btce:ethrur"`
	Btce_ethusd                       Summaries_sub7  `json:"btce:ethusd"`
	Btce_ltcbtc                       Summaries_sub7  `json:"btce:ltcbtc"`
	Btce_ltcusd                       Summaries_sub7  `json:"btce:ltcusd"`
	Btce_nmcusd                       Summaries_sub7  `json:"btce:nmcusd"`
	Btce_ppcusd                       Summaries_sub7  `json:"btce:ppcusd"`
	Btce_zecbtc                       Summaries_sub7  `json:"btce:zecbtc"`
	Btce_zecusd                       Summaries_sub7  `json:"btce:zecusd"`
	Cexio_bchusd                      Summaries_sub16 `json:"cexio:bchusd"`
	Cexio_btceur                      Summaries_sub22 `json:"cexio:btceur"`
	Cexio_btcusd                      Summaries_sub30 `json:"cexio:btcusd"`
	Cexio_ethbtc                      Summaries_sub7  `json:"cexio:ethbtc"`
	Cexio_ethusd                      Summaries_sub14 `json:"cexio:ethusd"`
	Cexio_ltcbtc                      Summaries_sub24 `json:"cexio:ltcbtc"`
	Cexio_ltcusd                      Summaries_sub24 `json:"cexio:ltcusd"`
	Gdax_btceur                       Summaries_sub22 `json:"gdax:btceur"`
	Gdax_btcgbp                       Summaries_sub30 `json:"gdax:btcgbp"`
	Gdax_btcusd                       Summaries_sub20 `json:"gdax:btcusd"`
	Gdax_ethbtc                       Summaries_sub7  `json:"gdax:ethbtc"`
	Gdax_etheur                       Summaries_sub7  `json:"gdax:etheur"`
	Gdax_ethusd                       Summaries_sub7  `json:"gdax:ethusd"`
	Gdax_ltcbtc                       Summaries_sub7  `json:"gdax:ltcbtc"`
	Gdax_ltceur                       Summaries_sub7  `json:"gdax:ltceur"`
	Gdax_ltcusd                       Summaries_sub39 `json:"gdax:ltcusd"`
	Gemini_btcusd                     Summaries_sub30 `json:"gemini:btcusd"`
	Gemini_ethbtc                     Summaries_sub7  `json:"gemini:ethbtc"`
	Gemini_ethusd                     Summaries_sub16 `json:"gemini:ethusd"`
	Kraken_bchbtc                     Summaries_sub7  `json:"kraken:bchbtc"`
	Kraken_bcheur                     Summaries_sub20 `json:"kraken:bcheur"`
	Kraken_bchusd                     Summaries_sub7  `json:"kraken:bchusd"`
	Kraken_btccad                     Summaries_sub20 `json:"kraken:btccad"`
	Kraken_btceur                     Summaries_sub41 `json:"kraken:btceur"`
	Kraken_btcgbp                     Summaries_sub12 `json:"kraken:btcgbp"`
	Kraken_btcjpy                     Summaries_sub12 `json:"kraken:btcjpy"`
	Kraken_btcusd                     Summaries_sub12 `json:"kraken:btcusd"`
	Kraken_dashbtc                    Summaries_sub7  `json:"kraken:dashbtc"`
	Kraken_dasheur                    Summaries_sub18 `json:"kraken:dasheur"`
	Kraken_dashusd                    Summaries_sub30 `json:"kraken:dashusd"`
	Kraken_dogebtc                    Summaries_sub31 `json:"kraken:dogebtc"`
	Kraken_eosbtc                     Summaries_sub4  `json:"kraken:eosbtc"`
	Kraken_eoseth                     Summaries_sub7  `json:"kraken:eoseth"`
	Kraken_etcbtc                     Summaries_sub7  `json:"kraken:etcbtc"`
	Kraken_etceth                     Summaries_sub7  `json:"kraken:etceth"`
	Kraken_etceur                     Summaries_sub7  `json:"kraken:etceur"`
	Kraken_etcusd                     Summaries_sub7  `json:"kraken:etcusd"`
	Kraken_ethbtc                     Summaries_sub7  `json:"kraken:ethbtc"`
	Kraken_ethcad                     Summaries_sub38 `json:"kraken:ethcad"`
	Kraken_etheur                     Summaries_sub30 `json:"kraken:etheur"`
	Kraken_ethgbp                     Summaries_sub14 `json:"kraken:ethgbp"`
	Kraken_ethjpy                     Summaries_sub12 `json:"kraken:ethjpy"`
	Kraken_ethusd                     Summaries_sub30 `json:"kraken:ethusd"`
	Kraken_gnobtc                     Summaries_sub7  `json:"kraken:gnobtc"`
	Kraken_gnoeth                     Summaries_sub7  `json:"kraken:gnoeth"`
	Kraken_gnoeur                     Summaries_sub24 `json:"kraken:gnoeur"`
	Kraken_gnousd                     Summaries_sub24 `json:"kraken:gnousd"`
	Kraken_icnbtc                     Summaries_sub12 `json:"kraken:icnbtc"`
	Kraken_icneth                     Summaries_sub7  `json:"kraken:icneth"`
	Kraken_ltcbtc                     Summaries_sub7  `json:"kraken:ltcbtc"`
	Kraken_ltccad                     Summaries_sub24 `json:"kraken:ltccad"`
	Kraken_ltceur                     Summaries_sub18 `json:"kraken:ltceur"`
	Kraken_ltcusd                     Summaries_sub30 `json:"kraken:ltcusd"`
	Kraken_mlnbtc                     Summaries_sub7  `json:"kraken:mlnbtc"`
	Kraken_mlneth                     Summaries_sub7  `json:"kraken:mlneth"`
	Kraken_repbtc                     Summaries_sub7  `json:"kraken:repbtc"`
	Kraken_repcad                     Summaries_sub24 `json:"kraken:repcad"`
	Kraken_repeth                     Summaries_sub7  `json:"kraken:repeth"`
	Kraken_repeur                     Summaries_sub7  `json:"kraken:repeur"`
	Kraken_repjpy                     Summaries_sub24 `json:"kraken:repjpy"`
	Kraken_repusd                     Summaries_sub24 `json:"kraken:repusd"`
	Kraken_strbtc                     Summaries_sub31 `json:"kraken:strbtc"`
	Kraken_streur                     Summaries_sub24 `json:"kraken:streur"`
	Kraken_strusd                     Summaries_sub24 `json:"kraken:strusd"`
	Kraken_usdtusd                    Summaries_sub26 `json:"kraken:usdtusd"`
	Kraken_xmrbtc                     Summaries_sub7  `json:"kraken:xmrbtc"`
	Kraken_xmreur                     Summaries_sub7  `json:"kraken:xmreur"`
	Kraken_xmrusd                     Summaries_sub20 `json:"kraken:xmrusd"`
	Kraken_xrpbtc                     Summaries_sub31 `json:"kraken:xrpbtc"`
	Kraken_xrpcad                     Summaries_sub24 `json:"kraken:xrpcad"`
	Kraken_xrpeur                     Summaries_sub26 `json:"kraken:xrpeur"`
	Kraken_xrpjpy                     Summaries_sub24 `json:"kraken:xrpjpy"`
	Kraken_xrpusd                     Summaries_sub26 `json:"kraken:xrpusd"`
	Kraken_zecbtc                     Summaries_sub7  `json:"kraken:zecbtc"`
	Kraken_zeccad                     Summaries_sub24 `json:"kraken:zeccad"`
	Kraken_zeceur                     Summaries_sub16 `json:"kraken:zeceur"`
	Kraken_zecgbp                     Summaries_sub24 `json:"kraken:zecgbp"`
	Kraken_zecjpy                     Summaries_sub24 `json:"kraken:zecjpy"`
	Kraken_zecusd                     Summaries_sub7  `json:"kraken:zecusd"`
	Luno_btczar                       Summaries_sub12 `json:"luno:btczar"`
	Okcoin_btccny                     Summaries_sub24 `json:"okcoin:btccny"`
	Okcoin_btcusd                     Summaries_sub7  `json:"okcoin:btcusd"`
	Okcoin_btcusd_biweekly_futures    Summaries_sub26 `json:"okcoin:btcusd-biweekly-futures"`
	Okcoin_btcusd_quarterly_futures   Summaries_sub42 `json:"okcoin:btcusd-quarterly-futures"`
	Okcoin_btcusd_weekly_futures      Summaries_sub34 `json:"okcoin:btcusd-weekly-futures"`
	Okcoin_ltccny                     Summaries_sub24 `json:"okcoin:ltccny"`
	Okcoin_ltcusd                     Summaries_sub43 `json:"okcoin:ltcusd"`
	Okcoin_ltcusd_biweekly_futures    Summaries_sub26 `json:"okcoin:ltcusd-biweekly-futures"`
	Okcoin_ltcusd_quarterly_futures   Summaries_sub26 `json:"okcoin:ltcusd-quarterly-futures"`
	Okcoin_ltcusd_weekly_futures      Summaries_sub26 `json:"okcoin:ltcusd-weekly-futures"`
	Poloniex_ampbtc                   Summaries_sub12 `json:"poloniex:ampbtc"`
	Poloniex_ardrbtc                  Summaries_sub12 `json:"poloniex:ardrbtc"`
	Poloniex_bchbtc                   Summaries_sub7  `json:"poloniex:bchbtc"`
	Poloniex_bcheth                   Summaries_sub7  `json:"poloniex:bcheth"`
	Poloniex_bchusdt                  Summaries_sub45 `json:"poloniex:bchusdt"`
	Poloniex_bcnbtc                   Summaries_sub31 `json:"poloniex:bcnbtc"`
	Poloniex_bcnxmr                   Summaries_sub12 `json:"poloniex:bcnxmr"`
	Poloniex_bcybtc                   Summaries_sub12 `json:"poloniex:bcybtc"`
	Poloniex_belabtc                  Summaries_sub12 `json:"poloniex:belabtc"`
	Poloniex_blkbtc                   Summaries_sub12 `json:"poloniex:blkbtc"`
	Poloniex_blkxmr                   Summaries_sub4  `json:"poloniex:blkxmr"`
	Poloniex_btcdbtc                  Summaries_sub7  `json:"poloniex:btcdbtc"`
	Poloniex_btcdxmr                  Summaries_sub7  `json:"poloniex:btcdxmr"`
	Poloniex_btcusdt                  Summaries_sub18 `json:"poloniex:btcusdt"`
	Poloniex_btmbtc                   Summaries_sub12 `json:"poloniex:btmbtc"`
	Poloniex_btsbtc                   Summaries_sub31 `json:"poloniex:btsbtc"`
	Poloniex_burstbtc                 Summaries_sub31 `json:"poloniex:burstbtc"`
	Poloniex_clambtc                  Summaries_sub4  `json:"poloniex:clambtc"`
	Poloniex_cvcbtc                   Summaries_sub12 `json:"poloniex:cvcbtc"`
	Poloniex_cvceth                   Summaries_sub4  `json:"poloniex:cvceth"`
	Poloniex_daobtc                   Summaries_sub24 `json:"poloniex:daobtc"`
	Poloniex_daoeth                   Summaries_sub24 `json:"poloniex:daoeth"`
	Poloniex_dashbtc                  Summaries_sub7  `json:"poloniex:dashbtc"`
	Poloniex_dashusdt                 Summaries_sub20 `json:"poloniex:dashusdt"`
	Poloniex_dashxmr                  Summaries_sub7  `json:"poloniex:dashxmr"`
	Poloniex_dcrbtc                   Summaries_sub7  `json:"poloniex:dcrbtc"`
	Poloniex_dgbbtc                   Summaries_sub31 `json:"poloniex:dgbbtc"`
	Poloniex_dogebtc                  Summaries_sub31 `json:"poloniex:dogebtc"`
	Poloniex_emc2btc                  Summaries_sub45 `json:"poloniex:emc2btc"`
	Poloniex_etcbtc                   Summaries_sub7  `json:"poloniex:etcbtc"`
	Poloniex_etceth                   Summaries_sub7  `json:"poloniex:etceth"`
	Poloniex_etcusdt                  Summaries_sub18 `json:"poloniex:etcusdt"`
	Poloniex_ethbtc                   Summaries_sub7  `json:"poloniex:ethbtc"`
	Poloniex_ethusdt                  Summaries_sub7  `json:"poloniex:ethusdt"`
	Poloniex_expbtc                   Summaries_sub4  `json:"poloniex:expbtc"`
	Poloniex_fctbtc                   Summaries_sub7  `json:"poloniex:fctbtc"`
	Poloniex_fldcbtc                  Summaries_sub31 `json:"poloniex:fldcbtc"`
	Poloniex_flobtc                   Summaries_sub12 `json:"poloniex:flobtc"`
	Poloniex_gamebtc                  Summaries_sub4  `json:"poloniex:gamebtc"`
	Poloniex_gasbtc                   Summaries_sub4  `json:"poloniex:gasbtc"`
	Poloniex_gaseth                   Summaries_sub7  `json:"poloniex:gaseth"`
	Poloniex_gnobtc                   Summaries_sub7  `json:"poloniex:gnobtc"`
	Poloniex_gnoeth                   Summaries_sub7  `json:"poloniex:gnoeth"`
	Poloniex_gntbtc                   Summaries_sub12 `json:"poloniex:gntbtc"`
	Poloniex_gnteth                   Summaries_sub7  `json:"poloniex:gnteth"`
	Poloniex_grcbtc                   Summaries_sub31 `json:"poloniex:grcbtc"`
	Poloniex_hucbtc                   Summaries_sub12 `json:"poloniex:hucbtc"`
	Poloniex_lbcbtc                   Summaries_sub12 `json:"poloniex:lbcbtc"`
	Poloniex_lskbtc                   Summaries_sub4  `json:"poloniex:lskbtc"`
	Poloniex_lsketh                   Summaries_sub7  `json:"poloniex:lsketh"`
	Poloniex_ltcbtc                   Summaries_sub7  `json:"poloniex:ltcbtc"`
	Poloniex_ltcusdt                  Summaries_sub30 `json:"poloniex:ltcusdt"`
	Poloniex_ltcxmr                   Summaries_sub7  `json:"poloniex:ltcxmr"`
	Poloniex_maidbtc                  Summaries_sub12 `json:"poloniex:maidbtc"`
	Poloniex_maidxmr                  Summaries_sub4  `json:"poloniex:maidxmr"`
	Poloniex_nautbtc                  Summaries_sub24 `json:"poloniex:nautbtc"`
	Poloniex_navbtc                   Summaries_sub4  `json:"poloniex:navbtc"`
	Poloniex_neosbtc                  Summaries_sub4  `json:"poloniex:neosbtc"`
	Poloniex_nmcbtc                   Summaries_sub4  `json:"poloniex:nmcbtc"`
	Poloniex_notebtc                  Summaries_sub24 `json:"poloniex:notebtc"`
	Poloniex_nxcbtc                   Summaries_sub31 `json:"poloniex:nxcbtc"`
	Poloniex_nxtbtc                   Summaries_sub31 `json:"poloniex:nxtbtc"`
	Poloniex_nxtusdt                  Summaries_sub7  `json:"poloniex:nxtusdt"`
	Poloniex_nxtxmr                   Summaries_sub7  `json:"poloniex:nxtxmr"`
	Poloniex_omgbtc                   Summaries_sub4  `json:"poloniex:omgbtc"`
	Poloniex_omgeth                   Summaries_sub7  `json:"poloniex:omgeth"`
	Poloniex_omnibtc                  Summaries_sub4  `json:"poloniex:omnibtc"`
	Poloniex_pascbtc                  Summaries_sub12 `json:"poloniex:pascbtc"`
	Poloniex_pinkbtc                  Summaries_sub12 `json:"poloniex:pinkbtc"`
	Poloniex_potbtc                   Summaries_sub12 `json:"poloniex:potbtc"`
	Poloniex_ppcbtc                   Summaries_sub4  `json:"poloniex:ppcbtc"`
	Poloniex_radsbtc                  Summaries_sub4  `json:"poloniex:radsbtc"`
	Poloniex_repbtc                   Summaries_sub7  `json:"poloniex:repbtc"`
	Poloniex_repeth                   Summaries_sub7  `json:"poloniex:repeth"`
	Poloniex_repusdt                  Summaries_sub30 `json:"poloniex:repusdt"`
	Poloniex_ricbtc                   Summaries_sub31 `json:"poloniex:ricbtc"`
	Poloniex_sbdbtc                   Summaries_sub4  `json:"poloniex:sbdbtc"`
	Poloniex_scbtc                    Summaries_sub31 `json:"poloniex:scbtc"`
	Poloniex_sdcbtc                   Summaries_sub24 `json:"poloniex:sdcbtc"`
	Poloniex_sjcxbtc                  Summaries_sub24 `json:"poloniex:sjcxbtc"`
	Poloniex_steembtc                 Summaries_sub45 `json:"poloniex:steembtc"`
	Poloniex_steemeth                 Summaries_sub7  `json:"poloniex:steemeth"`
	Poloniex_storjbtc                 Summaries_sub31 `json:"poloniex:storjbtc"`
	Poloniex_stratbtc                 Summaries_sub4  `json:"poloniex:stratbtc"`
	Poloniex_strbtc                   Summaries_sub31 `json:"poloniex:strbtc"`
	Poloniex_strusdt                  Summaries_sub26 `json:"poloniex:strusdt"`
	Poloniex_sysbtc                   Summaries_sub12 `json:"poloniex:sysbtc"`
	Poloniex_viabtc                   Summaries_sub4  `json:"poloniex:viabtc"`
	Poloniex_vrcbtc                   Summaries_sub31 `json:"poloniex:vrcbtc"`
	Poloniex_vtcbtc                   Summaries_sub4  `json:"poloniex:vtcbtc"`
	Poloniex_xbcbtc                   Summaries_sub4  `json:"poloniex:xbcbtc"`
	Poloniex_xcpbtc                   Summaries_sub4  `json:"poloniex:xcpbtc"`
	Poloniex_xembtc                   Summaries_sub31 `json:"poloniex:xembtc"`
	Poloniex_xmrbtc                   Summaries_sub7  `json:"poloniex:xmrbtc"`
	Poloniex_xmrusdt                  Summaries_sub30 `json:"poloniex:xmrusdt"`
	Poloniex_xpmbtc                   Summaries_sub12 `json:"poloniex:xpmbtc"`
	Poloniex_xrpbtc                   Summaries_sub31 `json:"poloniex:xrpbtc"`
	Poloniex_xrpusdt                  Summaries_sub26 `json:"poloniex:xrpusdt"`
	Poloniex_xvcbtc                   Summaries_sub12 `json:"poloniex:xvcbtc"`
	Poloniex_zecbtc                   Summaries_sub7  `json:"poloniex:zecbtc"`
	Poloniex_zeceth                   Summaries_sub7  `json:"poloniex:zeceth"`
	Poloniex_zecusdt                  Summaries_sub38 `json:"poloniex:zecusdt"`
	Poloniex_zecxmr                   Summaries_sub7  `json:"poloniex:zecxmr"`
	Poloniex_zrxbtc                   Summaries_sub12 `json:"poloniex:zrxbtc"`
	Poloniex_zrxeth                   Summaries_sub7  `json:"poloniex:zrxeth"`
	Qryptos_etcbtc                    Summaries_sub10 `json:"qryptos:etcbtc"`
	Qryptos_ethbtc                    Summaries_sub7  `json:"qryptos:ethbtc"`
	Qryptos_ltcbtc                    Summaries_sub7  `json:"qryptos:ltcbtc"`
	Qryptos_repbtc                    Summaries_sub24 `json:"qryptos:repbtc"`
	Qryptos_xmrbtc                    Summaries_sub10 `json:"qryptos:xmrbtc"`
	Qryptos_xrpbtc                    Summaries_sub12 `json:"qryptos:xrpbtc"`
	Qryptos_zecbtc                    Summaries_sub24 `json:"qryptos:zecbtc"`
	Quoine_bchjpy                     Summaries_sub30 `json:"quoine:bchjpy"`
	Quoine_bchsgd                     Summaries_sub30 `json:"quoine:bchsgd"`
	Quoine_bchusd                     Summaries_sub7  `json:"quoine:bchusd"`
	Quoine_btcaud                     Summaries_sub7  `json:"quoine:btcaud"`
	Quoine_btccny                     Summaries_sub24 `json:"quoine:btccny"`
	Quoine_btceur                     Summaries_sub7  `json:"quoine:btceur"`
	Quoine_btchkd                     Summaries_sub22 `json:"quoine:btchkd"`
	Quoine_btcidr                     Summaries_sub12 `json:"quoine:btcidr"`
	Quoine_btcinr                     Summaries_sub24 `json:"quoine:btcinr"`
	Quoine_btcjpy                     Summaries_sub38 `json:"quoine:btcjpy"`
	Quoine_btcphp                     Summaries_sub7  `json:"quoine:btcphp"`
	Quoine_btcsgd                     Summaries_sub7  `json:"quoine:btcsgd"`
	Quoine_btcusd                     Summaries_sub7  `json:"quoine:btcusd"`
	Quoine_ethaud                     Summaries_sub24 `json:"quoine:ethaud"`
	Quoine_ethbtc                     Summaries_sub7  `json:"quoine:ethbtc"`
	Quoine_ethcny                     Summaries_sub24 `json:"quoine:ethcny"`
	Quoine_etheur                     Summaries_sub24 `json:"quoine:etheur"`
	Quoine_ethhkd                     Summaries_sub24 `json:"quoine:ethhkd"`
	Quoine_ethidr                     Summaries_sub16 `json:"quoine:ethidr"`
	Quoine_ethinr                     Summaries_sub24 `json:"quoine:ethinr"`
	Quoine_ethjpy                     Summaries_sub30 `json:"quoine:ethjpy"`
	Quoine_ethphp                     Summaries_sub24 `json:"quoine:ethphp"`
	Quoine_ethsgd                     Summaries_sub7  `json:"quoine:ethsgd"`
	Quoine_ethusd                     Summaries_sub38 `json:"quoine:ethusd"`
}

type Summaries_sub3 struct {
	Change Summaries_sub2 `json:"change"`
	High   float64        `json:"high"`
	Last   float64        `json:"last"`
	Low    float64        `json:"low"`
}

type Summaries_sub40 struct {
	Change Summaries_sub2 `json:"change"`
	High   float64        `json:"high"`
	Last   float64          `json:"last"`
	Low    float64        `json:"low"`
}

type Summaries_sub44 struct {
	Change Summaries_sub2 `json:"change"`
	High   float64        `json:"high"`
	Last   float64          `json:"last"`
	Low    float64          `json:"low"`
}

type Summaries_sub32 struct {
	Change Summaries_sub2 `json:"change"`
	High   float64          `json:"high"`
	Last   float64        `json:"last"`
	Low    float64        `json:"low"`
}

type Summaries_sub35 struct {
	Change Summaries_sub2 `json:"change"`
	High   float64          `json:"high"`
	Last   float64        `json:"last"`
	Low    float64          `json:"low"`
}

type Summaries_sub21 struct {
	Change Summaries_sub2 `json:"change"`
	High   float64          `json:"high"`
	Last   float64          `json:"last"`
	Low    float64        `json:"low"`
}

type Summaries_sub11 struct {
	Change Summaries_sub2 `json:"change"`
	High   float64          `json:"high"`
	Last   float64          `json:"last"`
	Low    float64          `json:"low"`
}

type Summaries_sub6 struct {
	Change Summaries_sub5 `json:"change"`
	High   float64        `json:"high"`
	Last   float64        `json:"last"`
	Low    float64        `json:"low"`
}

type Summaries_sub19 struct {
	Change Summaries_sub5 `json:"change"`
	High   float64        `json:"high"`
	Last   float64        `json:"last"`
	Low    float64          `json:"low"`
}

type Summaries_sub37 struct {
	Change Summaries_sub5 `json:"change"`
	High   float64        `json:"high"`
	Last   float64          `json:"last"`
	Low    float64        `json:"low"`
}

type Summaries_sub13 struct {
	Change Summaries_sub5 `json:"change"`
	High   float64        `json:"high"`
	Last   float64          `json:"last"`
	Low    float64          `json:"low"`
}

type Summaries_sub29 struct {
	Change Summaries_sub5 `json:"change"`
	High   float64          `json:"high"`
	Last   float64        `json:"last"`
	Low    float64        `json:"low"`
}

type Summaries_sub17 struct {
	Change Summaries_sub5 `json:"change"`
	High   float64          `json:"high"`
	Last   float64        `json:"last"`
	Low    float64          `json:"low"`
}

type Summaries_sub27 struct {
	Change Summaries_sub5 `json:"change"`
	High   float64          `json:"high"`
	Last   float64          `json:"last"`
	Low    float64        `json:"low"`
}

type Summaries_sub15 struct {
	Change Summaries_sub5 `json:"change"`
	High   float64          `json:"high"`
	Last   float64          `json:"last"`
	Low    float64          `json:"low"`
}

type Summaries_sub9 struct {
	Change Summaries_sub8 `json:"change"`
	High   float64        `json:"high"`
	Last   float64        `json:"last"`
	Low    float64        `json:"low"`
}

type Summaries_sub23 struct {
	Change Summaries_sub8 `json:"change"`
	High   float64          `json:"high"`
	Last   float64          `json:"last"`
	Low    float64          `json:"low"`
}

type Summaries_sub1 struct {
	Cost      float64 `json:"cost"`
	Remaining float64 `json:"remaining"`
}

type Summaries_sub12 struct {
	Price  Summaries_sub11 `json:"price"`
	Volume float64         `json:"volume"`
}

type Summaries_sub31 struct {
	Price  Summaries_sub11 `json:"price"`
	Volume float64           `json:"volume"`
}

type Summaries_sub14 struct {
	Price  Summaries_sub13 `json:"price"`
	Volume float64         `json:"volume"`
}

type Summaries_sub42 struct {
	Price  Summaries_sub13 `json:"price"`
	Volume float64           `json:"volume"`
}

type Summaries_sub16 struct {
	Price  Summaries_sub15 `json:"price"`
	Volume float64         `json:"volume"`
}

type Summaries_sub39 struct {
	Price  Summaries_sub15 `json:"price"`
	Volume float64           `json:"volume"`
}

type Summaries_sub18 struct {
	Price  Summaries_sub17 `json:"price"`
	Volume float64         `json:"volume"`
}

type Summaries_sub20 struct {
	Price  Summaries_sub19 `json:"price"`
	Volume float64         `json:"volume"`
}

type Summaries_sub22 struct {
	Price  Summaries_sub21 `json:"price"`
	Volume float64         `json:"volume"`
}

type Summaries_sub24 struct {
	Price  Summaries_sub23 `json:"price"`
	Volume float64           `json:"volume"`
}

type Summaries_sub28 struct {
	Price  Summaries_sub27 `json:"price"`
	Volume float64         `json:"volume"`
}

type Summaries_sub30 struct {
	Price  Summaries_sub29 `json:"price"`
	Volume float64         `json:"volume"`
}

type Summaries_sub34 struct {
	Price  Summaries_sub29 `json:"price"`
	Volume float64           `json:"volume"`
}

type Summaries_sub4 struct {
	Price  Summaries_sub3 `json:"price"`
	Volume float64        `json:"volume"`
}

type Summaries_sub25 struct {
	Price  Summaries_sub3 `json:"price"`
	Volume float64          `json:"volume"`
}

type Summaries_sub43 struct {
	Price  Summaries_sub32 `json:"price"`
	Volume float64         `json:"volume"`
}

type Summaries_sub33 struct {
	Price  Summaries_sub32 `json:"price"`
	Volume float64           `json:"volume"`
}

type Summaries_sub36 struct {
	Price  Summaries_sub35 `json:"price"`
	Volume float64           `json:"volume"`
}

type Summaries_sub38 struct {
	Price  Summaries_sub37 `json:"price"`
	Volume float64         `json:"volume"`
}

type Summaries_sub41 struct {
	Price  Summaries_sub40 `json:"price"`
	Volume float64         `json:"volume"`
}

type Summaries_sub45 struct {
	Price  Summaries_sub44 `json:"price"`
	Volume float64         `json:"volume"`
}

type Summaries_sub7 struct {
	Price  Summaries_sub6 `json:"price"`
	Volume float64        `json:"volume"`
}

type Summaries_sub26 struct {
	Price  Summaries_sub6 `json:"price"`
	Volume float64          `json:"volume"`
}

type Summaries_sub10 struct {
	Price  Summaries_sub9 `json:"price"`
	Volume float64        `json:"volume"`
}
