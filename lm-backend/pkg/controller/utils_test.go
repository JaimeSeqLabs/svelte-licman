package controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/ent/enttest"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func makeClient(t *testing.T) *ent.Client {
	return enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
}

func readJSON(r *httptest.ResponseRecorder, ptr any) {

	body, err := io.ReadAll(r.Result().Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, ptr)
	if err != nil {
		panic(err)
	}
}

// callChiRouter exchanges request and response with a chi Router.
//
// example: res := callChiRouter(r, http.MethodGet, "/", nil, &responseStruct)
func callChiRouter(router chi.Router, method, path string, requestObj any, responseObjPtr any) *http.Response {
	
	var body io.Reader = nil

	if requestObj != nil {
		data, err := json.Marshal(requestObj)
		if err != nil {
			panic(err)
		}
		body = bytes.NewReader(data)
	}

	req := httptest.NewRequest(method, path, body)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if responseObjPtr != nil {
		readJSON(rec, responseObjPtr)
	}

	return rec.Result()
	
}