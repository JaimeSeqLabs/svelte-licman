package main

import (
	"context"
	"fmt"
	"license-manager/pkg/controller"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/license"
	"license-manager/pkg/repositories/ent-fw/organization"
	"license-manager/pkg/repositories/ent-fw/product"
	"license-manager/pkg/repositories/ent-fw/token"
	"license-manager/pkg/repositories/ent-fw/user"
	"license-manager/pkg/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"

	_ "github.com/mattn/go-sqlite3"
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

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run migrations
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	jwtRepo := token_repo.NewJwtTokenEntRepo(client)
	userRepo := user_repo.NewUserEntRepo(client)
	orgRepo := organization_repo.NewOrganizationEntRepo(organization_repo.WithEntClient(client))
	licenseRepo := license_repo.NewLicenseEntRepo(client)
	prodRepo := product_repo.NewProductEntRepo(client)

	jwtService := service.NewJWTService("<this_is_a_secret>", jwtRepo)
	authService := service.NewAuthService(userRepo, jwtService)
	licenseService := service.NewLicenseService(licenseRepo, orgRepo, prodRepo)
	certificateService := service.NewCertificateService(licenseService, prodRepo)

	helloController := controller.NewHelloController()
	loginController := controller.NewLoginController(authService)
	adminController := controller.NewAdminController(jwtService)
	organizationController := controller.NewOrganizationController(orgRepo)
	productController := controller.NewProductController(prodRepo)
	licenseController := controller.NewLicenseController(licenseService, certificateService)
	validationController := controller.NewValidationController(certificateService)

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

		r.Mount("/organizations", organizationController.Routes())
		r.Mount("/licenses", licenseController.Routes())
		r.Mount("/products", productController.Routes())
		r.Mount("/validate", validationController.Routes())
	})

	addr := "localhost:8080"
	fmt.Printf("> Server is running at 'http://%s'\n", addr)
	http.ListenAndServe(addr, router)

}
