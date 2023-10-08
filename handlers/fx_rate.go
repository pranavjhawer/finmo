package handlers

import (
	"encoding/json"
	"finmo/db"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type reqFxRate struct {
	FromCurrency string `json:"fromCurrency"`
	ToCurrency   string `json:"toCurrency"`
}

type resFxRate struct {
	QuoteID  string `json:"quoteId"`
	ExpiryAt string `json:"expiry_at"`
}

const (
	alphaVantageURL   = "https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE"
	expiryDurationSec = 20
)

func FxRate(c echo.Context) error {
	var req reqFxRate
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	rateOfConversion, err := alphavantage(req.FromCurrency, req.ToCurrency)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	//generating quote id
	quoteID := uuid.New().String()
	// find current time + 20 seconds
	expiryTime := time.Now().Add(expiryDurationSec * time.Second)
	// store quote id and expiry time in DB
	quoteData := db.QuoteData{
		ConversionRate: rateOfConversion,
		FromCurrency:   req.FromCurrency,
		ToCurrency:     req.ToCurrency,
		ExpiryAt:       expiryTime,
	}
	db.SetQuote(quoteID, quoteData)
	res := resFxRate{
		QuoteID:  quoteID,
		ExpiryAt: fmt.Sprint(expiryTime.Unix()),
	}
	return c.JSON(http.StatusCreated, res)
}
