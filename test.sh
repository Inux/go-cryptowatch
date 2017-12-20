#!/bin/bash

( cd m/types; stringer -type=PairType; stringer -type=CurrencyType; stringer -type=MarketType )
# go test ./... -v 