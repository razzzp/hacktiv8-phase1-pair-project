package repos

import (
	"database/sql"
	"fmt"
	"roc-gameshop-app/entities"
)

type GamesRepo interface {
	GetAllGames(name string, limit int, start int) ([]*entities.Game, error)
	CreateGame(game entities.Game) error
	UpdateGame(game entities.Game) error
	GetGameById(id int) (*entities.Game, error)
}

type gamesRepo struct {
	db *sql.DB
}

// CreateGame implements GameRepo.
func (gR *gamesRepo) CreateGame(game entities.Game) error {
	query := `INSERT INTO Games(
		Name,Description,Genre,SalePrice,RentalPrice,Studio,Stock
	) VALUES (
		?,?,?,?,?,?,?
	)`
	_, err := gR.db.Exec(
		query,
		game.Name,
		game.Description,
		game.Genre,
		game.SalePrice,
		game.RentalPrice,
		game.Studio,
		game.Stock,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetAllGames implements GameRepo.
func (gR *gamesRepo) GetAllGames(name string, limit int, start int) ([]*entities.Game, error) {

	var rows *sql.Rows
	var err error
	nameQry := fmt.Sprintf("%%%s%%", name)
	if name == "" {
		query := `
		SELECT * FROM Games
		`
		rows, err = gR.db.Query(query)
	} else {
		query := `
		SELECT * FROM Games WHERE Name LIKE ?
		`
		rows, err = gR.db.Query(query, nameQry)
	}
	if err != nil {
		return nil, err
	}

	result := []*entities.Game{}
	for rows.Next() {
		var game entities.Game
		err = rows.Scan(
			&game.GameId,
			&game.Name,
			&game.Description,
			&game.Genre,
			&game.SalePrice,
			&game.RentalPrice,
			&game.Studio,
			&game.Stock,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, &game)
	}
	return result, nil
}

// GetGameById implements GameRepo.
func (gR *gamesRepo) GetGameById(id int) (*entities.Game, error) {

	query := `SELECT * FROM Games WHERE GameId = ?`
	rows, err := gR.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		var game entities.Game
		err = rows.Scan(
			&game.GameId,
			&game.Name,
			&game.Description,
			&game.Genre,
			&game.SalePrice,
			&game.RentalPrice,
			&game.Studio,
			&game.Stock,
		)
		if err != nil {
			return nil, err
		}

		return &game, nil
	}
	return nil, fmt.Errorf("game with id %d not found", id)
}

// UpdateGame implements GameRepo.
func (gR *gamesRepo) UpdateGame(game entities.Game) error {
	query := `UPDATE Games
		SET Name=?,Description=?,Genre=?,SalePrice=?,RentalPrice=?,Studio=?,Stock=?
		WHERE GameId = ?`
	_, err := gR.db.Exec(
		query,
		game.Name,
		game.Description,
		game.Genre,
		game.SalePrice,
		game.RentalPrice,
		game.Studio,
		game.Stock,
		game.GameId,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewGamesRepo(db *sql.DB) GamesRepo {
	return &gamesRepo{
		db: db,
	}
}
