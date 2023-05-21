package domain

import (
	"backend-pdv/src/domain/aggregate"
	"backend-pdv/src/infra/gateway"
	"time"

	"github.com/savsgio/gotils/uuid"
)

type RepositoryClient interface {
	GetClientByTelephone(telephone string) (client gateway.Client, err error)
	CreateClient(client gateway.CreateCLientParams) (err error)
}

type ClientEntity struct {
	ID           string
	EnterpriseId string
	Name         string
	Telephone    string
	Email        string
	Password     string
	Address      aggregate.Address
	CreatedAt    string
}

func (this_client *ClientEntity) ValidateClient() (errors []string) {
	this_client.ID = uuid.V4()
	this_client.CreatedAt = time.Now().Format(time.RFC3339)
	if len(this_client.Telephone) < 11 {
		errors = append(errors, "Not-Valid-Telephone")
	}
	if len(this_client.Address.City) == 0 {
		errors = append(errors, "Required-Complete-Address")
	}
	if len(this_client.EnterpriseId) == 0 {
		errors = append(errors, "Required-Field-EnterpriseId")
	}
	return
}
