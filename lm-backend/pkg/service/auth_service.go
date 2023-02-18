package service

import "license-manager/pkg/domain"

type AuthService interface {
	
	RegisterUser(user domain.User) error
	IsRegistered(user domain.User) (bool, error)
	
	FindUserByMailAndPsswd(mail, passwdHash string) (domain.User, error)

	MergeClaimsFor(user domain.User, claims domain.Claims) error
	SetClaimsFor(user domain.User, claims domain.Claims) error

	CreateTokenFor(user domain.User) (domain.Token, error)
	RevokeTokensFor(user domain.User) (revoked int, err error)

}