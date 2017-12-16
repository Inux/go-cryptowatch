package endpoint

import "go-cryptowatch/models"

// Endpoint interface
type Endpoint interface {
	Run()
	Get(models.CurrencyType) interface{}
	GetLast(length int) []interface{}
}
