#!/bin/bash

declare -a URLs=(
    'https://api.cryptowat.ch/assets'
    'https://api.cryptowat.ch/assets/btc'
    'https://api.cryptowat.ch/pairs'
    'https://api.cryptowat.ch/pairs/ethbtc'
    'https://api.cryptowat.ch/exchanges'
    'https://api.cryptowat.ch/exchanges/kraken'
    'https://api.cryptowat.ch/markets'
    'https://api.cryptowat.ch/markets/gdax/btcusd'
    'https://api.cryptowat.ch/markets/gdax/btcusd/price'
    'https://api.cryptowat.ch/markets/gdax/btcusd/summary'
    'https://api.cryptowat.ch/markets/gdax/btcusd/trades'
    'https://api.cryptowat.ch/markets/gdax/btcusd/orderbook'
    'https://api.cryptowat.ch/markets/gdax/btcusd/ohlc'
    'https://api.cryptowat.ch/markets/prices'
    'https://api.cryptowat.ch/markets/summaries'
)

declare -a Names=(
    'Assets'
    'Asset'
    'Pairs'
    'Pair'
    'Exchanges'
    'Exchange'
    'Markets'
    'Market'
    'Price'
    'Summary'
    'Trades'
    'OrderBook'
    'OHLC'
    'Prices'
    'Summaries'
)

i=0
for name in "${Names[@]}"
do
    filename=$(echo "$name" | tr '[:upper:]' '[:lower:]')
    echo 'fetch from '${URLs[$i]}
    echo 'and writing '$name' into models/gen_'$filename'.go'
    curl -s ${URLs[$i]} | gojson -name $name -o models/gen_$filename.go -pkg models -forcefloats
    i=$((i + 1))
done
