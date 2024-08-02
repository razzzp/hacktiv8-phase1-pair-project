package main

import (
	"bufio"
	"log"
	"os"
	"roc-gameshop-app/cli"
	"roc-gameshop-app/config"
	"roc-gameshop-app/handlers"
	"roc-gameshop-app/repos"
	"roc-gameshop-app/routes"
)

func main() {
	config.InitGoDotEnv()

	db, err := config.CreateDBInstance()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Connected to DB.")

	// create router
	router := cli.NewRouter()

	reader := bufio.NewReader(os.Stdin)
	// create clis
	gameDetailsCli := cli.NewGameDetailsCli(router, reader, handlers.NewGamesHandler(repos.NewGamesRepo(db)), handlers.NewReviewHandler(repos.NewReviewsRepo(db)), handlers.NewRentalHandler(repos.NewRentalRepo(db)))
	registerCli := cli.NewUserCli(router, reader, handlers.NewUserHandler(repos.NewUserRepo(db)))
	homepageCli := cli.NewHomepageCli(router, reader)
	loginCli := cli.NewLoginCli(router, reader, handlers.NewAuthHandler(repos.NewUserRepo(db)))
	gamesCli := cli.NewGamesCli(router, reader, handlers.NewGamesHandler(repos.NewGamesRepo(db)))
	cartCli := cli.NewCartCli(router, reader, handlers.NewSaleHandler(repos.NewSaleRepo(db)))
	salesReportCli := cli.NewSalesReportCli(router, reader, handlers.NewSalesReportHandler(db))
	overdueCli := cli.NewOverdueCli(router, reader, handlers.NewRentalHandler(repos.NewRentalRepo(db)))
	loginRegisterCli := cli.NewLoginRegisterCli(router, reader)

	// assign routes
	router.AddRouteCli(routes.HOME_PAGE_ROUTE, homepageCli)
	router.AddRouteCli(routes.GAME_DETAILS_ROUTE, gameDetailsCli)
	router.AddRouteCli(routes.REGISTER_ROUTE, registerCli)
	router.AddRouteCli(routes.LOGIN_ROUTE, loginCli)
	router.AddRouteCli(routes.GAMES_ROUTE, gamesCli)
	router.AddRouteCli(routes.CART_ROUTE, cartCli)
	router.AddRouteCli(routes.SALES_REPORT_ROUTE, salesReportCli)
	router.AddRouteCli(routes.RENTALS_OVERDUE_ROUTE, overdueCli)
	router.AddRouteCli(routes.LOGIN_REGISTER, loginRegisterCli)

	// create session
	session := cli.NewSession()

	// push starting route
	router.Push(routes.HOME_PAGE_ROUTE, cli.RouteArgs{})

	// start
	router.Run(session)
}
