package cli

import (
	"bufio"
	"fmt"
	"roc-gameshop-app/entities"
	"roc-gameshop-app/handlers"
	"roc-gameshop-app/routes"
	"strconv"
	"strings"
	"time"
)

type gameDetailsCli struct {
	gameHandler   handlers.GamesHandler
	reviewHandler handlers.ReviewHandler
	rentalHandler handlers.RentalHandler
	router        Router
	reader        *bufio.Reader
}

func NewGameDetailsCli(router Router, reader *bufio.Reader, gameHandler handlers.GamesHandler, reviewHandler handlers.ReviewHandler, rentalHandler handlers.RentalHandler) Cli {
	return &gameDetailsCli{
		router:        router,
		reader:        reader,
		gameHandler:   gameHandler,
		reviewHandler: reviewHandler,
		rentalHandler: rentalHandler,
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
				subActions := []Action{
					{
						Name: "Buy Now",
						ActionFunc: func() {
							//TODO
							// fmt.Printf("Enter game qty to buy: ")
							// qty, err := gDC.reader.ReadString('\n')
							// if err != nil {
							// 	fmt.Println("Error reading game qty input", err)
							// 	time.Sleep(time.Second)
							// }
							// qty = strings.TrimSpace(qty)
							// qtyInt, err := strconv.Atoi(qty)
							// if err != nil {
							// 	fmt.Println("Invalid qty input, integer only")
							// 	time.Sleep(time.Second)
							// }
						},
					},
					{
						Name: "Add To Cart",
						ActionFunc: func() {
							fmt.Printf("Enter game qty to buy: ")
							qty, err := gDC.reader.ReadString('\n')
							if err != nil {
								fmt.Println("Error reading game qty input", err)
								time.Sleep(time.Second)
							}
							qty = strings.TrimSpace(qty)
							qtyInt, err := strconv.Atoi(qty)
							if err != nil {
								fmt.Println("Invalid qty input, integer only")
								time.Sleep(time.Second)
							}
							ci := CartItem{
								Game:      game,
								Qty:       qtyInt,
								BuyOrRent: "Buy",
								RentDays:  0,
							}
							session.CurrentCart.AddItem(&ci)
						},
					},
				}
				PromptUserForActions(subActions, gDC.reader)
			},
		},
		{
			Name: "Rent",
			ActionFunc: func() {
				fmt.Printf("Enter End Date of your rental (yyyy--mm--dd): ")
				endDate, err := gDC.reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading rental end date input", err)
					time.Sleep(time.Second)
					return
				}
				endDate = strings.TrimSpace(endDate)
				//parse date input
				layout := "2006-01-02"
				ed, err := time.Parse(layout, endDate)
				if err != nil {
					fmt.Println("Invalid date format, insert date as yyyy-mm-dd")
					time.Sleep(time.Second)
					return
				}
				if session.CurrentUser == nil {
					fmt.Println("You must logged in to make a rental")
					time.Sleep(time.Second)
					return
				} else {
					rental := entities.Rental{
						UserId:    session.CurrentUser.UserId,
						GameId:    game.GameId,
						StartDate: time.Now(),
						EndDate:   ed,
						Status:    "Not Returned",
					}
					err = gDC.rentalHandler.Create(rental)
					if err != nil {
						fmt.Println(err)
					}
					time.Sleep(time.Second)
				}
			},
		},
		{
			Name: "View Cart",
			ActionFunc: func() {
				gDC.router.Push(routes.CART_ROUTE, RouteArgs{})
			},
		},
		{
			Name: "Add Review",
			ActionFunc: func() {
				if session.CurrentUser == nil {
					fmt.Println("You must logged in to add review")
					time.Sleep(time.Second)
					return
				} else {
					fmt.Printf("Enter Rating (0.0 - 5.0): ")
					rating, err := gDC.reader.ReadString('\n')
					if err != nil {
						fmt.Println("Error reading rating input", err)
						time.Sleep(time.Second)
						return
					}
					rating = strings.TrimSpace(rating)
					ratingFlt, err := strconv.ParseFloat(rating, 64)
					if err != nil {
						fmt.Println("Error converting rating to float64")
						time.Sleep(time.Second)
						return
					}
					if ratingFlt > 5 || ratingFlt < 0 {
						fmt.Println("Rating must be between 0.0 - 5.0")
						time.Sleep(time.Second)
						return
					}
					fmt.Printf("Enter review message: ")
					message, err := gDC.reader.ReadString('\n')
					if err != nil {
						fmt.Println("Error reading message input", err)
						time.Sleep(time.Second)
						return
					}
					message = strings.TrimSpace(message)
					review := entities.Review{
						UserId:    session.CurrentUser.UserId,
						GameId:    game.GameId,
						Rating:    ratingFlt,
						ReviewMsg: message,
					}
					err = gDC.reviewHandler.Create(review)
					if err != nil {
						fmt.Println(err)
					}
					time.Sleep(time.Second)
				}

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
