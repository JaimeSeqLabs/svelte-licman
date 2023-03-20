package service_test

import (
	"fmt"
	"license-manager/pkg/repositories/ent-fw/ent/enttest"
	"license-manager/pkg/repositories/ent-fw/token"
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

	secret := "<test_secret>"
	jwts := service.NewJWTService2(secret, jwtRepo)

	mw := service.NewJWTMiddleware(jwts)

	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Use(mw)
		r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hiii")
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	t.Log(rec.Result())

	// TODO: create token, call with token, check no errors
	// then modify token, make call and check error

}