package exchange

type LicenseValidateRequest struct {

	Certificate string `json:"certificate"`
	ProductSKU string `json:"product"`

}
