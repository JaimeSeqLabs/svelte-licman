package exchange

import "license-manager/pkg/domain"

type ListAllOrgsResponse struct {
	Organizations []ListAllOrgsItem `json:"organizations"`
}

type ListAllOrgsItem struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Mail    string `json:"mail"`
	Country string `json:"country"`
}

type DescribeOrgResponse domain.Organization

type UpdateOrgRequest domain.Organization

type CreateOrgRequest domain.Organization