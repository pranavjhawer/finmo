package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type alphavantageResponse struct {
	RealtimeCurrencyExchangeRate struct {
		FromCurrencyCode string `json:"1. From_Currency Code"`
		FromCurrencyName string `json:"2. From_Currency Name"`
		ToCurrencyCode   string `json:"3. To_Currency Code"`
		ToCurrencyName   string `json:"4. To_Currency Name"`
		ExchangeRate     string `json:"5. Exchange Rate"`
		LastRefreshed    string `json:"6. Last Refreshed"`
		TimeZone         string `json:"7. Time Zone"`
		BidPrice         string `json:"8. Bid Price"`
		AskPrice         string `json:"9. Ask Price"`
	} `json:"Realtime Currency Exchange Rate"`
}

func alphavantage(from string, to string) (float64, error) {
	// Fetch conversion rate from alpha	vantage
	// https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=USD&to_currency=SGD&apikey=Z7GQ6EMA24HXJBCV
	URL := alphaVantageURL + "&from_currency=" + from + "&to_currency=" + to + "&apikey=Z7GQ6EMA24HXJBCV"
	resp, err := http.Get(URL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, echo.ErrBadGateway
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var alphavantageResponse alphavantageResponse
	err = json.Unmarshal(bs, &alphavantageResponse)
	if err != nil {
		return 0, err
	}
	b, _ := json.Marshal(alphavantageResponse)
	fmt.Println(string(b))
	rate := alphavantageResponse.RealtimeCurrencyExchangeRate.ExchangeRate

	// Convert the string to a float32
	convertedRate, err := strconv.ParseFloat(rate, 32)
	if err != nil {
		return 0, err
	}

	return convertedRate, nil
}
