package handlers

import (
	"encoding/json"
	"finmo/db"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type reqFxConversion struct {
	QuoteID      string  `json:"quoteId"`
	FromCurrency string  `json:"fromCurrency"`
	ToCurrency   string  `json:"toCurrency"`
	Amount       float64 `json:"amount"`
}

type resFxConversion struct {
	ConvertedAmount float64 `json:"convertedAmount"`
	Currency        string  `json:"currency"`
}

func FxConversion(c echo.Context) error {
	accounID := c.Param("account_id")
	var req reqFxConversion
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	conversionRate, err := getConversionRate(req)
	if err != nil {
		return err
	}
	convertedAmount, err := moveAmount(accounID, req, conversionRate)
	if err != nil {
		return err
	}
	res := resFxConversion{
		ConvertedAmount: convertedAmount, //only returning the converted amount NOT the balance
		Currency:        req.ToCurrency,
	}
	return c.JSON(http.StatusCreated, res)
}

func moveAmount(
	accountID string,
	req reqFxConversion,
	conversionRate float64,
) (float64, error) {
	fromCurrency := req.FromCurrency
	toCurrency := req.ToCurrency
	amount := req.Amount
	wallet, isExist := db.GetWallet(accountID)
	if !isExist {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "User wallet not found")
	}
	if wallet[fromCurrency] < amount {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "Insufficient balance")
	}
	wallet[fromCurrency] -= amount

	convertedAmount := conversionRate * amount
	wallet[toCurrency] += convertedAmount

	db.UpdateWallet(accountID, wallet)
	return convertedAmount, nil
}

func getConversionRate(req reqFxConversion) (float64, error) {
	if req.QuoteID == "" {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "Quote ID is required")
	}
	val, isExist := db.GetQuote(req.QuoteID)
	if !isExist {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "Quote ID invalid")
	}
	if val.ExpiryAt.Before(time.Now()) {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "Quote ID expired")
	}
	if val.FromCurrency != req.FromCurrency || val.ToCurrency != req.ToCurrency {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "Quote ID invalid for this currency pair")
	}
	return val.ConversionRate, nil
}
