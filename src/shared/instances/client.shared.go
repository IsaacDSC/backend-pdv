package shared_instances

import (
	app_services "backend-pdv/src/app"
	"backend-pdv/src/domain"
	"backend-pdv/src/domain/aggregate"
	"backend-pdv/src/infra/repositories"
	shared_types "backend-pdv/src/shared/types"
)

func GetInstanceClientService(createOnSet shared_types.ClientCreateOnSet_DTO) app_services.ClientService {
	repository := &repositories.ClientRepository{}
	entity := &domain.ClientEntity{
		EnterpriseId: createOnSet.EnterpriseId,
		Telephone:    createOnSet.Telephone,
		Name:         createOnSet.Name,
		Email:        createOnSet.Email,
		Address:      aggregate.Address(createOnSet.Address),
	}
	service := app_services.ClientService{
		Entity:     entity,
		Repository: repository,
	}
	return service
}
