package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"license-manager/pkg/config"
	"license-manager/pkg/controller"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/license"
	"license-manager/pkg/repositories/ent-fw/organization"
	"license-manager/pkg/repositories/ent-fw/product"
	"license-manager/pkg/repositories/ent-fw/token"
	"license-manager/pkg/repositories/ent-fw/user"
	"license-manager/pkg/service"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"

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

	cfg := getCfg()

	client := getEntClient(cfg)
	defer client.Close()

	jwtRepo := token_repo.NewJwtTokenEntRepo(client)
	userRepo := user_repo.NewUserEntRepo(client)
	orgRepo := organization_repo.NewOrganizationEntRepo(organization_repo.WithEntClient(client))
	licenseRepo := license_repo.NewLicenseEntRepo(client)
	prodRepo := product_repo.NewProductEntRepo(client)

	jwtService := service.NewJWTService(cfg.JWTSecret, jwtRepo)
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

	// initial data
	registerAdmin(cfg, authService)

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

	fmt.Println("> Application endpoints:")
	chi.Walk(router, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("   [%-7s] '%s'\n", method, route)
		return nil
	})

	addr := cfg.ServerAddress
	fmt.Printf("> Server is running at 'http://%s'\n", addr)
	http.ListenAndServe(addr, router)

}

func getCfg() config.ApplicationCfg {

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cfg := config.LoadCfg(cwd)

	fmt.Printf("> Using config file %s\n", viper.ConfigFileUsed())

	return cfg
}

func getEntClient(cfg config.ApplicationCfg) *ent.Client {

	client, err := ent.Open(cfg.DBDriver, cfg.DBURL)
	if err != nil {
		log.Fatalf("failed opening connection to %s %s, reason: %v", cfg.DBDriver, cfg.DBURL, err)
	}
	
	// Run migrations
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func registerAdmin(cfg config.ApplicationCfg, authService service.AuthService) {

	if cfg.AdminMail == "" || cfg.AdminPassword == "" {
		panic("initial admin mail and password needed to start the server")
	}

	sha := sha256.New()
	sha.Write([]byte(cfg.AdminPassword))
	hash := hex.EncodeToString(sha.Sum(nil))

	err := authService.RegisterUser(domain.User{
		Name: "Admin",
		Mail: cfg.AdminMail,
		PasswordHash: hash,
		Claims: domain.Claims{
			domain.UserKindClaim: "admin",
		},
	})
	if err != nil {
		panic(fmt.Errorf("cannot create initial admin user, failure to register: %w", err))
	}

}
