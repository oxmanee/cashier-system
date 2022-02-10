package validator

import (
	"cashier/model"
	"math"

	validate "gopkg.in/go-playground/validator.v9"
)

func CheckRequest(request model.CalculateRequest) (err error) {
	valid := validate.New()

	valid.RegisterValidation("receivePriceInvalid", func(fl validate.FieldLevel) bool {
		if fl.Field().Float() < float64(request.ProductPrice) {
			return false
		}
		return true
	})

	valid.RegisterValidation("negativePrice", func(fl validate.FieldLevel) bool {
		if math.Signbit(fl.Field().Float()) {
			return false
		}
		return true
	})

	valid.RegisterValidation("formatPrice", func(fl validate.FieldLevel) bool {
		_, decimal := math.Modf(fl.Field().Float())
		if decimal != 0.5 && decimal != 0 && decimal != 0.75 && decimal != 0.25 {
			return false
		}

		return true
	})

	err = valid.Struct(request)

	if err != nil {
		return
	}

	return
}
