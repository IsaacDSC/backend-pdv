package domain

import (
	"backend-pdv/src/domain/aggregate"
	"backend-pdv/src/infra/gateway"
	"database/sql"
)

type RepositoryClient interface {
	GetClientByTelephone(telephone string) (client gateway.Client, err error)
	CreateClient(client gateway.CreateCLientParams) (err error)
}

type ClientEntity struct {
	ID        string
	Name      string
	Telephone string
	Email     string
	Address   aggregate.Address
	CreatedAt string

	Repository RepositoryClient
}

func (this_client *ClientEntity) ValidateClient() (errors []string) {
	if len(this_client.Telephone) < 11 {
		errors = append(errors, "Not-Valid-Telephone")
		return
	}
	return
}

func (this_client *ClientEntity) AlreadyExist() (exist bool, err error) {
	exist = false
	client, err := this_client.Repository.GetClientByTelephone(this_client.Telephone)
	if len(client.ID) > 0 {
		this_client.ID = client.ID
		this_client.Name = client.Name
		this_client.Telephone = client.Telephone
		this_client.CreatedAt = client.CreatedAt.Time.String()
		exist = true
		return
	}
	return
}

func (this_client *ClientEntity) CreateClient() (err error) {
	err = this_client.Repository.CreateClient(gateway.CreateCLientParams{
		Name:      this_client.Name,
		Email:     sql.NullString{String: this_client.Email, Valid: true},
		Telephone: this_client.Telephone,
	})
	return
}
