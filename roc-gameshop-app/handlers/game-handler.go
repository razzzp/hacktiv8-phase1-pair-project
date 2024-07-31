package handlers

import "roc-gameshop-app/repos"

type GameHandler interface {
}

type gameHandler struct {
	gameRepo *repos.GameRepo
}

func NewGameHandler(gameRepo repos.GameRepo) GameHandler {
	return &gameHandler{
		gameRepo: &gameRepo,
	}
}
