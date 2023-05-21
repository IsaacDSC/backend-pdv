package tests_end2end

import (
	shared_types "backend-pdv/src/shared/types"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	url = "http://localhost:3333/enterprise"
)

func TestCreateEnterpriseE2E(t *testing.T) {
	payload := strings.NewReader("{\n    \"corporateName\": \"corporateName\",\n    \"fantasyName\": \"enterpriseFake\",\n    \"logo\": \"logo.png\",\n    \"CPJ\": \"1233231321-0001\",\n    \"token\": \"d4198b5e-fac0-4a9a-b123-13d79f727879\"\n}")
	req, _ := http.NewRequest("POST", url, payload)
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
