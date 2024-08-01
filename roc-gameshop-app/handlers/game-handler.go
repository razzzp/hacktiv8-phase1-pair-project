package handlers

import (
	"fmt"
	"roc-gameshop-app/entities"
	"roc-gameshop-app/repos"
)

type GamesHandler interface {
	GetAll(name string, limit, start int) ([]*entities.Game, error)
	GetById(id int) (*entities.Game, error)
	UpdateGame(game entities.Game) error
	CreateGame(game entities.Game) error 
}

type gamesHandler struct {
	gamesRepo repos.GamesRepo
}

func NewGamesHandler(gamesRepo repos.GamesRepo) GamesHandler {
	return &gamesHandler{
		gamesRepo: gamesRepo,
	}
}

func (g *gamesHandler) GetAll(name string, limit, start int) ([]*entities.Game, error) {
	games, err := g.gamesRepo.GetAllGames(name, limit, start)
	if err != nil {
		fmt.Println("error getting all games")
		return nil, err
	}
	return games, nil
}

func (g *gamesHandler) GetById(id int) (*entities.Game, error) {
	game, err := g.gamesRepo.GetGameById(id)
	if err != nil {
		fmt.Println("error get a Game")
		return nil, err
	}
	return game, nil
}

func (g *gamesHandler) UpdateGame(game entities.Game) error {
	err := g.gamesRepo.UpdateGame(game)
	if err != nil {
		fmt.Println("error update game")
		return err
	}
	return nil
}

func (g *gamesHandler) CreateGame(game entities.Game) error {
	err := g.gamesRepo.CreateGame(game)
	if err != nil {
		fmt.Println("error creating game")
		return err
	}
	return nil
}