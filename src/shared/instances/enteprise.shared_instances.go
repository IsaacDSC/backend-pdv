package shared_instances

import (
	app_services "backend-pdv/src/app"
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/repositories"
	shared_types "backend-pdv/src/shared/types"
)

func GetInstanceEnterpriseService(product shared_types.CreateEnterprise_DTO) app_services.EnterpriseService {
	entity := &domain.EnterpriseEntity{
		CorporateName: product.CorporateName,
		FantasyName:   product.FantasyName,
		Logo:          product.Logo,
		CPJ:           product.CPJ,
		Token:         product.Token,
	}
	repository := &repositories.EnterpriseRepository{}
	service := app_services.EnterpriseService{
		Entity:     entity,
		Repository: repository,
	}
	return service
}
