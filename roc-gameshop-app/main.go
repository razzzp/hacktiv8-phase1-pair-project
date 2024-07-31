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

	log.Println("Connect to DB.")

	// create router
	router := cli.NewRouter()

	reader := bufio.NewReader(os.Stdin)
	// create clis
	gameDetailsCli := cli.NewGameDetailsCli(router, reader, handlers.NewGameHandler(repos.NewGameRepo(db)))
	registerCli := cli.NewUserCli(router, reader, handlers.NewUserHandler(repos.NewUserRepo(db)))
	homepageCli := cli.NewHomepageCli(router, reader)
	// assign routes
	router.AddRouteCli(routes.HOME_PAGE_ROUTE, homepageCli)
	router.AddRouteCli(routes.GAME_DETAILS_ROUTE, gameDetailsCli)
	router.AddRouteCli(routes.REGISTER_ROUTE, registerCli)

	// create session
	session := cli.Session{}

	// push starting route
	router.Push(routes.HOME_PAGE_ROUTE, cli.RouteArgs{})

	// start
	router.Run(&session)
}
