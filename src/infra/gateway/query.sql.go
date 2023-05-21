// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package gateway

import (
	"context"
	"database/sql"
)

const createCLient = `-- name: CreateCLient :exec
INSERT INTO "clients" (
        "name",
        "email",
        "password",
        "telephone",
        "home_number",
        city,
        neighborhood,
        cep,
        "address",
        observation,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, Now())
`

type CreateCLientParams struct {
	Name         string
	Email        sql.NullString
	Password     sql.NullString
	Telephone    string
	HomeNumber   sql.NullString
	City         string
	Neighborhood sql.NullString
	Cep          sql.NullString
	Address      sql.NullString
	Observation  sql.NullString
}

func (q *Queries) CreateCLient(ctx context.Context, arg CreateCLientParams) error {
	_, err := q.db.ExecContext(ctx, createCLient,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Telephone,
		arg.HomeNumber,
		arg.City,
		arg.Neighborhood,
		arg.Cep,
		arg.Address,
		arg.Observation,
	)
	return err
}

const createCategory = `-- name: CreateCategory :exec
INSERT INTO "categories" (
        "name",
        "description",
        "enterprise_id",
        "updated_at"
    )
VALUES($1, $2, $3, Now())
`

type CreateCategoryParams struct {
	Name         string
	Description  sql.NullString
	EnterpriseID sql.NullString
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, createCategory, arg.Name, arg.Description, arg.EnterpriseID)
	return err
}

const createEnterprise = `-- name: CreateEnterprise :exec
INSERT INTO "enterprise" (
        id,
        corporate_name,
        fantasy_name,
        logo,
        cnpj,
        token,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6, Now())
`

type CreateEnterpriseParams struct {
	ID            string
	CorporateName sql.NullString
	FantasyName   sql.NullString
	Logo          sql.NullString
	Cnpj          sql.NullString
	Token         sql.NullString
}

func (q *Queries) CreateEnterprise(ctx context.Context, arg CreateEnterpriseParams) error {
	_, err := q.db.ExecContext(ctx, createEnterprise,
		arg.ID,
		arg.CorporateName,
		arg.FantasyName,
		arg.Logo,
		arg.Cnpj,
		arg.Token,
	)
	return err
}

const createProduct = `-- name: CreateProduct :exec
INSERT INTO "products" (
        "category_id",
        "enterprise_id",
        "name",
        "description",
        "image",
        "price",
        "updated_at"
    )
VALUES($1, $2, $3, $4, $5, $6, Now())
`

type CreateProductParams struct {
	CategoryID   string
	EnterpriseID string
	Name         string
	Description  sql.NullString
	Image        sql.NullString
	Price        sql.NullFloat64
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) error {
	_, err := q.db.ExecContext(ctx, createProduct,
		arg.CategoryID,
		arg.EnterpriseID,
		arg.Name,
		arg.Description,
		arg.Image,
		arg.Price,
	)
	return err
}

const findClient = `-- name: FindClient :one
SELECT id, name, email, password, telephone, city, home_number, neighborhood, cep, address, observation, updated_at, created_at
FROM "clients"
WHERE telephone = $1
`

func (q *Queries) FindClient(ctx context.Context, telephone string) (Client, error) {
	row := q.db.QueryRowContext(ctx, findClient, telephone)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Telephone,
		&i.City,
		&i.HomeNumber,
		&i.Neighborhood,
		&i.Cep,
		&i.Address,
		&i.Observation,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateCLient = `-- name: UpdateCLient :exec
UPDATE "clients"
SET "name" = $1
WHERE telephone = $2
`

type UpdateCLientParams struct {
	Name      string
	Telephone string
}

func (q *Queries) UpdateCLient(ctx context.Context, arg UpdateCLientParams) error {
	_, err := q.db.ExecContext(ctx, updateCLient, arg.Name, arg.Telephone)
	return err
}