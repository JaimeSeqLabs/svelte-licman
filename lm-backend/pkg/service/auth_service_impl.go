package service

import (
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
)


type authService struct {
	credsRepo repositories.CredentialsRepository
	jwtService JWTService
}

func NewAuthService(credsRepo repositories.CredentialsRepository, jwtService JWTService) AuthService {
	return &authService{
		credsRepo: credsRepo,
		jwtService: jwtService,
	}
}

func (auth *authService) Register(creds domain.Credentials) error {
	return auth.credsRepo.Save(creds)
}

func (auth *authService) IsRegistered(user string, passwdHash string) bool {
	creds, err := auth.credsRepo.FindByUserNameAndPasswordHash(user, passwdHash)
	if err != nil {
		return false
	}
	// ensure struct is not empty
	return creds.UserName == user && creds.PasswordHash == passwdHash
}

func (auth *authService) FindByUserNameAndPasswordHash(user string, passwdHash string) (domain.Credentials, error) {
	return auth.credsRepo.FindByUserNameAndPasswordHash(user, passwdHash)
}

func (auth *authService) MergeClaimsFor(user string, passwdHash string, claims domain.Claims) error {
	_, err := auth.credsRepo.MergeClaimsFor(user, passwdHash, claims)
	return err
}

func (auth *authService) SetClaimsFor(user string, passwdHash string, claims domain.Claims) error {
	return auth.credsRepo.Update(domain.Credentials{
		UserName: user,
		PasswordHash: passwdHash,
		Claims: claims,
	})
}

func (auth *authService) CreateTokenFor(creds domain.Credentials) (domain.Token, error) {
	token, err := auth.jwtService.GenTokenFor(creds.Claims)
	if err != nil {
		return domain.Token{}, err
	}
	return token, nil
}

func (auth *authService) RevokeCreds(user string, passwdHash string) error {
	// TODO: maybe cascade to jwt tokens also
	return auth.credsRepo.DeleteByUserNameAndPasswordHash(user, passwdHash)
}