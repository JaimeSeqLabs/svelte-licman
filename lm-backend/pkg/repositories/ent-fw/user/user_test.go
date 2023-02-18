package user_repo_test

import (
	"context"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/ent/enttest"
	"license-manager/pkg/repositories/ent-fw/user"
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
	
	_ = user_repo.NewUserEntRepo(client)
}

func TestCreateData(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	repo := user_repo.NewUserEntRepo(client)

	usr := domain.User{
		Name: "jaime",
		Mail: "jaime@mail.com",
		PasswordHash: "<passwd_hash>",
		Claims: domain.Claims{
			domain.UserKindClaim: "admin",
		},
	}

	err := repo.Save(usr)
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
	
	repo := user_repo.NewUserEntRepo(client)

	target := domain.User{
		Name: "jaime",
		Mail: "jaime@mail.com",
		PasswordHash: "<passwd_hash>",
		Claims: domain.Claims{
			domain.UserKindClaim: "admin",
		},
	}

	err := repo.Save(target)
	if err != nil {
		t.Fatal(err)
	}

	usr1, err := repo.FindByNameAndMail(target.Name, target.Mail)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(target, usr1) {
		t.Fatalf("find by name and mail failed: want %v but got %v", target, usr1)
	}

	usr2, err := repo.FindByMailAndPassword(target.Mail, target.PasswordHash)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(target, usr2) {
		t.Fatalf("find by mail and password hash failed: want %v but got %v", target, usr1)
	}

}

func TestReadDataNotFound(t *testing.T) {
	
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	repo := user_repo.NewUserEntRepo(client)

	_, err := repo.FindByNameAndMail("jaime", "jaime@mail.com")

	if !ent.IsNotFound(err) {
		t.Fatal("'not found error' expected")
	}

}

func TestUpdateData(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	repo := user_repo.NewUserEntRepo(client)

	original := domain.User{
		Name: "jaime",
		Mail: "jaime@mail.com",
		PasswordHash: "<passwd_hash>",
		Claims: domain.Claims{
			domain.UserKindClaim: "admin",
		},
	}

	err := repo.Save(original)
	if err != nil {
		t.Fatal(err)
	}

	updated := original
	updated.PasswordHash = "<new_password_hash>"
	updated.Claims[domain.UserKindClaim] = "normal_user"

	ok, err := repo.Update(updated)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("no changes were detected in the DB")
	}

	got, _ := repo.FindByNameAndMail(original.Name, original.Mail)

	if !reflect.DeepEqual(updated, got) {
		t.Fatalf("updated failed: want %v but got %v", updated, got)
	}

}

func TestUpdateDataNotFound(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	repo := user_repo.NewUserEntRepo(client)

	user := domain.User{
		Name: "jaime",
		Mail: "jaime@mail.com",
		PasswordHash: "<passwd_hash>",
		Claims: domain.Claims{
			domain.UserKindClaim: "admin",
		},
	}

	updated, err := repo.Update(user)

	if err != nil {
		t.Fatal(err)
	}

	if updated {
		t.Fatal("'updated' flag expected to be false")
	}

}

func TestDeleteData(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	repo := user_repo.NewUserEntRepo(client)

	usr := domain.User{
		Name: "jaime",
		Mail: "jaime@mail.com",
		PasswordHash: "<passwd_hash>",
		Claims: domain.Claims{
			domain.UserKindClaim: "admin",
		},
	}

	err := repo.Save(usr)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.DeleteByNameAndMail(usr.Name, usr.Mail)
	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.FindByNameAndMail(usr.Name, usr.Mail)

	if !ent.IsNotFound(err) {
		t.Fatal("'not found' error expected")
	}

}