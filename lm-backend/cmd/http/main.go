package main

import (
	"fmt"
	"license-manager/pkg/controller"
	"license-manager/pkg/service/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

/*

echo "> Login as admin"
curl -i -X POST -d '{ "user": "admin", "password": "secret" }' http://localhost:8080/login

echo "> Login as non admin"
curl -i -X POST -d '{ "user": "user", "password": "1234" }' http://localhost:8080/login

echo "> Calling protected endpoint"
curl -i http://localhost:8080/api/is_admin -H "Accept: application/json" -H "Authorization: Bearer <INSERT_TOKEN>"

*/

func main() {

	jwtService := auth.NewJWTService("<this_is_a_secret>") // TODO

	helloController := controller.NewHelloController()
	loginController := controller.NewLoginController(jwtService)

	adminController := controller.NewAdminController(jwtService)

	router := chi.NewRouter()

	// Public endpoints
	router.Route("/", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Mount("/login", loginController.Routes())
		r.Mount("/hello", helloController.Routes())
	})

	// Private endpoints
	router.Route("/api", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(jwtauth.Verifier(jwtService.GetJWTAuth()))
		r.Mount("/is_admin", adminController.Routes())
	})

	addr := "localhost:8080"
	fmt.Printf("> Server is running at 'http://%s'\n", addr)
	http.ListenAndServe(addr, router)

}
