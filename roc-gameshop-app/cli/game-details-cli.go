package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/handlers"
	"strconv"
)

type gameDetailsCli struct {
	gameHandler   handlers.GamesHandler
	reviewHandler handlers.ReviewHandler
	router        Router
	reader        *bufio.Reader
}

func NewGameDetailsCli(router Router, reader *bufio.Reader, gameHandler handlers.GamesHandler, reviewHandler handlers.ReviewHandler) Cli {
	return &gameDetailsCli{
		router:        router,
		reader:        reader,
		gameHandler:   gameHandler,
		reviewHandler: reviewHandler,
	}
}

func (gDC *gameDetailsCli) HandleRoute(args RouteArgs, session *Session) {
	id, err := strconv.Atoi(args["gameId"])
	if err != nil {
		fmt.Println("Error converting gameId ", err)
		return
	}
	game, err := gDC.gameHandler.GetById(id)
	if err != nil {
		fmt.Println(err)
	}
	reviews, err := gDC.reviewHandler.GetByGameId(id)
	if err != nil {
		fmt.Println(err)
	}
	// logic of game details page goes here
	// TODO
	fmt.Println("Game Details Page")
	fmt.Println("")
	fmt.Printf("Genre: %s\n", game.Genre)
	fmt.Printf("Description: %s\n", game.Description)
	fmt.Println("Rating: ...")
	fmt.Printf("Sale Price: %.2f\n", game.SalePrice)
	fmt.Printf("Rent Price: %.2f/day\n", game.RentalPrice)
	fmt.Println("Reviews:")
	fmt.Println("")
	for i, review := range reviews {
		fmt.Printf("%d. %s     %.2f\n", i+1, review.UserName, review.Rating)
		fmt.Println(review.ReviewMsg)
	}
	fmt.Println("")
	actions := []string{
		"Buy",
		"Rent",
		"View Cart",
		"Add Review",
		"Back",
	}
	for i, v := range actions {
		fmt.Printf("%d. %s\n", i+1, v)
	}

	// // temp for testing

	// fmt.Println("Actions")
	// fmt.Println("1. Go to Game 2")
	// fmt.Println("2. Back")
	// for {
	// 	fmt.Print("What would you like to do? ")
	// 	input, _ := gDC.reader.ReadString('\n')
	// 	input = strings.TrimSpace(input)
	// 	if input == "1" {
	// 		// push to router and return to go to another route
	// 		gDC.router.Push(routes.GAME_DETAILS_ROUTE, RouteArgs{"gameId": "2"})
	// 		return
	// 	} else if input == "2" {
	// 		// pop and return to return to previous route
	// 		gDC.router.Pop()
	// 		return
	// 	} else {
	// 		fmt.Println("Invalid Action.")
	// 	}
	// }

}
