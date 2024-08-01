package cli

import (
	"roc-gameshop-app/entities"
)

type CartItem struct {
	Game      *entities.Game
	Qty       int
	BuyOrRent string
	RentDays  int
}

type Cart struct {
	items []*CartItem
}
