package main

import (
	"bytes"
	"encoding/json"
	"go-cryptowatch/models"
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
	var assets models.Assets
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
	var asset models.Asset
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
