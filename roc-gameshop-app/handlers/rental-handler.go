package handlers

import (
	"fmt"
	"roc-gameshop-app/entities"
	"roc-gameshop-app/repos"
	"strconv"
)

type RentalHandler interface {
	GetAll() ([]entities.Rental, error)
	GetById(id int) (*entities.Rental, error)
	Create(rental entities.Rental) error
	Update(id int, rental entities.RentalDTOUpdate) error
}

type rentalHandler struct {
	rentalRepo repos.RentalRepo
}

func NewRentalHandler(rentalRepo repos.RentalRepo) RentalHandler {
	return &rentalHandler{
		rentalRepo: rentalRepo,
	}
}

func (r *rentalHandler) GetAll() ([]entities.Rental, error) {
	rentals, err := r.rentalRepo.GetAllRentals()
	if err != nil {
		fmt.Println("error getting All Rentals")
		return nil, err
	}
	return rentals, nil
}

func (r *rentalHandler) GetById(id int) (*entities.Rental, error) {
	rental, err := r.rentalRepo.GetRentalById(id)
	if err != nil {
		fmt.Println("error get a Rental")
		return nil, err
	}
	return rental, nil
}

func (r *rentalHandler) Create(rental entities.Rental) error {
	// Check if UserId is zero
	if rental.UserId == 0 {
		return fmt.Errorf("UserId can't be zero")
	}

	// Check if GameId is zero
	if rental.GameId == 0 {
		return fmt.Errorf("GameId can't be zero")
	}

	// Check if StartDate is not empty (not the zero value)
	if rental.StartDate.IsZero() {
		return fmt.Errorf("Start date can't be empty")
	}

	// Check if StartDate is before EndDate
	if rental.StartDate.After(rental.EndDate) {
		return fmt.Errorf("Start date must be before end date")
	}

	// Check if StartDate is not empty (not the zero value)
	if rental.EndDate.IsZero() {
		return fmt.Errorf("Start date can't be empty")
	}

	// Check if Status is empty
	if rental.Status == "" {
		return fmt.Errorf("Status can't be empty")
	}

	err := r.rentalRepo.CreateRental(rental)
	if err != nil {
		fmt.Println("error create Rental")
		return err
	}
	return nil
}

func (r *rentalHandler) Update(id int, rental entities.RentalDTOUpdate) error {
	userIdInt, err := strconv.Atoi(rental.UserId)
	if err != nil {
		return fmt.Errorf("Invalid int value for UserId in Rental")
	}
	gameIdInt, err := strconv.Atoi(rental.GameId)
	if err != nil {
		return fmt.Errorf("Invalid int value for UserId in Rental")
	}

	// Check if UserId is zero
	if userIdInt == 0 {
		return fmt.Errorf("UserId can't be zero")
	}

	// Check if GameId is zero
	if gameIdInt == 0 {
		return fmt.Errorf("GameId can't be zero")
	}

	// Check if StartDate is not empty (not the zero value)
	if rental.StartDate.IsZero() {
		return fmt.Errorf("Start date can't be empty")
	}

	// Check if Status is empty
	if rental.Status == "" {
		fmt.Errorf("Status can't be empty, default to Returned")
		rental.Status = "Returned"
	}

	// Check if StartDate is before EndDate
	if rental.StartDate.After(rental.EndDate) {
		return fmt.Errorf("Start date must be before end date")
	}

	updatedRental := entities.Rental{
		UserId:    userIdInt,
		GameId:    gameIdInt,
		StartDate: rental.StartDate,
		EndDate:   rental.EndDate,
		Status:    rental.Status,
	}
	err = r.rentalRepo.UpdateRental(id, updatedRental)
	if err != nil {
		fmt.Println("error update User")
		return err
	}
	return nil
}
