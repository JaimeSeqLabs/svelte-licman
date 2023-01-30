package controller

import (
	"fmt"
	"license-manager/pkg/service/auth"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type adminController struct {
	jwtService auth.JWTService
}

func NewAdminController(jwtService auth.JWTService) *adminController {
	return &adminController{
		jwtService: jwtService,
	}
}

func (ac *adminController) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ac.handle)

	return router
}

func (ac *adminController) handle(w http.ResponseWriter, r *http.Request) {
	
	claims, err := ac.jwtService.GetClaimsFromCtx(r.Context())
	if err != nil {
		log.Printf("Failed to extract claims from context: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	kind := claims["user_kind"]

	if kind == "admin" {
		fmt.Fprintf(w, "user is admin")
		return
	}

	if kind == "user" {
		fmt.Fprintf(w, "user is not an admin")
		return
	}

	if kind == "" {
		http.Error(w, "user_kind is not valid", http.StatusBadRequest)
		return		
	}
}