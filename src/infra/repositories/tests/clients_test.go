package tests_repository

import (
	"backend-pdv/src/domain"
	"backend-pdv/src/domain/aggregate"
	"backend-pdv/src/infra/environments"
	"backend-pdv/src/infra/repositories"
	"testing"

	"github.com/savsgio/gotils/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	client_repository repositories.ClientRepository
	address           aggregate.Address
	entity            domain.ClientEntity
)

func init() {
	address = aggregate.Address{
		HomeNumber:   "200",
		Cep:          "27325110",
		Neighborhood: "FAKE",
		Address:      "FAKE",
		Observation:  "NOT",
		City:         "Barra Mansa",
	}
	entity = domain.ClientEntity{
		ID:        uuid.V4(),
		Name:      "name_" + uuid.V4()[0:5],
		Email:     uuid.V4()[0:5] + "@gmail.com",
		Telephone: uuid.V4()[0:10],
		Address:   address,
	}
	environments.StartEnv()
	client_repository = repositories.ClientRepository{}
}

func TestCreateClient(t *testing.T) {
	err := client_repository.CreateClient(entity)
	assert.NoError(t, err)
}

func TestGetClient(t *testing.T) {
	client, err := client_repository.GetClientByTelephone(entity.Telephone)
	assert.NoError(t, err)
	assert.Equal(t, client.Name, entity.Name)
	assert.Equal(t, client.Telephone, entity.Telephone)
}
