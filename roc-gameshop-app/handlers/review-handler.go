package handlers

import (
	"roc-gameshop-app/entities"
	"roc-gameshop-app/repos"
)

type ReviewHandler interface {
	GetByGameId(id int) ([]*entities.ReviewPerGame, error)
	GetAvgRating(id int) (*float64, error)
}

type reviewHandler struct {
	reviewRepo repos.ReviewRepo
}

func NewReviewHandler(reviewRepo repos.ReviewRepo) ReviewHandler {
	return &reviewHandler{
		reviewRepo: reviewRepo,
	}
}

func (r *reviewHandler) GetByGameId(gameId int) ([]*entities.ReviewPerGame, error) {
	reviews, err := r.reviewRepo.GetGameReviews(gameId)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *reviewHandler) GetAvgRating(id int) (*float64, error) {
	avgRating, err := r.reviewRepo.GetGameAvgRating(id)
	if err != nil {
		return nil, err
	}
	return avgRating, nil
}
