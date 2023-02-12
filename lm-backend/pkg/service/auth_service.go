package service

import "license-manager/pkg/domain"

type AuthService interface {
	
	Register(creds domain.Credentials) error
	IsRegistered(user string, passwdHash string) bool
	
	FindByUserNameAndPasswordHash(user, passwdHash string) (domain.Credentials, error)

	MergeClaimsFor(user string, passwdHash string, claims domain.Claims) error
	SetClaimsFor(user string, passwdHash string, claims domain.Claims) error
	
	RevokeCreds(user string, passwdHash string) error

	CreateTokenFor(creds domain.Credentials) (domain.Token, error)

}