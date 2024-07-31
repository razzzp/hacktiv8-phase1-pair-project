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

	// db, err := config.CreateDBInstance()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	log.Println("Connect to DB.")

	// create router
	router := cli.NewRouter()

	reader := bufio.NewReader(os.Stdin)
	// create clis
	gameDetailsCli := cli.NewGameDetailsCli(router, reader, handlers.NewGameHandler(repos.NewGameRepo()))

	// assign routes
	router.AddRouteCli(routes.GAME_DETAILS_ROUTE, gameDetailsCli)

	// create session
	session := cli.Session{}

	// push starting route
	routeArgs := cli.RouteArgs{
		"gameId": "1",
	}
	router.Push(routes.GAME_DETAILS_ROUTE, routeArgs)

	// start
	router.Run(&session)
}
