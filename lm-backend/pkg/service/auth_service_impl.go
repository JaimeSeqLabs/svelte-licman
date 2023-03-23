package service

import (
	"errors"
	"license-manager/pkg/domain"
	"license-manager/pkg/pkgerrors"
	"license-manager/pkg/repositories"
	ent_fw_common "license-manager/pkg/repositories/ent-fw/common"
	"license-manager/pkg/repositories/ent-fw/ent"
	"log"
)


type authService struct {
	userRepo repositories.UserRepository
	jwtService JWTService
}

func NewAuthService(userRepo repositories.UserRepository, jwtService JWTService) AuthService {
	return &authService{
		userRepo: userRepo,
		jwtService: jwtService,
	}
}

func (auth *authService) RegisterUser(user domain.User) error {
	registered, err := auth.IsRegistered(user)
	if err != nil {
		return err
	}
	if registered {
		return errors.New("user already registered")
	}
	return auth.userRepo.Save(user)
}

func (auth *authService) IsRegistered(user domain.User) (bool, error) {
	found, err := auth.userRepo.FindByNameAndMail(user.Name, user.Mail)
	if err != nil {

		if ent.IsNotFound(err) { // TODO: remove this dependency from ent
			return false, nil
		}
		return false, err
	}
	return (found.ID != ""), nil
}

func (auth *authService) FindUserByMailAndPsswd(mail, passwdHash string) (domain.User, error) {
	return auth.userRepo.FindByMailAndPassword(mail, passwdHash)
}

func (auth *authService) MergeClaimsFor(user domain.User, claims domain.Claims) error {
	
	user.Claims = ent_fw_common.MergeClaims(user.Claims, claims)
	
	updated, err := auth.userRepo.Update(user)
	if err != nil {
		return err
	}
	if !updated {
		return errors.New("user is not registered, unable to update claims")
	}

	return nil
}

func (auth *authService) SetClaimsFor(user domain.User, claims domain.Claims) error {
	
	user.Claims = claims
	
	updated, err := auth.userRepo.Update(user)
	if err != nil {
		return err
	}
	if !updated {
		return errors.New("user is not registered, unable to set claims")
	}
	
	return nil
}

func (auth *authService) CreateTokenFor(user domain.User) (domain.Token, error) {

	// if ID not provided double check if user is registered
	if user.ID == "" {
		dbUser, err := auth.userRepo.FindByMailAndPassword(user.Mail, user.PasswordHash)
		if err != nil {
			return domain.Token{}, err
		}
		user = dbUser
	}

	token, err := auth.jwtService.GenTokenFor(user, user.Claims) // for now same claims as user
	if err != nil {
		return domain.Token{}, err
	}

	log.Printf("new token issued to user %s %s\n", user.Name, user.Mail)

	return token, err
}

func (auth *authService) GetFirstTokenFor(user domain.User) (domain.Token, error) {
	
	usr, err := auth.userRepo.FindByMailAndPassword(user.Mail, user.PasswordHash)
	if err != nil {
		return domain.Token{}, err
	}

	tokens, err := auth.jwtService.GetIssuedBy(usr.ID)
	if err != nil {
		return domain.Token{}, err
	}

	if len(tokens) < 1 {
		return domain.Token{}, pkgerrors.ErrNoTokensFound
	}

	return tokens[0], nil
}

func (auth *authService) RevokeTokensFor(user domain.User) (revoked int, err error) {
	
	if user.ID == "" {
		dbUser, err := auth.userRepo.FindByMailAndPassword(user.Mail, user.PasswordHash)
		if err != nil {
			return 0, err
		}
		user = dbUser
	}

	return auth.jwtService.RevokeTokensFor(user)
}