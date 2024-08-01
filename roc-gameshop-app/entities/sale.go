package entities

import "time"

type Sale struct {
	SaleId         int
	GameId         int
	UserId         int
	SaleDate       time.Time
	PurchasedPrice float64
	Quantity       int
}
