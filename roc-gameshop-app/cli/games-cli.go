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

func (gC *gamesCli) GetActions(session *Session) []Action {
	result := []Action{}
	result = append(result, Action{Name: "View Game", ActionFunc: func() {
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
			return
		}
		gC.router.Push(routes.GAME_DETAILS_ROUTE, RouteArgs{"gameId": gameId})

	}})

	result = append(result, Action{Name: "Search Again", ActionFunc: func() {
		//get game name to search
		fmt.Printf("Enter game name to search: ")
		name, err := gC.reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading game name input")
			return
		}
		name = strings.TrimSpace(name)
		gC.router.Push(routes.GAMES_ROUTE, RouteArgs{"gameName": name})

	}})

	if session.CurrentUser == nil || !session.CurrentUser.IsAdmin() {
		result = append(result, Action{Name: "View Cart", ActionFunc: func() {
			gC.router.Push(routes.CART_ROUTE, RouteArgs{})
		}})
	}

	result = append(result, Action{Name: "Back", ActionFunc: func() {
		gC.router.Pop()
	}})

	return result
}

func (gC *gamesCli) HandleRoute(args RouteArgs, session *Session) {

	games, err := gC.gamesHandler.GetAll(args["gameName"], 10, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Search Results")
	fmt.Println("")
	gamesTable := table.New("No", "Name", "Rating", "Sale Price", "Stock")
	for _, game := range games {
		gamesTable.AddRow(game.GameId, game.Name, game.Genre, game.SalePrice, game.Stock).WithPadding(1)
	}
	gamesTable.Print()

	fmt.Println("")

	actions := gC.GetActions(session)
	PromptUserForActions(actions, gC.reader)
}
