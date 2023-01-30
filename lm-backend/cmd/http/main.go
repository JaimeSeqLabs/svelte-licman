package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type Credentials struct {
	User string `json:"user"`
	Password string `json:"password"`
}

type JWTResponse struct {
	AccessToken string `json:"access_token"`
}

/*

echo "> Login as admin"
curl -i -X POST -d '{ "user": "jaime", "password": "limao" }' http://localhost:8080/login

echo "> Calling protected endpoint"
curl -i http://localhost:8080/greet/jaime -H "Accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2tpbmQiOiJhZG1pbiJ9.xuqrsOgJAfzcWRqOaSRjyDuaGBV_totB67lVX6xM4Qg"

echo "> Login as non admin"
curl -i -X POST -d '{ "user": "alice", "password": "lima" }' http://localhost:8080/login

echo "> Calling protected endpoint"
curl -i http://localhost:8080/greet/alice -H "Accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2tpbmQiOiJ1c2VyIn0.TvcTC5kMMj5oqOjFENgA_I01R1kPWAz_WyMEHm8IT9I"

*/

func main() {

	secret := []byte("secret")
	tokenAuth := jwtauth.New("HS256", secret, nil)

	r := chi.NewRouter()

	// Public
	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger)

		r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world"))
		})

		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {

			var creds Credentials
			err := json.NewDecoder(r.Body).Decode(&creds)
			if err != nil {
				panic(err)
			}

			var kind string
			if creds.User == "jaime" && creds.Password == "limao" {
				kind = "admin"
			} else {
				kind = "user"
			}

			_, tokenStr, err := tokenAuth.Encode(map[string]interface{}{
				"user_kind": kind,
			})
			if err != nil {
				panic(err)
			}
			
			res := JWTResponse{ AccessToken: tokenStr }

			err = json.NewEncoder(w).Encode(res)
			if err != nil {
				panic(err)
			}

		})
	})

	// Private
	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, claims, err := jwtauth.FromContext(r.Context())
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}

				if kind := claims["user_kind"]; kind != "admin" {
					http.Error(w, "user_kind is not admin", http.StatusUnauthorized)
					return
				}

				next.ServeHTTP(w, r)
			})
		})

		r.Get("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
			name := chi.URLParam(r, "name")
			w.Write([]byte(fmt.Sprintf("Hello %s :)", name)))
		})
	})

	http.ListenAndServe(":8080", r)
}