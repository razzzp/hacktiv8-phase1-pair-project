package repos

import "roc-gameshop-app/entities"

type UserRepo interface {
	GetAllUsers() ([]*entities.User, error)
	CreateUser(user entities.User) error
	UpdateUser(user entities.User) error
	GetUserById(id int) (*entities.User, error)
}
