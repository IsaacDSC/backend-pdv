package domain

import (
	"time"

	"github.com/savsgio/gotils/uuid"
)

type CategoryEntity struct {
	ID           string
	Name         string
	Description  string
	EnterpriseId string
	CreatedAt    string
}

func (this_entity *CategoryEntity) Validate() (errors []string) {
	this_entity.ID = uuid.V4()
	this_entity.CreatedAt = time.Now().Format(time.RFC3339)
	if len(this_entity.EnterpriseId) == 0 {
		errors = append(errors, "EnterpriseId-Is-Required")
	}
	return
}
