package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/handlers"
	"strings"

	"github.com/rodaine/table"
)

type gamesCli struct {
	gamesHandler handlers.GamesHandler
	router Router
	reader *bufio.Reader
}

func NewGamesCli(router Router, reader *bufio.Reader, gamesHandler handlers.GamesHandler) Cli {
	return &gamesCli{
		router: router,
		reader: reader,
		gamesHandler: gamesHandler,
	}
}

func (gC *gamesCli) HandleRoute(args RouteArgs, session *Session) {
	//get game name to search
	fmt.Printf("Enter game name to search: ")
	name, err := gC.reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading game name input")
	}
	name = strings.TrimSpace(name)

	games, err := gC.gamesHandler.GetAll(name, 10, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Search Results:")
	gamesTable := table.New("No", "Name", "Rating", "Sale Price", "Stock")
	for i, game := range games {
		gamesTable.AddRow(i + 1, game.Name, game.Genre, game.SalePrice, game.Stock).WithPadding(1)
	}
	gamesTable.Print()

	actions:= []string{
		"View Game",
		"Search Again",
		"View Cart",
		"Back",
	}
	fmt.Println("")
	fmt.Println("Actions:")
	for i, action := range actions{
		fmt.Printf("%d. %s\n", i + 1, action)
	}
	fmt.Printf("Enter your action: ")
	action, err := gC.reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading action input")
	}
	action = strings.TrimSpace(action)
	fmt.Println(action)
}