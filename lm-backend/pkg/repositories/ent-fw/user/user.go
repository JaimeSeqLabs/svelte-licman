package user_repo

import (
	"context"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/ent/user"
)

type userEntRepo struct {
	client *ent.Client
}

func NewUserEntRepo(client *ent.Client) repositories.UserRepository {
	return &userEntRepo{
		client: client,
	}
}

func (repo *userEntRepo) Save(usr domain.User) error {
	_, err := repo.client.User.Create().
		SetUsername(usr.Name).
		SetMail(usr.Mail).
		SetPasswordHash(usr.PasswordHash).
		SetClaims(usr.Claims).
		Save(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (repo *userEntRepo) FindByNameAndMail(name string, mail string) (domain.User, error) {
	usr, err := repo.client.User.Query().
		Where(
			user.And(
				user.UsernameEQ(name),
				user.MailEQ(mail),
			),
		).
		Only(context.TODO())
	if err != nil {
		return domain.User{}, err
	}
	return toEntity(usr), nil
}

func (repo *userEntRepo) FindByMailAndPassword(mail string, passwordHash string) (domain.User, error) {
	usr, err := repo.client.User.Query().
		Where(
			user.And(
				user.MailEQ(mail),
				user.PasswordHashEQ(passwordHash),
			),
		).
		Only(context.TODO())
	if err != nil {
		return domain.User{}, err
	}
	return toEntity(usr), nil
}

func (repo *userEntRepo) Update(usr domain.User) (updated bool, err error) {
	changes, err := repo.client.User.Update().
		Where(
			user.And(
				user.UsernameEQ(usr.Name),
				user.MailEQ(usr.Mail),
			),
		).
		SetPasswordHash(usr.PasswordHash).
		SetClaims(usr.Claims).
		Save(context.TODO())
	if err != nil {
		return false, err
	}
	return (changes > 0), nil
}

func (repo *userEntRepo) DeleteByNameAndMail(name string, mail string) error {
	deleted, err := repo.client.User.Delete().
		Where(user.And(
			user.UsernameEQ(name),
			user.MailEQ(mail),
		)).
		Exec(context.TODO())
	if err != nil {
		return err
	}
	if deleted == 0 {
		return &ent.NotFoundError{}
	}
	return nil
}

func toEntity(dto *ent.User) domain.User {
	return domain.User {
		ID: dto.ID,
		Name: dto.Username,
		Mail: dto.Mail,
		PasswordHash: dto.PasswordHash,
		Claims: domain.Claims(dto.Claims),
	}
}