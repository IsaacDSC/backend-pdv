package tests_repository

import (
	"backend-pdv/src/infra/environments"
	"backend-pdv/src/infra/gateway"
	"backend-pdv/src/infra/repositories"
	"database/sql"
	"testing"

	"github.com/savsgio/gotils/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	input_client      gateway.CreateCLientParams
	client_repository repositories.ClientRepository
)

func init() {
	environments.StartEnv()
	input_client = gateway.CreateCLientParams{
		Name:         "name_" + uuid.V4()[0:5],
		Email:        sql.NullString{String: uuid.V4()[0:5] + "@gmail.com", Valid: true},
		Telephone:    uuid.V4()[0:10],
		HomeNumber:   sql.NullString{String: "200", Valid: true},
		Cep:          sql.NullString{String: "27325110", Valid: true},
		Neighborhood: sql.NullString{String: "FAKE", Valid: true},
		Address:      sql.NullString{String: "Rua Fake", Valid: true},
		Observation:  sql.NullString{String: "FINAL DA RUA", Valid: true},
		City:         "Barra Mansa",
	}
	client_repository = repositories.ClientRepository{}
}

func TestCreateClient(t *testing.T) {
	err := client_repository.CreateClient(input_client)
	assert.NoError(t, err)
}

func TestGetClient(t *testing.T) {
	client, err := client_repository.GetClientByTelephone(input_client.Telephone)
	assert.NoError(t, err)
	assert.Equal(t, client.Name, input_client.Name)
	assert.Equal(t, client.Telephone, input_client.Telephone)
}
