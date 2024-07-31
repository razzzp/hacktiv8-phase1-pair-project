package repos_test

import (
	"fmt"
	"log"
	"roc-gameshop-app/config"
	"roc-gameshop-app/entities"
	"roc-gameshop-app/repos"
	"testing"
)

func TestGamesRepo(t *testing.T) {

	config.InitGoDotEnv()

	db, err := config.CreateDBInstance()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	gameRepo := repos.NewGameRepo(db)
	result, err := gameRepo.GetAllGames("", 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, g := range result {
		fmt.Println(g)
	}

	game, err := gameRepo.GetGameById(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(game)

	newGame := entities.Game{
		Name:        "newgame",
		Description: "brand new",
		Genre:       "unknown",
		SalePrice:   10,
		RentalPrice: 10,
		Studio:      "unknown",
		Stock:       10,
	}
	err = gameRepo.CreateGame(newGame)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created game")

	game.Description = "modif"
	err = gameRepo.UpdateGame(*game)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated game")
}
