package cli_test

import (
	"bufio"
	"roc-gameshop-app/cli"
	"strings"
	"testing"
)

func TestCartCli(t *testing.T) {
	in := bufio.NewReader(strings.NewReader("1\n"))
	cartCli := cli.NewCartCli(NewMockRouter(), in)
	session := cli.NewSession()
	session.CurrentCart.AddItem(&cli.CartItem{
		Game:      cli.NewTestGame(1),
		BuyOrRent: "Buy",
		Qty:       1,
	})
	session.CurrentCart.AddItem(&cli.CartItem{
		Game:      cli.NewTestGame(2),
		BuyOrRent: "Rent",
		RentDays:  2,
	})
	session.CurrentCart.AddItem(&cli.CartItem{
		Game:      cli.NewTestGame(3),
		BuyOrRent: "Buy",
		Qty:       10,
	})

	cartCli.HandleRoute(cli.RouteArgs{}, session)
}
