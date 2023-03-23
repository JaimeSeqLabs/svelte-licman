package service_test

import (
	"fmt"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories/ent-fw/ent/enttest"
	"license-manager/pkg/repositories/ent-fw/token"
	user_repo "license-manager/pkg/repositories/ent-fw/user"
	"license-manager/pkg/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

func TestJWTMiddleware(t *testing.T) {

	db := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer db.Close()

	jwtRepo := token_repo.NewJwtTokenEntRepo(db)
	userRepo := user_repo.NewUserEntRepo(db)

	secret := "<test_secret>"
	jwts := service.NewJWTService(secret, jwtRepo)

	authService := service.NewAuthService(userRepo, jwts)

	mw := service.NewJWTMiddleware(jwts)

	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Use(mw)
		r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hiii")
		})
	})

	// TODO: create token, call with token, check no errors
	// then modify token, make call and check error

	usr := domain.User{
		Name: "Jaime",
		Mail: "jaime@mail.com",
		PasswordHash: "<hash>",
		Claims: domain.Claims{
			UserKind: "usr",
		},
	}
	if err := authService.RegisterUser(usr); err != nil {
		panic(err)
	}

	token, err := authService.CreateTokenFor(usr)
	if err != nil {
		panic(err)
	}

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Value))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	t.Log(rec.Result())






}