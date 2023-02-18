package domain

type User struct {
	ID string
	Name string
	Mail string
	PasswordHash string
	Claims Claims
}