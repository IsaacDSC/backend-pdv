package app_services

import (
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/repositories"
	shared_types "backend-pdv/src/shared/types"
)

type ProductService struct {
	Entity     *domain.ProductEntity
	Repository *repositories.ProductsRepository
}

func (this_service *ProductService) CreateProduct() (output shared_types.CreateProduct_DTO, errors []string) {
	errors = this_service.Entity.Validate()
	if len(errors) == 0 {
		err := this_service.Repository.CreateProduct(*this_service.Entity)
		output = shared_types.CreateProduct_DTO(*this_service.Entity)
		if err != nil {
			errors = append(errors, "Error: "+err.Error())
		}
	}
	return
}
