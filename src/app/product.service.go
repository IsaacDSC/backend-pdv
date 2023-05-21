package app_services

import (
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/repositories"
)

type ProductService struct {
	Entity     *domain.ProductEntity
	Repository *repositories.ProductsRepository
}

func (this_service *ProductService) CreateProduct() (errors []string) {
	errors = this_service.Entity.Validate()
	if len(errors) == 0 {
		err := this_service.Repository.CreateProduct(*this_service.Entity)
		if err != nil {
			errors = append(errors, "Error: "+err.Error())
		}
	}
	return
}
