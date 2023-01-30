package controller

import (
	"encoding/json"
	"fmt"
	"license-manager/pkg/controller/exchange"
	"license-manager/pkg/service/auth"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type loginController struct {
	jwtService auth.JWTService
}

func NewLoginController(jwtService auth.JWTService) *loginController {
	return &loginController{
		jwtService: jwtService,
	}
}

func (lc *loginController) Routes() chi.Router {
	router := chi.NewRouter()

	router.Post("/", lc.handleLoginPOST)

	return router
}

func (lc *loginController) handleLoginPOST(w http.ResponseWriter, r *http.Request) {
	
	// extract
	creds, err := lc.getCredentialsFrom(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// authenticate
	kind := lc.getUserKind(creds)
	if kind == "" {
		http.Error(w, fmt.Sprintf("User %s is unauthorized", creds.User), http.StatusUnauthorized)
		return
	}

	// sign claims
	token, err := lc.generateTokenClaiming(kind)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// response
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (lc *loginController) getCredentialsFrom(r *http.Request) (exchange.LoginCredentials, error) {
	var creds exchange.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	return creds, err
}

func (lc *loginController) getUserKind(creds exchange.LoginCredentials) string {
	// TODO: this is obviously dumb

	if creds.User == "admin" && creds.Password == "secret" {
		return "admin"
	}

	if creds.User == "user" && creds.Password == "1234" {
		return "user"
	}

	return "" // no kind, unauthorized
}

func (lc *loginController) generateTokenClaiming(userKind string) (exchange.JWTResponse, error) {
	
	tokenStr, err := lc.jwtService.GenTokenFor(
		map[string]any {
			"user_kind": userKind,
		},
	)
	if err != nil {
		return exchange.JWTResponse{}, err
	}
	
	return exchange.JWTResponse{ AccessToken: tokenStr }, nil
}




