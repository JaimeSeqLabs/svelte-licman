package service

import (
	"errors"
	"license-manager/pkg/domain"
	"license-manager/pkg/pkgerrors"
	"license-manager/pkg/repositories"
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
	
	user.Claims = merge(user.Claims, claims)
	
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


func merge(original, updated map[string]any) map[string]any {
	for k, v := range updated {
		original[k] = v
	}
	return original
}

/* func (auth *authService) Register(creds domain.Credentials) error {
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
} */