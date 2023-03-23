package token_repo_test

import (
	"context"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/ent/enttest"
	"license-manager/pkg/repositories/ent-fw/token"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateRepo(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	_ = token_repo.NewJwtTokenEntRepo(client)
}

func TestCreateData(t *testing.T) {
	
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	repo := token_repo.NewJwtTokenEntRepo(client)

	usr := createUserX(client)

	token := domain.Token {
		Value: "<token_value>",
		Revoked: false,
		Claims: domain.Claims {
			UserKind: "admin",
		},
		IssuerID: usr.ID,
	}

	_, err := repo.Save(token)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadData(t *testing.T) {
	
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	repo := token_repo.NewJwtTokenEntRepo(client)

	usr := createUserX(client)

	token := domain.Token {
		Value: "<token_value>",
		Revoked: true,
		Claims: domain.Claims {
			UserKind: "admin",
		},
		IssuerID: usr.ID,
	}

	_, _ = repo.Save(token)

	got, err := repo.FindByToken(token.Value)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(token, got) {
		t.Fatalf("want %v but got %v", token, got)
	}

	claims, err := repo.FindClaimsByToken(token.Value)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(token.Claims, claims) {
		t.Fatalf("want %v but got %v", token.Claims, claims)
	}

	revoked, err := repo.IsRevoked(token.Value)
	if err != nil {
		t.Fatal(err)
	}

	if token.Revoked != revoked {
		t.Fatalf("want %v but got %v", token.Revoked, revoked)
	}

}

func TestDeleteData(t *testing.T) {
	
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	repo := token_repo.NewJwtTokenEntRepo(client)

	usr := createUserX(client)

	token := domain.Token {
		Value: "<token_value>",
		Revoked: false,
		Claims: domain.Claims {
			UserKind: "admin",
		},
		IssuerID: usr.ID,
	}

	_, _ = repo.Save(token)

	err := repo.Delete(token.Value)
	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.FindByToken(token.Value)

	if !ent.IsNotFound(err) {
		t.Fatal("'not found error' expected")
	}

	err = repo.Delete(token.Value)

	if !ent.IsNotFound(err) {
		t.Fatal("'not found error' expected when trying to delete a non-existent token")
	}

}

func createUserX(client *ent.Client) *ent.User {
	usr := client.User.Create().
		SetUsername("user").
		SetMail("mail").
		SetPasswordHash("hash").
		SetClaims(domain.Claims{}).
		SaveX(context.TODO())
	return usr
}