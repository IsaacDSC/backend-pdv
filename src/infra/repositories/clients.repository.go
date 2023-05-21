package repositories

import (
	"backend-pdv/src/infra/gateway"
	"context"
)

type ClientRepository struct{}

func (this_client *ClientRepository) GetClientByTelephone(telephone string) (client gateway.Client, err error) {
	ctx := context.Background()
	conn := gateway.GetConnection()
	defer conn.Close()
	database := gateway.New(conn)
	client, err = database.FindClient(ctx, telephone)
	return
}

func (this_client *ClientRepository) CreateClient(client gateway.CreateCLientParams) (err error) {
	ctx := context.Background()
	conn := gateway.GetConnection()
	defer conn.Close()
	database := gateway.New(conn)
	err = database.CreateCLient(ctx, client)
	return
}
