package service

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"license-manager/pkg/controller/exchange"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
	"strings"
	"time"
)

type certificateService struct {
	quotaService QuotaService
    licenseService LicenseService
	prodRepo     repositories.ProductRepository
}

func NewCertificateService(qs QuotaService, ls LicenseService, prodRepo repositories.ProductRepository) CertificateService {
	return &certificateService{qs, ls, prodRepo}
}

func (cs *certificateService) ValidateCertificate(req exchange.LicenseValidateRequest) (domain.License, error) {

	now := time.Now()
	cert := req.Certificate
	data, signature, jsonStr, err := decodeCert(cert)
	if err != nil {
		return domain.License{}, err
	}
	licenseID := data.ID

	if licenseID == "" {
		return domain.License{}, fmt.Errorf("invalid certificate, missing ID")
	}

    license, err := cs.licenseService.DescribeLicense(licenseID)
    if err != nil {
        return domain.License{}, fmt.Errorf("unable to describe license ID %s: %w", licenseID, err)
    }

    licenseSKUs, err := cs.getProdIDsAsSKUs(license.ProductIDs)
    if err != nil {
        return domain.License{}, fmt.Errorf("unable to obtain license ID %s associated products: %w", licenseID, err)
    }

	// check if requested product matches the license product
	var validProduct bool = false
    for _, sku := range licenseSKUs {
		if req.ProductSKU == sku {
			validProduct = true
			break
		}
	}
	if !validProduct {
		return domain.License{}, fmt.Errorf("invalid certificate, unexpected product code %s", req.ProductSKU)
	}

	if license.Status == "suspended" {
		return domain.License{}, fmt.Errorf("invalid certificate, license suspended")
	}

	if license.Status == "deleted" {
		return domain.License{}, fmt.Errorf("invalid certificate, license deleted")
	}

	if now.After(license.ExpirationDate) {
		return domain.License{}, fmt.Errorf("invalid certificate, license expired")
	}

	// check signature from certificate (constant time comparison)
	if !hmac.Equal([]byte(signature), []byte(getHMACSignature(jsonStr, license.Secret))) {
		return domain.License{}, fmt.Errorf("invalid certificate, signature doesn't match")
	}

	return license, nil
}

func (cs *certificateService) CreateCertificate(license domain.License) (string, error) {

	if license.Secret == "" {
		return "", fmt.Errorf("license %s missing a valid secret", license.ID)
	}

	qMap := cs.quotaService.BuildLicenseQuotaMap(license)
	if qMap == nil {
		return "", fmt.Errorf("unable to build quota map for license %s", license.ID)
	}

	skus, err := cs.getProdIDsAsSKUs(license.ProductIDs)
    if err != nil {
        return "", err
    }

	cert := domain.CertificateData{
		ID:             license.ID,
		ProductSKUs:    skus,
		Quotas:         qMap,
		ActivationDate: license.ActivationDate,
		ExpirationDate: license.ExpirationDate,
	}

	return encodeCert(cert, license.Secret)

}

func (cs *certificateService) getProdIDsAsSKUs(ids []string) ([]string, error) {

    prods, err := cs.prodRepo.FindByIDs(ids)
	if err != nil {
		return []string{}, err
	}

	skus := []string{}
	for _, p := range prods {
		skus = append(skus, p.SKU)
	}

    return skus, nil
}

func (cs *certificateService) CreateCertificateFile(_ domain.License) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (cs *certificateService) DescribeCertificateData(cert string) (domain.CertificateData, error) {
	data, _, _, err := decodeCert(cert)
	return data, err
}

func encodeCert(cert domain.CertificateData, secret string) (string, error) {

	payload, err := json.Marshal(cert)
	if err != nil {
		return "", err
	}

	signature := getHMACSignature(string(payload), secret)

	// {"ver":version}.payload.signature
	str := fmt.Sprintf(`{"ver":%s}.%s.%s`, domain.LicenseCurrentVersion, string(payload), signature)

	return base64.StdEncoding.EncodeToString([]byte(str)), nil
}

// decodeCert returns CertificateData, signature, json payload as str and error if any
func decodeCert(cert string) (data domain.CertificateData, signature string, jsonStr string, err error) {

	bytes, err := base64.StdEncoding.DecodeString(cert)
	if err != nil {
		return
	}
	str := string(bytes)

	parts := strings.Split(str, ".")
	if parts == nil || len(parts) != 3 {
		err = fmt.Errorf("cannot decode certificate %s, missing parts", cert)
		return
	}

	header := map[string]string{}

	err = json.Unmarshal([]byte(parts[0]), &header)
	if err != nil {
		return
	}

	if ver := header["ver"]; ver != domain.LicenseCurrentVersion {
		err = fmt.Errorf("cannot decode certificate, invalid version %s", ver)
		return
	}

	jsonStr = parts[1]
	signature = parts[2]

	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return
	}

	err = nil
	return
}

func getHMACSignature(payload, secret string) string {
	
	mac := hmac.New(sha1.New, []byte(secret))
	
	mac.Write([]byte(payload))
	
	return hex.EncodeToString(mac.Sum(nil))
}