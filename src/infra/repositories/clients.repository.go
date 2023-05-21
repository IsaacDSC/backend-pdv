package repositories

import (
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/gateway"
	"context"
	"database/sql"
)

type ClientRepository struct{}

func (*ClientRepository) GetClientByTelephone(telephone string) (client gateway.Client, err error) {
	ctx := context.Background()
	conn := gateway.GetConnection()
	defer conn.Close()
	database := gateway.New(conn)
	client, err = database.FindClient(ctx, telephone)
	return
}

func (*ClientRepository) CreateClient(client domain.ClientEntity) (err error) {
	ctx := context.Background()
	conn := gateway.GetConnection()
	defer conn.Close()
	database := gateway.New(conn)
	err = database.CreateCLient(ctx, gateway.CreateCLientParams{
		ID:           client.ID,
		EnterpriseID: client.EnterpriseId,
		Name:         client.Name,
		Email:        sql.NullString{String: client.Email, Valid: true},
		Password:     sql.NullString{String: client.Password, Valid: true},
		Telephone:    client.Telephone,
		City:         sql.NullString{String: client.Address.City, Valid: true},
		HomeNumber:   sql.NullString{String: client.Address.HomeNumber, Valid: true},
		Neighborhood: sql.NullString{String: client.Address.Neighborhood, Valid: true},
		Cep:          sql.NullString{String: client.Address.Cep, Valid: true},
		Address:      sql.NullString{String: client.Address.Address, Valid: true},
		Observation:  sql.NullString{String: client.Address.Observation, Valid: true},
	})
	return
}
