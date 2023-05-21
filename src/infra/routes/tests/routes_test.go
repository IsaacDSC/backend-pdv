package tests_end2end

import (
	shared_types "backend-pdv/src/shared/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/savsgio/gotils/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	url = "http://localhost:3333"
)

func TestCreateEnterpriseE2E(t *testing.T) {
	path := "/enterprise"
	payload := strings.NewReader("{\n    \"corporateName\": \"corporateName\",\n    \"fantasyName\": \"enterpriseFake\",\n    \"logo\": \"logo.png\",\n    \"CPJ\": \"1233231321-0001\",\n    \"token\": \"d4198b5e-fac0-4a9a-b123-13d79f727879\"\n}")
	URL := fmt.Sprintf("%s%s", url, path)
	req, _ := http.NewRequest("POST", URL, payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var response shared_types.OutputCreateEnterprise
	json.Unmarshal(body, &response)
	assert.Equal(t, 201, res.StatusCode)
	assert.NotEmpty(t, response.EnterpriseId)
	assert.NotEmpty(t, response.Token)
}

func TestCreateCategoryE2E(t *testing.T) {
	input := shared_types.CreateCategory_DTO{
		Name:         "Category" + uuid.V4(),
		Description:  uuid.V4(),
		EnterpriseId: "8861d131-7274-466b-851e-414677e203e6",
	}
	var payload bytes.Buffer
	err := json.NewEncoder(&payload).Encode(input)
	assert.NoError(t, err)
	path := "/category"
	URL := fmt.Sprintf("%s%s", url, path)
	req, _ := http.NewRequest("POST", URL, &payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, 201, res.StatusCode)
}

func TestCreateProductE2E(t *testing.T) {
	input := shared_types.CreateProduct_DTO{
		Name:         "ProductNamed" + uuid.V4(),
		Description:  uuid.V4(),
		EnterpriseId: "8861d131-7274-466b-851e-414677e203e6",
		CategoryId:   "82d71c4a-838b-4d45-952d-36559fd1ba89",
		Image:        uuid.V4() + ".png",
		Price:        99.99,
	}
	var payload bytes.Buffer
	err := json.NewEncoder(&payload).Encode(input)
	assert.NoError(t, err)
	path := "/product"
	URL := fmt.Sprintf("%s%s", url, path)
	req, _ := http.NewRequest("POST", URL, &payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, 201, res.StatusCode)
}

func TestCreateClientE2E(t *testing.T) {
	input := shared_types.ClientCreateOnSet_DTO{
		Name:         "ClientName",
		EnterpriseId: "87332675-50e8-4eea-a35a-e273e676dee4",
		Telephone:    uuid.V4()[0:11],
		Email:        uuid.V4()[0:10],
		Address: shared_types.AddressDTO{
			City:         uuid.V4()[0:10],
			HomeNumber:   "200",
			Neighborhood: uuid.V4()[0:10],
			Cep:          uuid.V4()[0:10],
			Address:      uuid.V4()[0:10],
			Observation:  uuid.V4()[0:10],
		},
	}
	var payload bytes.Buffer
	err := json.NewEncoder(&payload).Encode(input)
	assert.NoError(t, err)
	path := "/client"
	URL := fmt.Sprintf("%s%s", url, path)
	req, _ := http.NewRequest("POST", URL, &payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, 201, res.StatusCode)
}
