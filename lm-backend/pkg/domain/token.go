package domain

type Token struct {
	Value string
	Revoked bool
	Claims Claims
}