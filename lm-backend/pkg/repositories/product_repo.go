package repositories

import "license-manager/pkg/domain"

type ProductRepository interface {
	Save(prod domain.Product) error
	FindAll() []domain.Product
	FindByID(id string) (domain.Product, error)
	FindBySKU(sku string) (domain.Product, error)
	UpdateByID(prod domain.Product) (updated bool, err error)
	DeleteByID(prod domain.Product) error
}