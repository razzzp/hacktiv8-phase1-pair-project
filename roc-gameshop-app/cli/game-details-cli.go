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

func (gDC *gameDetailsCli) GetCustomerActions(game *entities.Game, session *Session) []Action {
	actions := []Action{
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
				fmt.Println("Added to cart")
				time.Sleep(time.Second)
			},
		},
		{
			Name: "Rent",
			ActionFunc: func() {
				if session.CurrentUser == nil {
					// redirect to login page
					gDC.router.Push(routes.LOGIN_REGISTER, RouteArgs{})
					return
				}
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
					gDC.router.Push(routes.LOGIN_REGISTER, RouteArgs{})
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
	}

	return actions
}

func (gDC *gameDetailsCli) GetAdminActions(game *entities.Game, session *Session) []Action {
	actions := []Action{
		{
			Name: "Edit Game",
			ActionFunc: func() {
				gameDto := gDC.gameHandler.GameToDTO(game)
				fmt.Println("Please enter details to update (leave empty to not modify)")
				// update game dto from user input
				gameDto.Name = PromptUserForString(fmt.Sprintf("Enter new name (%s):\n", gameDto.Name), gameDto.Name, gDC.reader)
				gameDto.Genre = PromptUserForString(fmt.Sprintf("Enter new genre (%s):\n", gameDto.Genre), gameDto.Genre, gDC.reader)
				gameDto.Description = PromptUserForString(fmt.Sprint("Enter new description:\n", ""), gameDto.Description, gDC.reader)
				gameDto.SalePrice = PromptUserForString(fmt.Sprintf("Enter new sale price (%s):\n", gameDto.SalePrice), gameDto.SalePrice, gDC.reader)
				gameDto.RentalPrice = PromptUserForString(fmt.Sprintf("Enter new rent price (%s/day):\n", gameDto.RentalPrice), gameDto.RentalPrice, gDC.reader)

				updateGame, err := gDC.gameHandler.ValidateGameDto(&gameDto)
				if err != nil {
					fmt.Printf("Error: %v", err)
				}

				gDC.gameHandler.UpdateGame(*updateGame)
			},
		},
		{
			Name: "Delete Game",
			ActionFunc: func() {
				confirm := PromptUserForString("Are you sure you want to delete this game? (y/n)", "n", gDC.reader)
				if strings.EqualFold(confirm, "y") {
					err := gDC.gameHandler.DeleteGame(game.GameId)
					if err != nil {
						fmt.Printf("Failed to delete game: %v", err)
					} else {
						fmt.Printf("Game '%s' successfully deleted.", game.Name)
						time.Sleep(time.Second)
						gDC.router.Pop()
					}
				}
			},
		},
	}

	return actions
}

func (gDC *gameDetailsCli) GetActions(game *entities.Game, session *Session) []Action {
	var result []Action
	if session.CurrentUser != nil && session.CurrentUser.IsAdmin() {
		result = gDC.GetAdminActions(game, session)
	} else {
		result = gDC.GetCustomerActions(game, session)
	}
	result = append(result, Action{
		Name: "Back",
		ActionFunc: func() {
			gDC.router.Pop()
		},
	})

	return result
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

	fmt.Println("Game Details Page")
	fmt.Println("")
	fmt.Println(game.Name)
	fmt.Println("")
	fmt.Printf("Genre: %s\n", game.Genre)
	fmt.Printf("Description: %s\n", game.Description)
	fmt.Printf("Average Rating: %.2f\n", *rating)
	fmt.Printf("Sale Price: %.2f\n", game.SalePrice)
	fmt.Printf("Rent Price: %.2f/day\n", game.RentalPrice)
	fmt.Println("")
	fmt.Println("Reviews:")
	fmt.Println("")
	for i, review := range reviews {
		fmt.Printf("%d. %s     Rating: %.2f\n", i+1, review.UserName, review.Rating)
		fmt.Println(review.ReviewMsg)
		fmt.Println("")
	}

	actions := gDC.GetActions(game, session)
	PromptUserForActions(actions, gDC.reader)
}
