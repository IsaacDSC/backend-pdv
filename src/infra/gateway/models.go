// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package gateway

import (
	"database/sql"
)

type Category struct {
	ID           string
	Name         string
	Description  sql.NullString
	EnterpriseID sql.NullString
	UpdatedAt    sql.NullTime
	CreatedAt    sql.NullTime
}

type Client struct {
	ID           string
	Name         string
	Email        sql.NullString
	Password     sql.NullString
	Telephone    string
	City         sql.NullString
	HomeNumber   sql.NullString
	Neighborhood sql.NullString
	Cep          sql.NullString
	Address      sql.NullString
	Observation  sql.NullString
	EnterpriseID string
	UpdatedAt    sql.NullTime
	CreatedAt    sql.NullTime
}

type Enterprise struct {
	ID            string
	CorporateName sql.NullString
	FantasyName   sql.NullString
	Logo          sql.NullString
	Cnpj          sql.NullString
	Status        sql.NullString
	Token         sql.NullString
	UpdatedAt     sql.NullTime
	CreatedAt     sql.NullTime
}

type Order struct {
	ID            string
	TotalPrice    sql.NullFloat64
	PriceDelivery sql.NullFloat64
	PriceItems    sql.NullFloat64
	AddressID     sql.NullString
	ClientID      sql.NullString
	Status        sql.NullString
	UpdatedAt     sql.NullTime
	CreatedAt     sql.NullTime
}

type OrderItem struct {
	ID                 string
	ProductID          sql.NullString
	ProductName        sql.NullString
	ProductDescription sql.NullString
	ProductImages      sql.NullString
	CategoryName       sql.NullString
	PriceProduct       sql.NullFloat64
	UpdatedAt          sql.NullTime
	CreatedAt          sql.NullTime
	OrderID            string
}

type Product struct {
	ID           string
	CategoryID   string
	EnterpriseID string
	Name         string
	Description  sql.NullString
	Image        sql.NullString
	Price        sql.NullFloat64
	UpdatedAt    sql.NullTime
	CreatedAt    sql.NullTime
}
