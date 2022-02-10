package http

import (
	"net/http"

	"cashier/model"
	usecase "cashier/usecase/interface"

	"github.com/labstack/echo/v4"

	validate "cashier/validator"
)

type cashierDelivery struct {
	cashierUsecase usecase.CashierUsecaseInterface
}

func NewCashierDelivery(e *echo.Echo, cashierUsecase usecase.CashierUsecaseInterface) {
	handler := &cashierDelivery{
		cashierUsecase: cashierUsecase,
	}

	cashierRoute := e.Group("/cashier")
	cashierRoute.GET("/amount", handler.Amount)
	cashierRoute.POST("/calculate", handler.Calculate)
}

func (handler *cashierDelivery) Calculate(c echo.Context) error {

	var response model.CalculateResponse

	request := new(model.CalculateRequest)

	err := c.Bind(request)
	if err != nil {
		return err
	}

	err = validate.CheckRequest(*request)
	if err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	balance := handler.cashierUsecase.GetBalance()
	change := handler.cashierUsecase.GetChange(*request)

	if balance < change {
		response.Message = "Sorry, Not enough money in the system."
		return c.JSON(http.StatusBadRequest, response)
	}

	amount, over := handler.cashierUsecase.Calculate(*request)
	if over {
		response.Message = "Sorry, Not enough money in the system."
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Amount = &amount
	response.Change = change

	return c.JSON(http.StatusOK, response)
}

func (handler *cashierDelivery) Amount(c echo.Context) error {
	var response model.AmountResponse

	response.Amount = handler.cashierUsecase.GetAmount()

	return c.JSON(http.StatusOK, response)
}
