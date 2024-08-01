package entities

import "time"

type Rental struct {
	RentalId  int
	UserId    int
	GameId    int
	StartDate time.Time
	EndDate   time.Time
	Status    string
}

type RentalOverdue struct {
	UserName  string
	GameName  string
	StartDate string
	EndDate   string
	Status    string
}

type RentalDTOUpdate struct {
	RentalId  string
	UserId    string
	GameId    string
	StartDate time.Time
	EndDate   time.Time
	Status    string
}
