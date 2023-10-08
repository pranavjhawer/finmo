package handlers

import (
	"finmo/db"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type resBalance struct {
	Balances map[string]float64 `json:"balances"`
}

func Balance(c echo.Context) error {
	accounID := c.Param("account_id")
	wallet, isExist := db.GetWallet(accounID)
	if !isExist {
		return echo.NewHTTPError(http.StatusBadRequest, "User wallet not found")
	}
	res := resBalance{
		Balances: wallet,
	}
	return c.JSON(http.StatusCreated, res)
}
