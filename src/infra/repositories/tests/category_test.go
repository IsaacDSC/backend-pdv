package tests_repository

import (
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/environments"
	"backend-pdv/src/infra/repositories"
	"testing"

	"github.com/savsgio/gotils/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	environments.StartEnv()
}

func TestCreateCategory(t *testing.T) {
	domain := domain.CategoryEntity{
		Name:         "Hamburguer" + uuid.V4()[0:10],
		Description:  uuid.V4()[0:10],
		EnterpriseId: "87332675-50e8-4eea-a35a-e273e676dee4",
	}
	repository_category := repositories.CategoryRepository{}
	err := repository_category.CreateCategory(domain)
	assert.NoError(t, err)
}
