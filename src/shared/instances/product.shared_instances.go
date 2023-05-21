package shared_instances

import (
	app_services "backend-pdv/src/app"
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/repositories"
	shared_types "backend-pdv/src/shared/types"
)

func GetInstanceProductService(product shared_types.CreateProduct_DTO) app_services.ProductService {
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
