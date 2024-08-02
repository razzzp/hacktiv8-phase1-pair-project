package repos

import (
	"database/sql"
	"fmt"
	"roc-gameshop-app/entities"
)

type UserRepo interface {
	GetAllUsers() ([]entities.User, error)
	CreateUser(user entities.User) error
	UpdateUser(id int, user entities.User) error
	GetUserById(id int) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
}

type userRepo struct {
	db *sql.DB
}

// Create User
func (u *userRepo) CreateUser(user entities.User) error {
	query := `
	INSERT INTO Users (Name, Role, Email, PhoneNumber, PasswordHash)
	VALUES (?,?,?,?,?)`
	// fmt.Println("role: ", user.Role)
	_, err := u.db.Exec(query, user.Name, user.Role, user.Email, user.PhoneNumber, user.PasswordHash)
	if err != nil {
		fmt.Println("Error executing create user query")
		return err
	}
	fmt.Printf("Success creating %s as user\n", user.Name)
	return nil
}

// Get All User
func (u *userRepo) GetAllUsers() ([]entities.User, error) {
	query :=
		`SELECT * FROM Users`

	rows, err := u.db.Query(query)
	if err != nil {
		fmt.Println("Error executing get all users query")
		return nil, err
	}
	defer rows.Close()

	users := []entities.User{}
	for rows.Next() {
		user := entities.User{}
		err := rows.Scan(&user.UserId, &user.Name, &user.Role, &user.Email, &user.PhoneNumber, &user.PasswordHash)
		if err != nil {
			fmt.Println("Error scanning returned users data")
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Update User
func (u *userRepo) UpdateUser(id int, user entities.User) error {
	query := `
		UPDATE Users
		SET Name = ?, Email = ?, Role = ?, PhoneNumber = ?, PasswordHash = ?
		WHERE UserId = ?
	`
	_, err := u.db.Exec(query, user.Name, user.Email, user.Role, user.PhoneNumber, user.PasswordHash, id)
	if err != nil {
		fmt.Println("Error executing update user query")
		return err
	}
	fmt.Println("Success updating user")
	return nil
}

// Get User By ID
func (u *userRepo) GetUserById(id int) (*entities.User, error) {
	query := `
		SELECT * FROM Users WHERE UserId = ?
	`

	rows, err := u.db.Query(query, id)
	if err != nil {
		fmt.Println("Error executing get user by id query")
		return nil, err
	}
	defer rows.Close()

	user := entities.User{}
	found := false
	for rows.Next() {
		err := rows.Scan(&user.UserId, &user.Name, &user.Role, &user.Email, &user.PhoneNumber, &user.PasswordHash)
		if err != nil {
			fmt.Println("Error scanning returned user data")
			return nil, err
		}
		found = true
	}
	if !found {
		return nil, fmt.Errorf("user with id '%d' not found", id)
	}
	return &user, nil
}

// GetUserByEmail retrieves a user by their email from the database.
func (u *userRepo) GetUserByEmail(email string) (*entities.User, error) {
	query := `
		SELECT * FROM Users WHERE Email = ?
	`

	rows, err := u.db.Query(query, email)
	if err != nil {
		fmt.Println("Error executing get user by email query")
		return nil, err
	}
	defer rows.Close()

	user := entities.User{}
	found := false
	for rows.Next() {
		err := rows.Scan(&user.UserId, &user.Name, &user.Role, &user.Email, &user.PhoneNumber, &user.PasswordHash)
		if err != nil {
			fmt.Println("Error scanning returned user data")
			return nil, err
		}
		found = true
	}

	if !found {
		return nil, fmt.Errorf("user with email '%s' not found", email)
	}

	return &user, nil
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{db}
}
