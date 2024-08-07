package repos

import (
	"database/sql"
	"fmt"
	"roc-gameshop-app/entities"
)

type RentalRepo interface {
	CreateRental(rental entities.Rental) error
	GetAllRentals() ([]entities.Rental, error)
	GetRentalById(id int) (*entities.Rental, error)
	UpdateRental(id int, rental entities.Rental) error
	GetRentalsOverdue() ([]*entities.RentalOverdue, error)
}

type rentalRepo struct {
	db *sql.DB
}

func (r *rentalRepo) CreateRental(rental entities.Rental) error {
	query := `
		INSERT INTO Rentals (UserId, GameId, StartDate, EndDate, Status)
		VALUES (?,?,?,?,?)
	`

	_, err := r.db.Exec(query, rental.UserId, rental.GameId, rental.StartDate, rental.EndDate, rental.Status)
	if err != nil {
		fmt.Println("Error executing create rental query")
		return err
	}
	fmt.Println("Success creating rental")
	return nil
}

func (r *rentalRepo) GetAllRentals() ([]entities.Rental, error) {
	query :=
		`SELECT * FROM Rentals`

	rows, err := r.db.Query(query)
	if err != nil {
		fmt.Println("Error executing get all rentals query")
		return nil, err
	}
	defer rows.Close()

	rentals := []entities.Rental{}
	for rows.Next() {
		rental := entities.Rental{}
		err := rows.Scan(&rental.RentalId, &rental.UserId, &rental.GameId, &rental.StartDate, &rental.EndDate, &rental.Status)
		if err != nil {
			fmt.Println("Error scanning returned rentals data")
			return nil, err
		}
		rentals = append(rentals, rental)
	}

	return rentals, nil
}

func (r *rentalRepo) UpdateRental(id int, rental entities.Rental) error {
	query := `
		UPDATE Rentals
		SET UserId = ?, GameId = ?, StartDate = ?, EndDate = ?, Status = ?
		WHERE RentalId = ?
	`
	_, err := r.db.Exec(query, rental.UserId, rental.GameId, rental.StartDate, rental.EndDate, rental.Status, id)
	if err != nil {
		fmt.Println("Error executing update rental query")
		return err
	}
	fmt.Println("Success updating rental")
	return nil
}

func (r *rentalRepo) GetRentalById(id int) (*entities.Rental, error) {
	query := `
		SELECT * FROM Rentals WHERE RentalId = ?
	`

	rows, err := r.db.Query(query, id)
	if err != nil {
		fmt.Println("Error executing get rental by id query")
		return nil, err
	}
	defer rows.Close()

	rental := entities.Rental{}
	for rows.Next() {
		err := rows.Scan(&rental.RentalId, &rental.UserId, &rental.GameId, &rental.StartDate, &rental.EndDate, &rental.Status)
		if err != nil {
			fmt.Println("Error scanning returned rental data")
			return nil, err
		}
	}
	return &rental, nil
}

func (r *rentalRepo) GetRentalsOverdue() ([]*entities.RentalOverdue, error) {
	query :=
		`
	SELECT Users.Name, Games.Name, Rentals.StartDate , Rentals.EndDate, Rentals.Status  FROM Rentals 
	INNER JOIN Users ON Rentals.UserId = Users.UserId 
	INNER JOIN Games ON Rentals.GameId = Games.GameId 
	WHERE Rentals.EndDate < NOW() AND Rentals.Status = "Not Returned"
	`
	rows, err := r.db.Query(query)
	if err != nil {
		fmt.Println("Error executing get all rentals query")
		return nil, err
	}
	defer rows.Close()

	rentals := []*entities.RentalOverdue{}
	for rows.Next() {
		rental := entities.RentalOverdue{}
		err := rows.Scan(&rental.UserName, &rental.GameName, &rental.StartDate, &rental.EndDate, &rental.Status)
		if err != nil {
			fmt.Println("Error scanning returned rentals overdue data")
			return nil, err
		}
		rentals = append(rentals, &rental)
	}

	return rentals, nil
}

func NewRentalRepo(db *sql.DB) RentalRepo {
	return &rentalRepo{db}
}
