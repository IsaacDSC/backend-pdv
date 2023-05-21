package shared

import (
	app_services "backend-pdv/src/app"
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/repositories"
)

type CreateCategory_DTO struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	EnterpriseId string `json:"enterpriseId"`
}

func GetInstanceCategoryService(category CreateCategory_DTO) app_services.CategoryService {
	entity := &domain.CategoryEntity{}
	repository := &repositories.CategoryRepository{}
	service := app_services.CategoryService{
		Repository:   repository,
		Entity:       entity,
		Name:         category.Name,
		Description:  category.Description,
		EnterpriseId: category.EnterpriseId,
	}
	return service
}
