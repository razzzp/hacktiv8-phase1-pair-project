package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/handlers"
	"roc-gameshop-app/routes"
	"strings"
)

type gameDetailsCli struct {
	gameHandler handlers.GamesHandler
	router      Router
	reader      *bufio.Reader
}

func NewGameDetailsCli(router Router, reader *bufio.Reader, gameHandler handlers.GamesHandler) Cli {
	return &gameDetailsCli{
		router:      router,
		reader:      reader,
		gameHandler: gameHandler,
	}
}

func (gDC *gameDetailsCli) HandleRoute(args RouteArgs, session *Session) {
	// logic of game details page goes here
	// TODO
	fmt.Println("Game Details Page")
	fmt.Println("Game ID: ", args["gameId"])

	fmt.Println("Actions")
	fmt.Println("1. Go to Game 2")
	fmt.Println("2. Back")

	// temp for testing
	for {
		fmt.Print("What would you like to do? ")
		input, _ := gDC.reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "1" {
			// push to router and return to go to another route
			gDC.router.Push(routes.GAME_DETAILS_ROUTE, RouteArgs{"gameId": "2"})
			return
		} else if input == "2" {
			// pop and return to return to previous route
			gDC.router.Pop()
			return
		} else {
			fmt.Println("Invalid Action.")
		}
	}

}
