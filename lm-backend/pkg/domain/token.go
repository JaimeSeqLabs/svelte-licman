package domain

type Token struct {
	ID string
	Value string
	Revoked bool
	Claims Claims
	IssuerID string
}