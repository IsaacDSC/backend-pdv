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

func TestCreateEnterprise(t *testing.T) {
	entity := domain.EnterpriseEntity{
		CorporateName: "corporate_name",
		FantasyName:   "FantasyName",
		Logo:          "image.png",
		CPJ:           "123321232313-12312",
		Token:         uuid.V4(),
	}
	repository_enterprise := repositories.EnterpriseRepository{}
	err := repository_enterprise.CreateEnterprise(entity)
	assert.NoError(t, err)
}
