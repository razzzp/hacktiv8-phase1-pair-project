package repos

import (
	"database/sql"
	"fmt"
	"roc-gameshop-app/entities"
)

type ReviewRepo interface {
	GetGameReviews(gameId int) ([]*entities.ReviewPerGame, error)
	GetGameAvgRating(gameId int) (*float64, error)
	CreateReview(review entities.Review) error
}

type reviewRepo struct {
	db *sql.DB
}

func NewReviewsRepo(db *sql.DB) ReviewRepo {
	return &reviewRepo{
		db: db,
	}
}

func (rR *reviewRepo) GetGameReviews(gameId int) ([]*entities.ReviewPerGame, error) {
	query :=
		`
	SELECT Reviews.ReviewId, Reviews.Rating , Reviews.ReviewMsg, Users.Name  FROM Reviews 
	INNER JOIN Users ON Reviews.UserId = Users.UserId 
	INNER JOIN Games ON Reviews.GameId = Games.GameId 
	WHERE Reviews.GameId = ?

	`

	rows, err := rR.db.Query(query, gameId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviews := []*entities.ReviewPerGame{}
	for rows.Next() {
		review := entities.ReviewPerGame{}

		err = rows.Scan(
			&review.ReviewId, &review.Rating, &review.ReviewMsg, &review.UserName,
		)
		if err != nil {
			return nil, err
		}

		reviews = append(reviews, &review)
	}

	return reviews, nil
}

func (rR *reviewRepo) GetGameAvgRating(gameId int) (*float64, error) {
	query :=
		`
		SELECT AVG(Reviews.Rating)  FROM Reviews 
		INNER JOIN Users ON Reviews.UserId = Users.UserId 
		INNER JOIN Games ON Reviews.GameId = Games.GameId 
		WHERE Reviews.GameId = ?
	`
	rows, err := rR.db.Query(query, gameId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var avgRating float64
	for rows.Next() {
		err = rows.Scan(
			&avgRating,
		)
		if err != nil {
			return nil, err
		}
	}

	return &avgRating, nil
}

func (rR *reviewRepo) CreateReview(review entities.Review) error {
	query :=
		`
		INSERT INTO Reviews (UserId, GameId, Rating, ReviewMsg)
		VALUES (?,?,?,?)
	`

	_, err := rR.db.Exec(query, review.UserId, review.GameId, review.Rating, review.ReviewMsg)
	if err != nil {
		fmt.Println("Error executing create review query")
		return err
	}
	fmt.Println("Success creating a review")
	return nil
}
