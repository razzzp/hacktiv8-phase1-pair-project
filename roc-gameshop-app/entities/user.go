package entities

import (
	"fmt"
	"strings"
)

type User struct {
	UserId       int
	Name         string
	Email        string
	Role         string
	PhoneNumber  string
	PasswordHash string
}

func (u *User) IsAdmin() bool {
	fmt.Println("role: ", u.Role)
	return strings.ToLower(u.Role) == "admin"
}
