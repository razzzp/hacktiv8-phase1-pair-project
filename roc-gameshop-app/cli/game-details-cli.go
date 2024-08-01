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
	rating, err := gDC.reviewHandler.GetAvgRating(id)
	if err != nil {
		fmt.Println(err)
	}
	// logic of game details page goes here
	// TODO
	fmt.Println("Game Details Page")
	fmt.Println(game.Name)
	fmt.Println("")
	fmt.Printf("Genre: %s\n", game.Genre)
	fmt.Printf("Description: %s\n", game.Description)
	fmt.Printf("Average Rating: %.2f\n", *rating)
	fmt.Printf("Sale Price: %.2f\n", game.SalePrice)
	fmt.Printf("Rent Price: %.2f/day\n", game.RentalPrice)
	fmt.Println("Reviews:")
	fmt.Println("")
	for i, review := range reviews {
		fmt.Printf("%d. %s     Rating: %.2f\n", i+1, review.UserName, review.Rating)
		fmt.Println(review.ReviewMsg)
	}
	fmt.Println("")
	actions := []Action{
		{
			Name: "Buy",
			ActionFunc: func() {
				//TODO
			},
		},
		{
			Name: "Rent",
			ActionFunc: func() {
				//TODO
			},
		},
		{
			Name: "View Cart",
			ActionFunc: func() {
				//TODO
			},
		},
		{
			Name: "Add Review",
			ActionFunc: func() {
				//TODO
			},
		},
		{
			Name: "Back",
			ActionFunc: func() {
				gDC.router.Pop()
			},
		},
	}
	PromptUserForActions(actions, gDC.reader)

}
