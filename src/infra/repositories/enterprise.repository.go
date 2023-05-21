package repositories

import (
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/gateway"
	"context"
	"database/sql"
)

type EnterpriseRepository struct{}

func (this_service *EnterpriseRepository) CreateEnterprise(enterprise domain.EnterpriseEntity) (err error) {
	ctx := context.Background()
	conn := gateway.GetConnection()
	defer conn.Close()
	database := gateway.New(conn)
	err = database.CreateEnterprise(ctx, gateway.CreateEnterpriseParams{
		ID:            enterprise.ID,
		CorporateName: sql.NullString{String: enterprise.CorporateName, Valid: true},
		FantasyName:   sql.NullString{String: enterprise.FantasyName, Valid: true},
		Logo:          sql.NullString{String: enterprise.Logo, Valid: true},
		Cnpj:          sql.NullString{String: enterprise.CPJ, Valid: true},
		Token:         sql.NullString{String: enterprise.Token, Valid: true},
	})
	return
}
