package shared_types

type CreateCategory_DTO struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	EnterpriseId string `json:"enterpriseId"`
	CreatedAt    string `json:"createdAt"`
}
