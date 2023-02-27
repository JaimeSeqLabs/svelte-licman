package exchange

type LicenseValidateRequest struct {

	Certificate string `json:"certificate"`
	ProductSKU string `json:"product"`

}

type LicenseValidateResponse struct {
	
	LicenseID string `json:"id"`
	LicenseVersion string `json:"version"`
	Quotas map[string]string `json:"quotas"`

}
