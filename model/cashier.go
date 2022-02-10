package model

type CalculateRequest struct {
	ProductPrice float64 `json:"productPrice" validate:"required,negativePrice,formatPrice"`
	ReceivePrice float64 `json:"receivePrice" validate:"required,negativePrice,formatPrice,receivePriceInvalid"`
}

type CalculateResponse struct {
	Message string  `json:"message,omitempty"`
	Amount  *Amount `json:"changeAmount,omitempty"`
	Change  float64 `json:"change,omitempty"`
}

type AmountResponse struct {
	Amount Amount `json:"amount"`
}

type Amount struct {
	OneThoundsandBank    int `json:"oneThoundsandBank"`
	FiveHundredBank      int `json:"fiveHundredBank"`
	OneHundredBank       int `json:"oneHundredBank"`
	FiftyBank            int `json:"fiftyBank"`
	TwentyBank           int `json:"twentyBank"`
	TenCoin              int `json:"tenCoin"`
	FiveCoin             int `json:"fiveCoin"`
	OneCoin              int `json:"oneCoin"`
	TwentyFiveSatangCoin int `json:"twentyFiveSatangCoin"`
}
