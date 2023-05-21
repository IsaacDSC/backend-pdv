package app_services

import (
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/repositories"
)

type CategoryService struct {
	Name         string
	Description  string
	EnterpriseId string
	Entity       *domain.CategoryEntity
	Repository   *repositories.CategoryRepository
}

func (this_service *CategoryService) CreateCategory() (errors []string) {
	this_service.Entity.Name = this_service.Name
	this_service.Entity.Description = this_service.Description
	this_service.Entity.EnterpriseId = this_service.EnterpriseId
	errors = this_service.Entity.Validate()
	if len(errors) == 0 {
		err := this_service.Repository.CreateCategory(*this_service.Entity)
		if err != nil {
			errors = append(errors, "Error: "+err.Error())
		}
	}
	return
}
