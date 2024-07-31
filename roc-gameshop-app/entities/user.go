package entities

type User struct {
	UserId       int
	Name         string
	Email        string
	Role         string
	PhoneNumber  string
	Salt         string
	PasswordHash string
}
