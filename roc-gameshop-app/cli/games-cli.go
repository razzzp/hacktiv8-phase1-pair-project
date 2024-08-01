package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/handlers"
	"roc-gameshop-app/routes"
	"strconv"
	"strings"

	"github.com/rodaine/table"
)

type gamesCli struct {
	gamesHandler handlers.GamesHandler
	router       Router
	reader       *bufio.Reader
}

func NewGamesCli(router Router, reader *bufio.Reader, gamesHandler handlers.GamesHandler) Cli {
	return &gamesCli{
		router:       router,
		reader:       reader,
		gamesHandler: gamesHandler,
	}
}

func (gC *gamesCli) HandleRoute(args RouteArgs, session *Session) {
gameLoop:
	for {

		games, err := gC.gamesHandler.GetAll(args["gameName"], 10, 1)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Search Results:")
		gamesTable := table.New("No", "Name", "Rating", "Sale Price", "Stock")
		for _, game := range games {
			gamesTable.AddRow(game.GameId, game.Name, game.Genre, game.SalePrice, game.Stock).WithPadding(1)
		}
		gamesTable.Print()

		actions := []string{
			"View Game",
			"Search Again",
			"View Cart",
			"Back",
		}
		fmt.Println("")
		fmt.Println("Actions:")
		for i, action := range actions {
			fmt.Printf("%d. %s\n", i+1, action)
		}
		fmt.Printf("Enter your action: ")
		action, err := gC.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading action input")
		}
		action = strings.TrimSpace(action)
		switch action {
		case "1":
			//TODO
			fmt.Printf("Enter a game you want to view: ")
			gameId, err := gC.reader.ReadString('\n')
			if err != nil {
				fmt.Println("error reading gameId input")
			}
			gameId = strings.TrimSpace(gameId)

			//check if it is valid int
			_, err = strconv.Atoi(gameId)
			if err != nil {
				fmt.Println("Please enter a valid int")
				continue gameLoop
			}
			gC.router.Push(routes.GAME_DETAILS_ROUTE, RouteArgs{"gameId": gameId})
			return
		case "2":
			continue
		case "3":
			gC.router.Push(routes.CART_ROUTE, RouteArgs{})
			return
		case "4":
			gC.router.Pop()
			return
		}
	}
}
