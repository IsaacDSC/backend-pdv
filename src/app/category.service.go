package app_services

import (
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/repositories"
	shared_types "backend-pdv/src/shared/types"
)

type CategoryService struct {
	Entity     *domain.CategoryEntity
	Repository *repositories.CategoryRepository
}

func (this_service *CategoryService) CreateCategory() (output shared_types.CreateCategory_DTO, errors []string) {
	errors = this_service.Entity.Validate()
	if len(errors) == 0 {
		err := this_service.Repository.CreateCategory(*this_service.Entity)
		output = shared_types.CreateCategory_DTO(*this_service.Entity)
		if err != nil {
			errors = append(errors, "Error: "+err.Error())
		}
	}
	return
}
