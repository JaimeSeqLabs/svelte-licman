package credentials_repo_test

import (
	"context"
	"license-manager/pkg/domain"
	"license-manager/pkg/pkgerrors"
	"license-manager/pkg/repositories/ent-fw/credentials"
	"license-manager/pkg/repositories/ent-fw/ent/enttest"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestNewCredentialsEntRepo(t *testing.T) {
	
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}

	_ = credentials_repo.NewCredentialsEntRepo(client)

}

func TestSave(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}

	repo := credentials_repo.NewCredentialsEntRepo(client)

	err := repo.Save(domain.Credentials{
		UserName: "jaime",
		PasswordHash: "<hash>",	
		Claims: domain.Claims{
			"user_kind": "admin",
		},
	})

	if err != nil {
		t.Fatal(err)
	}

}

func TestFindByNameAndHash(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}

	repo := credentials_repo.NewCredentialsEntRepo(client)

	creds := domain.Credentials{
		UserName: "jaime",
		PasswordHash: "<hash>",	
		Claims: domain.Claims{
			"user_kind": "admin",
		},
	}

	err := repo.Save(creds)
	if err != nil {
		t.Fatal(err)
	}

	res, err := repo.FindByUserNameAndPasswordHash(creds.UserName, creds.PasswordHash)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(creds, res) {
		t.Fatalf("want %v but got %v", creds, res)
	}
}

func TestUpdate(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}

	repo := credentials_repo.NewCredentialsEntRepo(client)

	creds := domain.Credentials{
		UserName: "jaime",
		PasswordHash: "<hash>",	
		Claims: domain.Claims{
			"user_kind": "admin",
		},
	}

	err := repo.Save(creds)
	if err != nil {
		t.Fatal(err)
	}

	creds.Claims["readonly"] = "true"

	err = repo.Update(creds)
	if err != nil {
		t.Fatal(err)
	}

	res, err := repo.FindByUserNameAndPasswordHash(creds.UserName, creds.PasswordHash)
	if err != nil {
		t.Fatal(err)
	}

	if res.Claims["readonly"] != "true" {
		t.Fatal("entity field was not updated")
	}

}

func TestMergeClaims(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}

	repo := credentials_repo.NewCredentialsEntRepo(client)

	originalClaims := domain.Claims {
		"const_claim": "const",
		"user_kind": "admin",
	}

	updateClaims := domain.Claims {
		"user_kind": "user",
		"new_claim": "new",
	}

	targetClaims := domain.Claims {
		"const_claim": "const",
		"user_kind": "user",
		"new_claim": "new",
	}

	creds := domain.Credentials{
		UserName: "jaime",
		PasswordHash: "<hash>",	
		Claims: originalClaims,
	}

	err := repo.Save(creds)
	if err != nil {
		t.Fatal(err)
	}

	res, err := repo.MergeClaimsFor(creds.UserName, creds.PasswordHash, updateClaims)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(res ,targetClaims) {
		t.Fatalf("want %v but got %v", targetClaims, res)
	}

	updatedCreds, err := repo.FindByUserNameAndPasswordHash(creds.UserName, creds.PasswordHash)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(updatedCreds.Claims, targetClaims) {
		t.Fatalf("want %v but got %v", targetClaims, updatedCreds.Claims)
	}

}

func TestDelete(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}

	repo := credentials_repo.NewCredentialsEntRepo(client)

	creds := domain.Credentials{
		UserName: "jaime",
		PasswordHash: "<hash>",	
		Claims: domain.Claims{
			"user_kind": "admin",
		},
	}

	repo.Save(creds)

	err := repo.DeleteByUserNameAndPasswordHash(creds.UserName, creds.PasswordHash)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.DeleteByUserNameAndPasswordHash(creds.UserName, creds.PasswordHash)

	if err != pkgerrors.ErrCredsNotFound {
		t.Fatal("expected 'pkgerrors.ErrCredsNotFound'")
	}
}




















