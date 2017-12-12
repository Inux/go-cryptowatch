#!/bin/bash

( cd models; stringer -type=CurrencyType; stringer -type=PairType )
go test ./... -v 