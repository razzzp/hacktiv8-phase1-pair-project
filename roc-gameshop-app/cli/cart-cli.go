package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/entities"
	"roc-gameshop-app/handlers"
	"roc-gameshop-app/routes"
	"strings"
	"time"

	"github.com/rodaine/table"
)

type cartCli struct {
	router      Router
	reader      *bufio.Reader
	salehandler handlers.SaleHandler
}

func NewCartCli(router Router, reader *bufio.Reader, saleHandler handlers.SaleHandler) Cli {
	return &cartCli{
		router:      router,
		reader:      reader,
		salehandler: saleHandler,
	}
}

func (cc *cartCli) processCartItem(cartItem *CartItem, total *float64, session *Session) error {
	if cartItem.BuyOrRent == "Buy" {
		newSale := entities.Sale{}
		newSale.GameId = cartItem.Game.GameId
		newSale.UserId = session.CurrentUser.UserId
		newSale.Quantity = cartItem.Qty
		newSale.SaleDate = time.Now()
		newSale.PurchasedPrice = float64(cartItem.Qty) * cartItem.Game.SalePrice
		*total += newSale.PurchasedPrice
		return cc.salehandler.Create(&newSale)
	} else if cartItem.BuyOrRent == "Rent" {
		// TODO
		return nil
	} else {
		return fmt.Errorf("unknow cart action '%s'", cartItem.BuyOrRent)
	}
}

func (cc *cartCli) checkout(session *Session) error {
	total := 0.0

	for i, cartItem := range session.CurrentCart.Items {
		err := cc.processCartItem(cartItem, &total, session)
		if err != nil {
			fmt.Printf("Failed to process cart item no. %d, for game: %s\n", i+1, cartItem.Game.Name)
			return err
		}
	}
	fmt.Printf("Thank you for your purchase! Total: Rp.%.0f\n", total)
	return nil
}

func (cc *cartCli) GetUserActions(session *Session) []Action {
	result := []Action{}
	// checkout cart
	result = append(result, Action{Name: "Check Out", ActionFunc: func() {
		if session.CurrentUser == nil {
			// have to login first
			cc.router.Push(routes.LOGIN_ROUTE, RouteArgs{})
			return
		} else {
			for {
				err := cc.checkout(session)
				if err != nil {
					fmt.Println(err)

					fmt.Print("Would you like to try again? (y/n): ")
					input, _ := cc.reader.ReadString('\n')
					if strings.EqualFold(input, "y") {
						continue
					}
					// else stay on same page
					break
				} else {
					time.Sleep(time.Second)
					// go back to home
					cc.router.Push(routes.HOME_PAGE_ROUTE, RouteArgs{})
					return
				}
			}
		}
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

// for testing purposes
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
