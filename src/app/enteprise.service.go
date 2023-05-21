package app_services

import (
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/repositories"
	shared_types "backend-pdv/src/shared/types"
)

type EnterpriseService struct {
	Entity     *domain.EnterpriseEntity
	Repository *repositories.EnterpriseRepository
}

func (this_service *EnterpriseService) CreateEnterprise() (output shared_types.OutputCreateEnterprise, errors []string) {
	errors = this_service.Entity.Validate()
	output.EnterpriseId = this_service.Entity.ID
	output.Token = this_service.Entity.Token
	if len(errors) == 0 {
		err := this_service.Repository.CreateEnterprise(*this_service.Entity)
		if err != nil {
			errors = append(errors, "Error: "+err.Error())
		}
	}
	return
}
