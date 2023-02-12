package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"license-manager/pkg/controller/exchange"
	"license-manager/pkg/domain"
	"license-manager/pkg/pkgerrors"
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
	login, err := lc.extractLoginFrom(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// authenticate
	creds, err := lc.findCredentials(login)
	if err != nil {
		if errors.Is(err, pkgerrors.ErrCredsNotFound) {
			http.Error(w, fmt.Sprintf("User %s is unauthorized", login.User), http.StatusUnauthorized)
			return
		} else {
			log.Println(err.Error())
			http.Error(w, fmt.Sprintf("Failed to authenticate user %s", login.User), http.StatusInternalServerError)
			return
		}
	}

	// check minimum access claims
	kind := creds.Claims.GetUserKind()
	if kind == "" {
		http.Error(w, fmt.Sprintf("User %s does not claim any user kind", login.User), http.StatusUnauthorized)
		return
	}

	// sign claims
	token, err := lc.authService.CreateTokenFor(creds)
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

func (lc *loginController) extractLoginFrom(r *http.Request) (exchange.LoginCredentials, error) {
	var creds exchange.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	return creds, err
}

func (lc *loginController) findCredentials(login exchange.LoginCredentials) (domain.Credentials, error) {
	creds, err := lc.authService.FindByUserNameAndPasswordHash(login.User, login.PasswordHash)
	return creds, err
}
