package shared

import (
	app_services "backend-pdv/src/app"
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/repositories"
)

type CreateProduct_DTO struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	EnterpriseId string  `json:"enterpriseId"`
	CategoryId   string  `json:"categoryId"`
	Image        string  `json:"image"`
	Price        float32 `json:"price"`
}

func GetInstanceProductService(product CreateProduct_DTO) app_services.ProductService {
	entity := &domain.ProductEntity{
		Name:         product.Name,
		Description:  product.Description,
		EnterpriseId: product.EnterpriseId,
		CategoryId:   product.CategoryId,
		Image:        product.Image,
		Price:        product.Price,
	}
	repository := &repositories.ProductsRepository{}
	service := app_services.ProductService{
		Entity:     entity,
		Repository: repository,
	}
	return service
}
