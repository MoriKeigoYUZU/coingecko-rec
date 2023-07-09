package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CoinData struct {
	MarketData struct {
		CurrentPrice struct {
			Jpy float64 `json:"jpy"`
			Usd float64 `json:"usd"`
		} `json:"current_price"`
	} `json:"market_data"`
}

func main() {
	apiID := "axie-infinity"

	url := `https://api.coingecko.com/api/v3/coins/` + apiID + `?localization=false&tickers=false&market_data=true&community_data=false&developer_data=false&sparkline=false`

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Request Error : %s\n", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Response Error : %s\n", err)
		return
	}

	var coinData CoinData
	err = json.Unmarshal(body, &coinData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("USD : %f\n", coinData.MarketData.CurrentPrice.Usd)
	fmt.Printf("JPY : %f\n", coinData.MarketData.CurrentPrice.Jpy)
}
