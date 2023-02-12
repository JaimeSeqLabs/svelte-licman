package domain

type Credentials struct {
	UserName string
	PasswordHash string
	Claims Claims
}