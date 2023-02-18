package domain

type User struct {
	Name string
	Mail string
	PasswordHash string
	Claims Claims
}