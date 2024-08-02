package handlers

import (
	"errors"
	"fmt"
	"roc-gameshop-app/entities"
	"roc-gameshop-app/repos"
	"strconv"
)

type GameDto struct {
	GameId      int
	Name        string
	Description string
	Genre       string
	SalePrice   string
	RentalPrice string
	Studio      string
	Stock       string
}

type GamesHandler interface {
	GetAll(name string, limit, start int) ([]*entities.Game, error)
	GetById(id int) (*entities.Game, error)
	UpdateGame(game entities.Game) error
	CreateGame(game entities.Game) error
	GameToDTO(game *entities.Game) GameDto
	ValidateGameDto(gameDto *GameDto) (*entities.Game, error)
	DeleteGame(gameId int) error
}

type gamesHandler struct {
	gamesRepo repos.GamesRepo
}

func NewGamesHandler(gamesRepo repos.GamesRepo) GamesHandler {
	return &gamesHandler{
		gamesRepo: gamesRepo,
	}
}

// DeleteGame implements GamesHandler.
func (g *gamesHandler) DeleteGame(gameId int) error {
	return g.gamesRepo.DeleteGame(gameId)
}

// ValidateGameDto implements GamesHandler.
func (g *gamesHandler) ValidateGameDto(gameDto *GameDto) (*entities.Game, error) {
	if gameDto == nil {
		return nil, errors.New("game dto is nil")
	}
	result := &entities.Game{}
	result.GameId = gameDto.GameId
	result.Name = gameDto.Name
	result.Description = gameDto.Description
	result.Genre = gameDto.Genre
	result.Studio = gameDto.Studio

	rentalPrice, err := strconv.ParseFloat(gameDto.RentalPrice, 64)
	if err != nil {
		return nil, errors.New("invalid rental price")
	}
	result.RentalPrice = rentalPrice

	salePrice, err := strconv.ParseFloat(gameDto.RentalPrice, 64)
	if err != nil {
		return nil, errors.New("invalid sale price")
	}
	result.SalePrice = salePrice

	stock, err := strconv.Atoi(gameDto.Stock)
	if err != nil {
		return nil, errors.New("invalid stock")
	}
	result.Stock = stock

	return result, nil
}

func (g *gamesHandler) GameToDTO(game *entities.Game) GameDto {
	result := GameDto{}
	result.GameId = game.GameId
	result.Name = game.Name
	result.Description = game.Description
	result.Genre = game.Genre
	result.Studio = game.Studio
	result.RentalPrice = fmt.Sprintf("%.2f", game.RentalPrice)
	result.SalePrice = fmt.Sprintf("%.2f", game.SalePrice)
	result.Stock = fmt.Sprintf("%d", game.Stock)

	return result
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
