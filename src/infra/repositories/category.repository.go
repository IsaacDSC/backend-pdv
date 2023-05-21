package repositories

import (
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/gateway"
	"context"
	"database/sql"
)

type CategoryRepository struct{}

func (this_repository *CategoryRepository) CreateCategory(category domain.CategoryEntity) (err error) {
	ctx := context.Background()
	conn := gateway.GetConnection()
	defer conn.Close()
	database := gateway.New(conn)
	err = database.CreateCategory(ctx, gateway.CreateCategoryParams{
		ID:           category.ID,
		Name:         category.Name,
		Description:  sql.NullString{String: category.Description, Valid: true},
		EnterpriseID: sql.NullString{String: category.EnterpriseId, Valid: true},
	})
	return
}
