package domain

type Credentials struct {
	ID string
	UserName string
	PasswordHash string
	Claims Claims
}