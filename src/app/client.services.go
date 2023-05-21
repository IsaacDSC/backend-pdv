package app_services

import (
	"backend-pdv/src/domain"
)

type ClientService struct {
	Entity    *domain.ClientEntity
	Telephone string
	Name      string
	Email     string
}

func (this_service *ClientService) CreateOnSet() (errors []string) {
	this_service.Entity.Telephone = this_service.Telephone
	errors = this_service.Entity.ValidateClient()
	if len(errors) > 0 {
		return
	}
	exist, err := this_service.Entity.AlreadyExist()
	if err != nil {
		errors = append(errors, "Error-In-Verify-Exist-Client")
		return
	}
	if !exist {
		err := this_service.Entity.CreateClient()
		if err != nil {
			errors = append(errors, "Error-In-Create-Client")
		}
	}
	return
}
