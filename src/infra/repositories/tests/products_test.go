package tests_repository

import (
	"backend-pdv/src/domain"
	"backend-pdv/src/infra/environments"
	"backend-pdv/src/infra/repositories"
	"testing"

	"github.com/savsgio/gotils/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	enterpriseId = "87332675-50e8-4eea-a35a-e273e676dee4"
	categoryId   = "7627745f-2c6d-405e-a1b6-64e5273359a3"
)

func init() {
	environments.StartEnv()
}

func TestCreateProducts(t *testing.T) {
	repository_product := repositories.ProductsRepository{}
	err := repository_product.CreateProduct(domain.ProductEntity{
		CategoryId:   categoryId,
		EnterpriseId: enterpriseId,
		Name:         uuid.V4(),
		Description:  uuid.V4(),
		Image:        uuid.V4() + ".png",
		Price:        57.99,
	})
	assert.NoError(t, err)
}
