package main

import (
	"finmo/db"
	"finmo/handlers"

	"github.com/labstack/echo/v4"
)

// TODO
// - Add validation for params

// v1/account/:account_id
func main() {
	e := echo.New()
	db.DBInit()
	// POST /topup
	e.POST("v1/account/:account_id/topup", func(c echo.Context) error {
		return handlers.TopUp(c)
	})

	// GET /fx-rates
	e.GET("v1/account/:account_id/fx-rates", func(c echo.Context) error {
		return handlers.FxRate(c)
	})

	// POST /fx-conversion
	e.POST("v1/account/:account_id/fx-conversion", func(c echo.Context) error {
		return handlers.FxConversion(c)
	})

	// GET /balance
	e.GET("v1/account/:account_id/balance", func(c echo.Context) error {
		return handlers.Balance(c)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
