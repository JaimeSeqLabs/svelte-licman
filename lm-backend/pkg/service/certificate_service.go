package service

import (
	"license-manager/pkg/domain"
	"license-manager/pkg/controller/exchange"
)
	

type CertificateService interface {
	
	ValidateCertificate(exchange.LicenseValidateRequest) (domain.License, error)

    CreateCertificate(domain.License) (string, error)

    CreateCertificateFile(domain.License) (string, error)

    DescribeCertificateData(cert string) (domain.CertificateData, error)

}