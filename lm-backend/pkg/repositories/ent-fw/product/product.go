package product_repo

import (
	"context"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/ent/product"
)

type productEntRepo struct {
	client *ent.Client
}

func NewProductEntRepo(client *ent.Client) repositories.ProductRepository {
	return &productEntRepo{
		client: client,
	}
}

func (repo * productEntRepo) Save(prod domain.Product) error {
	_, err := repo.client.Product.Create().
		SetSku(prod.SKU).
		SetName(prod.Name).
		SetInstallInstr(prod.InstallInstructions).
		SetLicenseCount(prod.LicenseCount).
		// dates set by default handlers
		Save(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (repo * productEntRepo) FindAll() []domain.Product {
	
	all, err := repo.client.Product.Query().All(context.TODO())
	
	if err != nil || len(all) == 0 {
		return []domain.Product{}
	}

	prods := make([]domain.Product, len(all))

	for i, dto := range all {
		prods[i] = toEntity(dto)
	}

	return prods
}

func (repo * productEntRepo) FindByID(id string) (domain.Product, error) {
	
	prod, err := repo.client.Product.Get(context.TODO(), id)
	if err != nil {
		return domain.Product{}, err
	}

	return toEntity(prod), nil
}

func (repo * productEntRepo) FindByIDs(ids []string) ([]domain.Product, error) {

	prods, err := repo.client.Product.Query().
		Where(product.IDIn(ids...)).
		All(context.TODO())
	
	if err != nil {
		return nil, err
	}

	return toEntitySlice(prods), nil
}

func (repo * productEntRepo) FindBySKU(sku string) (domain.Product, error) {
	
	prod, err := repo.client.Product.Query().
		Where(
			product.SkuEQ(sku),
		).
		Only(context.TODO())
	if err != nil {
		return domain.Product{}, err
	}

	return toEntity(prod), nil
}

func (repo * productEntRepo) UpdateByID(prod domain.Product) (bool, error) {
	
	updated, err := repo.client.Product.UpdateOneID(prod.ID).
		SetSku(prod.SKU).
		SetName(prod.Name).
		SetInstallInstr(prod.InstallInstructions).
		SetLicenseCount(prod.LicenseCount).
		Save(context.TODO())

	if err != nil {
		return false, err
	}

	if updated == nil {
		return false, nil
	}

	return true, nil
}

func (repo * productEntRepo) DeleteByID(prod domain.Product) error {
	return repo.client.Product.DeleteOneID(prod.ID).Exec(context.TODO())
}

func toEntity(dto *ent.Product) domain.Product {
	return domain.Product {
		ID: dto.ID,
		SKU: dto.Sku,
		Name: dto.Name,
		InstallInstructions: dto.InstallInstr,
		LicenseCount: dto.LicenseCount,
		DateCreated: dto.DateCreated,
		LastUpdated: dto.LastUpdated,
	}
}

func toEntitySlice(dtos []*ent.Product) []domain.Product {
	
	products := make([]domain.Product, len(dtos))

	for i, dto := range dtos {
		products[i] = toEntity(dto)
	}

	return products
}