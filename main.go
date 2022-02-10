package main

import (
	"cashier/delivery/http"
	"cashier/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	cashierUsecase := usecase.NewCashierUsecase()

	e := echo.New()
	http.NewCashierDelivery(e, cashierUsecase)

	e.Logger.Fatal(e.Start(":80"))
}
