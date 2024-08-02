package handlers

import (
	"fmt"
	"roc-gameshop-app/entities"
	"roc-gameshop-app/helpers"
	"roc-gameshop-app/repos"
)

type UserHandler interface {
	GetAll() ([]entities.User, error)
	GetById(id int) (*entities.User, error)
	Create(user entities.User) error
	Update(id int, user entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
}

type userHandler struct {
	userRepo repos.UserRepo
}

// GetUserByEmail implements UserHandler.
func (u *userHandler) GetUserByEmail(email string) (*entities.User, error) {
	return u.userRepo.GetUserByEmail(email)
}

func NewUserHandler(userRepo repos.UserRepo) UserHandler {
	return &userHandler{
		userRepo: userRepo,
	}
}

func (u *userHandler) GetAll() ([]entities.User, error) {
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		fmt.Println("error getting All Users")
		return nil, err
	}
	return users, nil
}

func (u *userHandler) GetById(id int) (*entities.User, error) {
	user, err := u.userRepo.GetUserById(id)
	if err != nil {
		fmt.Println("error get a User")
		return nil, err
	}
	return user, nil
}

func (u *userHandler) Create(user entities.User) error {
	if user.Name == "" {
		return fmt.Errorf("User name can't be empty")
	}
	if user.Email == "" {
		return fmt.Errorf("User email can't be empty")
	}
	if user.Role == "" {
		return fmt.Errorf("User role can't be empty")
	}
	if user.PasswordHash == "" {
		return fmt.Errorf("User password can't be empty")
	}
	user.PasswordHash = helpers.HashAndSalt([]byte(user.PasswordHash))
	err := u.userRepo.CreateUser(user)
	if err != nil {
		fmt.Println("error create User")
		return err
	}
	return nil
}

func (u *userHandler) Update(id int, user entities.User) error {
	if user.Name == "" {
		return fmt.Errorf(`User name can't be empty`)
	}
	if user.Email == "" {
		return fmt.Errorf("User email can't be empty")
	}
	if user.Role == "" {
		return fmt.Errorf("User role can't be empty")
	}
	if user.PasswordHash == "" {
		return fmt.Errorf("User password can't be empty")
	}
	user.PasswordHash = helpers.HashAndSalt([]byte(user.PasswordHash))
	err := u.userRepo.UpdateUser(id, user)
	if err != nil {
		fmt.Println("error update User")
		return err
	}
	return nil
}
