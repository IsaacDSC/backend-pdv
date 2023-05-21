package shared_instances

import (
	app_services "backend-pdv/src/app"
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/repositories"
	shared_types "backend-pdv/src/shared/types"
)

func GetInstanceCategoryService(category shared_types.CreateCategory_DTO) app_services.CategoryService {
	entity := &domain.CategoryEntity{
		Name:         category.Name,
		Description:  category.Description,
		EnterpriseId: category.EnterpriseId,
	}
	repository := &repositories.CategoryRepository{}
	service := app_services.CategoryService{
		Repository: repository,
		Entity:     entity,
	}
	return service
}
