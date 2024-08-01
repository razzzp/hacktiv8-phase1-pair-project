package repos

import (
	"database/sql"
	"fmt"
	"roc-gameshop-app/entities"
)

// ReviewRepo cara memasukkan ke reviews data.
type ReviewRepo interface {
	CreateReview(review entities.Review) error
	GetReviewsByGame(gameId int) ([]entities.Review, error)
	GetReviewsByUser(userId int) ([]entities.Review, error)
	UpdateReview(review entities.Review) error
	DeleteReview(reviewId int) error
}

// reviewRepo implementasi ke dalam ReviewRepo.
type reviewRepo struct {
	db *sql.DB
}

// NewReviewRepo membuat data baru ReviewRepo.
func NewReviewRepo(db *sql.DB) ReviewRepo {
	return &reviewRepo{db}
}

// CreateReview memasukkan review baru ke dalam database.
func (rR *reviewRepo) CreateReview(review entities.Review) error {
	query := `INSERT INTO Reviews (GameId, UserId, Rating, ReviewMsg) VALUES (?, ?, ?, ?)`
	_, err := rR.db.Exec(query, review.GameId, review.UserId, review.Rating, review.ReviewMsg)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully created review for game %d by user %d\n", review.GameId, review.UserId)
	return nil
}

// GetReviewsByGame mengambil/ membaca review dari game.
func (rR *reviewRepo) GetReviewsByGame(gameId int) ([]entities.Review, error) {
	query := `SELECT * FROM Reviews WHERE GameId = ?`
	rows, err := rR.db.Query(query, gameId)
	if err != nil {
		fmt.Println("Error executing get reviews by game query")
		return nil, err
	}
	defer rows.Close()

	reviews := []entities.Review{}
	for rows.Next() {
		var review entities.Review
		err := rows.Scan(&review.ReviewId, &review.GameId, &review.UserId, &review.Rating, &review.ReviewMsg)
		if err != nil {
			fmt.Println("Error scanning returned review data")
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

// GetReviewsByUser membaca/ melihat review yang dibuat oleh user.
func (rR *reviewRepo) GetReviewsByUser(userId int) ([]entities.Review, error) {
	query := `SELECT * FROM Reviews WHERE UserId = ?`
	rows, err := rR.db.Query(query, userId)
	if err != nil {
		fmt.Println("Error executing get reviews by user query")
		return nil, err
	}
	defer rows.Close()

	reviews := []entities.Review{}
	for rows.Next() {
		var review entities.Review
		err := rows.Scan(&review.ReviewId, &review.GameId, &review.UserId, &review.Rating, &review.ReviewMsg)
		if err != nil {
			fmt.Println("Error scanning returned review data")
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

// UpdateReview update review dalam database.
func (rR *reviewRepo) UpdateReview(review entities.Review) error {
	query := `UPDATE Reviews SET Rating = ?, ReviewMsg = ? WHERE ReviewId = ?`
	_, err := rR.db.Exec(query, review.Rating, review.ReviewMsg, review.ReviewId)
	if err != nil {
		fmt.Println("Error executing update review query")
		return err
	}
	fmt.Printf("Successfully updated review with id %d\n", review.ReviewId)
	return nil
}

// DeleteReview menghapus review dari database.
func (rR *reviewRepo) DeleteReview(reviewId int) error {
	query := `DELETE FROM Reviews WHERE ReviewId = ?`
	_, err := rR.db.Exec(query, reviewId)
	if err != nil {
		fmt.Println("Error executing delete review query")
		return err
	}
	fmt.Printf("Successfully deleted review with id %d\n", reviewId)
	return nil
}
