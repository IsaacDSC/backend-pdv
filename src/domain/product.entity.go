package domain

import (
	"time"

	"github.com/savsgio/gotils/uuid"
)

type ProductEntity struct {
	ID           string
	Name         string
	Description  string
	EnterpriseId string
	CategoryId   string
	Image        string
	Price        float32
	CreatedAt    string
}

func (this_entity *ProductEntity) Validate() (errors []string) {
	this_entity.ID = uuid.V4()
	this_entity.CreatedAt = time.Now().Format(time.RFC3339)
	if len(this_entity.EnterpriseId) == 0 {
		errors = append(errors, "EnterpriseId-Is-Required")
		return
	}
	if len(this_entity.CategoryId) == 0 {
		errors = append(errors, "CategoryId-Is-Required")
		return
	}
	if this_entity.Price < 0 {
		errors = append(errors, "Price-Is-Required-Type-Real")
		return
	}
	return
}
