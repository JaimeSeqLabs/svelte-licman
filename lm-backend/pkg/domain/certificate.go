package domain

import "time"

type CertificateData struct {
	ID             string            `json:"id"`
	ProductSKUs    []string          `json:"prod"`
	Quotas         map[string]string `json:"quotas"`
	ActivationDate time.Time         `json:"act"`
	ExpirationDate time.Time         `json:"exp"`
}
