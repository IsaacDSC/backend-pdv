package repositories

import (
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/gateway"
	"context"
	"database/sql"
)

type ProductsRepository struct{}

func (this_repository *ProductsRepository) CreateProduct(product domain.ProductEntity) (err error) {
	ctx := context.Background()
	conn := gateway.GetConnection()
	defer conn.Close()
	database := gateway.New(conn)
	err = database.CreateProduct(ctx, gateway.CreateProductParams{
		ID:           product.ID,
		CategoryID:   product.CategoryId,
		EnterpriseID: product.EnterpriseId,
		Name:         product.Name,
		Image:        sql.NullString{String: product.Image, Valid: true},
		Description:  sql.NullString{String: product.Description, Valid: true},
		Price:        sql.NullFloat64{Float64: float64(product.Price), Valid: true},
	})
	return
}
