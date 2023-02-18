package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"license-manager/pkg/controller/exchange"
	"license-manager/pkg/domain"
	"license-manager/pkg/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type loginController struct {
	authService service.AuthService
}

func NewLoginController(authService service.AuthService) *loginController {
	return &loginController{
		authService: authService,
	}
}

func (lc *loginController) Routes() chi.Router {
	router := chi.NewRouter()

	router.Post("/", lc.handleLoginPOST)

	return router
}

func (lc *loginController) handleLoginPOST(w http.ResponseWriter, r *http.Request) {

	// extract from request
	creds, err := lc.extractCredentialsFrom(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// authenticate
	user, err := lc.findUserWith(creds)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, fmt.Sprintf("Failed to authenticate user %s", creds.User), http.StatusUnauthorized)
		return
	}

	// check minimum access claims
	kind := user.Claims.GetUserKind()
	if kind == "" {
		http.Error(w, fmt.Sprintf("User %s does not claim any user kind", creds.User), http.StatusUnauthorized)
		return
	}

	// sign claims
	token, err := lc.authService.CreateTokenFor(user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// response
	err = json.NewEncoder(w).Encode(exchange.JWTResponse{AccessToken: token.Value})
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (lc *loginController) extractCredentialsFrom(r *http.Request) (exchange.LoginCredentials, error) {
	var creds exchange.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	return creds, err
}

func (lc *loginController) findUserWith(creds exchange.LoginCredentials) (domain.User, error) {

	if creds.Mail != "" && creds.PasswordHash != "" {
		user, err := lc.authService.FindUserByMailAndPsswd(creds.Mail, creds.PasswordHash)
		if err != nil {
			return domain.User{}, err
		}
		return user, nil
	}

	// maybe login with username + passwd if mail missing

	return domain.User{}, errors.New("login credentials are missing values")
}
