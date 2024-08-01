package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/entities"
	"roc-gameshop-app/routes"

	"github.com/rodaine/table"
)

type cartCli struct {
	router Router
	reader *bufio.Reader
}

func NewCartCli(router Router, reader *bufio.Reader) Cli {
	return &cartCli{
		router: router,
		reader: reader,
	}
}

func (cc *cartCli) checkout(session *Session) {
	// TODO
	if session.CurrentUser == nil {
		// have to login first
		cc.router.Push(routes.LOGIN_ROUTE, RouteArgs{})
		return
	}

}

func (cc *cartCli) GetUserActions(session *Session) []Action {
	result := []Action{}
	// checkout cart
	result = append(result, Action{Name: "Check Out", ActionFunc: func() {
		cc.checkout(session)
	}})
	// remove item from cart
	result = append(result, Action{Name: "Remove Item", ActionFunc: func() {
		input, err := PromptUserForInt("Enter item number to remove: ", cc.reader)
		if err != nil {
			// refresh
			cc.router.Push(routes.CART_ROUTE, RouteArgs{})
			return
		}
		session.CurrentCart.RemoveItem(input - 1)
		cc.router.Push(routes.CART_ROUTE, RouteArgs{})
	}})
	// go back
	result = append(result, Action{Name: "Back", ActionFunc: func() {
		cc.router.Pop()
	}})
	return result
}

func getSubtotalAsString(cartItem *CartItem) string {
	if cartItem.BuyOrRent == "Buy" {
		return fmt.Sprintf("Rp. %.0f", cartItem.Game.SalePrice*float64(cartItem.Qty))
	} else {
		return fmt.Sprintf("Rp. %.0f", cartItem.Game.RentalPrice*float64(cartItem.RentDays))
	}
}

func (cc *cartCli) HandleRoute(args RouteArgs, session *Session) {

	var name string
	if session.CurrentUser == nil {
		name = "Your"
	} else {
		name = fmt.Sprintf("%s's", session.CurrentUser.Name)
	}
	fmt.Printf("%s Cart\n", name)
	fmt.Println("")
	cartTable := table.New("No.", "Game", "Buy/Rent", "Qty/RentDays", "Subtotal")
	for i, cartItem := range session.CurrentCart.Items {
		if cartItem.BuyOrRent == "Buy" {
			cartTable.AddRow(i+1, cartItem.Game.Name, cartItem.BuyOrRent, cartItem.Qty, getSubtotalAsString(cartItem)).WithPadding(1)
		} else {
			cartTable.AddRow(i+1, cartItem.Game.Name, cartItem.BuyOrRent, cartItem.RentDays, getSubtotalAsString(cartItem)).WithPadding(1)
		}
	}
	cartTable.Print()

	fmt.Println("")

	// get user actions
	actions := cc.GetUserActions(session)
	PromptUserForActions(actions, cc.reader)
}

func NewTestGame(withId int) *entities.Game {
	return &entities.Game{
		GameId:      withId,
		Name:        fmt.Sprintf("Test Game %d", withId),
		Description: "Test desc",
		Genre:       "Test Genre",
		SalePrice:   10_000,
		RentalPrice: 5_000,
		Studio:      "Test studio",
		Stock:       10,
	}
}
