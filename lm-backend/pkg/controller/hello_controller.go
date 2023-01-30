package controller

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type helloController struct {
	// services
}

func NewHelloController() *helloController {
	return &helloController{}
}

func (hc *helloController) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		hc.greet(name, w)
	})

	return router
}

func (hc *helloController) greet(name string, w io.Writer) error {
	fmt.Fprintf(w, "Hello %s :)", name)
	return nil
}
