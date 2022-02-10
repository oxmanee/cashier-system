package usecase

import (
	"cashier/model"
	usecase "cashier/usecase/interface"
	"math"
)

var ONE_THOUSAND_BANK int = 10
var FIVE_HUNDRED_BANK int = 20
var ONE_HUNDRED_BANK int = 15
var FIFTY_BANK int = 20
var TWENTY_BANK int = 30
var TEN_COIN int = 20
var FIVE_COIN int = 20
var ONE_COIN int = 20
var TWENTY_FIVE_SATANG_COIN int = 50

type cashierlUsecase struct {
}

func NewCashierUsecase() usecase.CashierUsecaseInterface {
	return &cashierlUsecase{}
}

func (c cashierlUsecase) GetBalance() float64 {
	return (float64(ONE_THOUSAND_BANK) * 1000) + (float64(FIVE_HUNDRED_BANK) * 500) + (float64(ONE_HUNDRED_BANK) * 100) + (float64(FIFTY_BANK) * 50) + (float64(TWENTY_BANK) * 20) + (float64(TEN_COIN) * 10) + (float64(FIVE_COIN) * 5) + (float64(ONE_COIN)) + (float64(TWENTY_FIVE_SATANG_COIN) * 0.25)
}

func (c cashierlUsecase) GetAmount() (amount model.Amount) {

	amount.OneThoundsandBank = ONE_THOUSAND_BANK
	amount.FiveHundredBank = FIVE_HUNDRED_BANK
	amount.OneHundredBank = ONE_HUNDRED_BANK
	amount.FiftyBank = FIFTY_BANK
	amount.TwentyBank = TWENTY_BANK
	amount.TenCoin = TEN_COIN
	amount.FiveCoin = FIVE_COIN
	amount.OneCoin = ONE_COIN
	amount.TwentyFiveSatangCoin = TWENTY_FIVE_SATANG_COIN

	return
}

func (c cashierlUsecase) Calculate(request model.CalculateRequest) (amount model.Amount, over bool) {

	totalBaht := int(request.ReceivePrice - request.ProductPrice)
	_, totalSatang := math.Modf(request.ReceivePrice - request.ProductPrice)

	if totalBaht >= 1000 && ONE_THOUSAND_BANK != 0 {
		if totalBaht/1000 > ONE_THOUSAND_BANK {
			over := (totalBaht / 1000) - ONE_THOUSAND_BANK
			amount.OneThoundsandBank = ONE_THOUSAND_BANK
			ONE_THOUSAND_BANK = 0
			totalBaht = totalBaht%1000 + (over * 1000)
		} else {
			ONE_THOUSAND_BANK -= totalBaht / 1000
			amount.OneThoundsandBank = totalBaht / 1000
			totalBaht = totalBaht % 1000
		}
	}

	if totalBaht >= 500 && FIVE_HUNDRED_BANK != 0 {
		if totalBaht/500 > FIVE_HUNDRED_BANK {
			over := (totalBaht / 500) - FIVE_HUNDRED_BANK
			amount.FiveHundredBank = FIVE_HUNDRED_BANK
			FIVE_HUNDRED_BANK = 0
			totalBaht = totalBaht%500 + (over * 500)
		} else {
			FIVE_HUNDRED_BANK -= totalBaht / 500
			amount.FiveHundredBank = totalBaht / 500
			totalBaht = totalBaht % 500
		}
	}

	if totalBaht >= 100 && ONE_HUNDRED_BANK != 0 {
		if totalBaht/100 > ONE_HUNDRED_BANK {
			over := (totalBaht / 100) - ONE_HUNDRED_BANK
			amount.OneHundredBank = ONE_HUNDRED_BANK
			ONE_HUNDRED_BANK = 0
			totalBaht = totalBaht%100 + (over * 100)
		} else {
			ONE_HUNDRED_BANK -= totalBaht / 100
			amount.OneHundredBank = totalBaht / 100
			totalBaht = totalBaht % 100
		}
	}

	if totalBaht >= 50 && FIFTY_BANK != 0 {
		if totalBaht/50 > FIFTY_BANK {
			over := (totalBaht / 50) - FIFTY_BANK
			amount.FiftyBank = FIFTY_BANK
			FIFTY_BANK = 0
			totalBaht = totalBaht%50 + (over * 50)
		} else {
			FIFTY_BANK -= totalBaht / 50
			amount.FiftyBank = totalBaht / 50
			totalBaht = totalBaht % 50
		}
	}

	if totalBaht >= 20 && TWENTY_BANK != 0 {
		if totalBaht/20 > TWENTY_BANK {
			over := (totalBaht / 20) - TWENTY_BANK
			amount.TwentyBank = TWENTY_BANK
			TWENTY_BANK = 0
			totalBaht = totalBaht%20 + (over * 20)
		} else {
			TWENTY_BANK -= totalBaht / 20
			amount.TwentyBank = totalBaht / 20
			totalBaht = totalBaht % 20
		}
	}

	if totalBaht >= 10 && TEN_COIN != 0 {
		if totalBaht/10 > TEN_COIN {
			over := (totalBaht / 10) - TEN_COIN
			amount.TenCoin = TEN_COIN
			TEN_COIN = 0
			totalBaht = totalBaht%10 + (over * 10)
		} else {
			TEN_COIN -= totalBaht / 10
			amount.TenCoin = totalBaht / 10
			totalBaht = totalBaht % 10
		}
	}

	if totalBaht >= 5 && FIVE_COIN != 0 {
		if totalBaht/5 > FIVE_COIN {
			over := (totalBaht / 5) - FIVE_COIN
			amount.FiveCoin = FIVE_COIN
			FIVE_COIN = 0
			totalBaht = totalBaht%5 + (over * 5)
		} else {
			FIVE_COIN -= totalBaht / 5
			amount.FiveCoin = totalBaht / 5
			totalBaht = totalBaht % 5
		}
	}

	if totalBaht >= 1 && ONE_COIN != 0 {
		if totalBaht > ONE_COIN {
			over := totalBaht - ONE_COIN
			amount.OneCoin = ONE_COIN
			ONE_COIN = 0
			totalBaht = over
		} else {
			ONE_COIN -= totalBaht
			amount.OneCoin = totalBaht
			totalBaht = 0
		}
	}

	if float64(totalBaht) != 0 || totalSatang > 0 {
		sum := float64(totalBaht) + totalSatang
		if int(sum/0.25) > TWENTY_FIVE_SATANG_COIN {
			over = true
			rollbackAmount(amount)
		} else {
			amount.TwentyFiveSatangCoin = int(sum / 0.25)
			TWENTY_FIVE_SATANG_COIN -= int(sum / 0.25)
		}
	}

	return
}

func (c cashierlUsecase) GetChange(request model.CalculateRequest) float64 {
	return request.ReceivePrice - request.ProductPrice
}

func rollbackAmount(amount model.Amount) {
	ONE_THOUSAND_BANK += amount.OneThoundsandBank
	FIVE_HUNDRED_BANK += amount.FiveHundredBank
	ONE_HUNDRED_BANK += amount.OneHundredBank
	FIFTY_BANK += amount.FiftyBank
	TWENTY_BANK += amount.TwentyBank
	TEN_COIN += amount.TenCoin
	FIVE_COIN += amount.FiveCoin
	ONE_COIN += amount.OneCoin
}
