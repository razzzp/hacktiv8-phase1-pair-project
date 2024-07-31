package repos

import "roc-gameshop-app/entities"

type GameRepo interface {
	GetAllGames(name string, limit int, start int) ([]*entities.Game, error)
	CreateGame(game entities.Game) error
	UpdateGame(game entities.Game) error
	GetGameById(id int) (*entities.Game, error)
}

type gameRepo struct {
	// TODO
}

// CreateGame implements GameRepo.
func (g *gameRepo) CreateGame(game entities.Game) error {
	panic("unimplemented")
}

// GetAllGames implements GameRepo.
func (g *gameRepo) GetAllGames(name string, limit int, start int) ([]*entities.Game, error) {
	panic("unimplemented")
}

// GetGameById implements GameRepo.
func (g *gameRepo) GetGameById(id int) (*entities.Game, error) {
	panic("unimplemented")
}

// UpdateGame implements GameRepo.
func (g *gameRepo) UpdateGame(game entities.Game) error {
	panic("unimplemented")
}

func NewGameRepo() GameRepo {
	return &gameRepo{}
}
