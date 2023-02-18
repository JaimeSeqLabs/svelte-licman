package repositories

import "license-manager/pkg/domain"

type UserRepository interface {
	Save(usr domain.User) error
	FindByNameAndMail(name, mail string) (domain.User, error)
	FindByMailAndPassword(name, passwordHash string) (domain.User, error)
	Update(usr domain.User) (updated bool, err error)
	DeleteByNameAndMail(name, mail string) error
}