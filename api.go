package main

import (
	"fmt"
	"go-cryptowatch/models"
)

type Change struct {
}

type Summary struct {
	High float64
}

type CurrencySummary struct {
	Currency currency
	Summary
}

type currency struct {
	Name    string
	Summary models.Summary
}

func GetBitfinexIndex() {

}

func main() {
	fmt.Println("for go build...")
}
