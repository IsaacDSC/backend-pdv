package shared_types

type CreateProduct_DTO struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	EnterpriseId string  `json:"enterpriseId"`
	CategoryId   string  `json:"categoryId"`
	Image        string  `json:"image"`
	Price        float32 `json:"price"`
	CreatedAt    string  `json:"createdAt"`
}
