package shared

import (
	app_services "backend-pdv/src/app"
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/repositories"
)

type ClientCreateOnSet_DTO struct {
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Name      string `json:"name"`
}

func GetInstanceClientService(createOnSet ClientCreateOnSet_DTO) app_services.ClientService {
	repository := &repositories.ClientRepository{}
	entity := &domain.ClientEntity{Repository: repository}
	service := app_services.ClientService{
		Entity:    entity,
		Telephone: createOnSet.Telephone,
		Name:      createOnSet.Name,
		Email:     createOnSet.Email,
	}
	return service
}
