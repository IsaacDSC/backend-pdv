package app_services

import (
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/repositories"
	shared_types "backend-pdv/src/shared/types"
)

type ClientService struct {
	Entity     *domain.ClientEntity
	Repository *repositories.ClientRepository
}

func (this_service *ClientService) CreateOnSet() (client shared_types.ClientCreateOnSet_DTO, errors []string) {
	errors = this_service.Entity.ValidateClient()
	if len(errors) > 0 {
		return
	}
	exist, err := this_service.alreadyExist()
	if err != nil {
		errors = append(errors, "Error-In-Verify-Exist-Client")
		errors = append(errors, err.Error())
		return
	}
	if exist {
		errors = append(errors, "User-Already-Exist")
		return
	}
	err = this_service.createClient()
	if err != nil {
		errors = append(errors, "Error-In-Create-Client")
	}
	client = shared_types.ClientCreateOnSet_DTO{
		ID:           this_service.Entity.ID,
		EnterpriseId: this_service.Entity.EnterpriseId,
		Name:         this_service.Entity.Name,
		Telephone:    this_service.Entity.Telephone,
		Email:        this_service.Entity.Email,
		Address: struct {
			City         string "json:\"city\""
			HomeNumber   string "json:\"homeNumber\""
			Neighborhood string "json:\"neighborhood\""
			Cep          string "json:\"cep\""
			Address      string "json:\"address\""
			Observation  string "json:\"observation\""
		}(this_service.Entity.Address),
		CreatedAt: this_service.Entity.CreatedAt,
	}
	return
}

func (this_service *ClientService) alreadyExist() (exist bool, err error) {
	exist = false
	client, err := this_service.Repository.GetClientByTelephone(this_service.Entity.Telephone)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			err = nil
		}
		return
	}
	if len(client.ID) > 0 {
		exist = true
		return
	}
	return
}

func (this_service *ClientService) createClient() (err error) {
	err = this_service.Repository.CreateClient(*this_service.Entity)
	return
}
