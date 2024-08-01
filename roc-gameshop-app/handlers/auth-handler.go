package handlers

import (
	"fmt"
	"roc-gameshop-app/entities"
	"roc-gameshop-app/helpers"
	"roc-gameshop-app/repos"
)

type AuthHandler interface {
	Login(email, password string) (*entities.User, error)
}

type authHandler struct {
	userRepo repos.UserRepo
}

func NewAuthHandler(userRepo repos.UserRepo) AuthHandler {
	return &authHandler{
		userRepo: userRepo,
	}
}

func (a *authHandler) Login(email, password string) (*entities.User, error) {
	if email == "" {
		return nil, fmt.Errorf("Email can't be empty")
	}
	if password == "" {
		return nil, fmt.Errorf("Password can't be empty")
	}
	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil {
		fmt.Println("error getting user by email")
		return nil, err
	}

	ok := helpers.ComparePasswords(user.PasswordHash, []byte(password))

	if !ok {
		return nil, fmt.Errorf("Invalid username and password")
	} else {
		return user, nil
	}

}