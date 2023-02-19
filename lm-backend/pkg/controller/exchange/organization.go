package exchange

type ListAllOrgsResponse struct {
	Organizations []ListAllOrgsItem `json:"organizations"`
}

type ListAllOrgsItem struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
}

type DescribeOrgResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
	ContactID string `json:"contact_id,omitempty"`
}

type UpdateOrgRequest struct {
	Name string `json:"name"`
	Location string `json:"location"`
}

type CreateOrgRequest struct {
	Name string `json:"name"`
	Location string `json:"location"`
}