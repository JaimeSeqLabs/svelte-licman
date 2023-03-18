package service_test

import (
	"license-manager/pkg/repositories/ent-fw/token"
	"license-manager/pkg/repositories/ent-fw/ent/enttest"
	"license-manager/pkg/service"
	"testing"
)

func TestJWTMiddleware(t *testing.T) {

	db := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer db.Close()

	jwtRepo := token_repo.NewJwtTokenEntRepo(db)

	secret := "<test_secret>"
	jwts := service.NewJWTService(secret, jwtRepo)

	// TODO: create middleware
}