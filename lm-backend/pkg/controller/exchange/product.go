package exchange

import (
	"time"
)

type CreateProductRequest struct {
	SKU                 string `json:"sku"`
	Name                string `json:"name"`
	InstallInstructions string `json:"install_instructions"`
}

type ListAllProductsResponse struct {
	Products []ListAllProductsItem `json:"products"`
}

type DescribeProductResponse struct {
	ID                  string 	`json:"id"`
	SKU                 string 	`json:"sku"`
	Name                string 	`json:"name"`
	InstallInstructions string 	`json:"install_instructions"`
	LicenseCount        int 	`json:"license_count"`
	
	DateCreated time.Time		`json:"date_created"`
	LastUpdated time.Time		`json:"last_updated"`	
}

type ListAllProductsItem DescribeProductResponse

type UpdateProductRequest CreateProductRequest