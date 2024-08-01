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
	Items []*CartItem
}

func (c *Cart) AddItem(item *CartItem) {
	c.Items = append(c.Items, item)
}

func (c *Cart) RemoveItem(index int) *CartItem {
	if index > len(c.Items)-1 || index < 0 {
		return nil
	}
	// get remove to be removed
	removed := c.Items[index]
	// remove item
	c.Items = append(c.Items[:index], c.Items[index+1:]...)
	return removed
}
