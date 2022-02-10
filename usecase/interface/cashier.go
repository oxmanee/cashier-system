package usecase

import "cashier/model"

type CashierUsecaseInterface interface {
	GetAmount() (amount model.Amount)
	Calculate(request model.CalculateRequest) (amount model.Amount, over bool)
	GetChange(request model.CalculateRequest) float64
	GetBalance() float64
}
