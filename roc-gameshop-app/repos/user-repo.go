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
}

type userRepo struct {
	db *sql.DB
}

// Create User
func (u *userRepo) CreateUser(user entities.User) error {
	query := `
	INSERT INTO users (Name, Role, Email, PhoneNumber, Salt, PasswordHash)
	VALUES (?,?,?,?,?,?)`

	_, err := u.db.Exec(query, user.Name, user.Role, user.Email, user.PhoneNumber, user.Salt, user.PasswordHash)
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
		`SELECT * FROM users`

	rows, err := u.db.Query(query)
	if err != nil {
		fmt.Println("Error executing get all users query")
		return nil, err
	}
	users := []entities.User{}
	for rows.Next() {
		user := entities.User{}
		err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Role, &user.PhoneNumber, &user.Salt, &user.PasswordHash)
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
		UPDATE users
		SET Name = ?, Email = ?, Role = ?, PhoneNumber = ?, Salt = ?, PasswordHash = ?
		WHERE UserId = ?
	`
	_, err := u.db.Exec(query, user.Name, user.Role, user.Email, user.PhoneNumber, user.Salt, user.PasswordHash, id)
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
		SELECT * FROM users WHERE UserId = ?
	`

	rows, err := u.db.Query(query, id)
	if err != nil {
		fmt.Println("Error executing get user by id query")
		return nil, err
	}
	user := entities.User{}
	for rows.Next() {
		err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Role, &user.PhoneNumber, &user.Salt, &user.PasswordHash)
		if err != nil {
			fmt.Println("Error scanning returned user data")
			return nil, err
		}
	}
	return &user, nil
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{db}
}
