package handlers

import (
	"encoding/json"
	"finmo/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type reqTopup struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

func TopUp(c echo.Context) error {
	accounID := c.Param("account_id")
	var req reqTopup
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	db.TopUpWallet(accounID, req.Currency, req.Amount)
	return c.JSON(http.StatusCreated, req)
}
