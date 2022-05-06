package useCases

import (
	"bytes"
	"encoding/json"
	"github.com/rodrinoblega/client-form3/src/frameworks"
	"github.com/rodrinoblega/client-form3/src/useCases/output"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockClientFetch struct{}

func (mc *MockClientFetch) Execute(req frameworks.Request) (frameworks.Response, error) {
	account := output.AccountData{
		ID:             "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
		Attributes: &output.AccountAttributes{
			Country:      "AR",
			BaseCurrency: "ARS",
			BankID:       "400301",
			BankIDCode:   "GAL",
			Bic:          "NWBKGB22",
			Name:         []string{"Rodrigo"},
		},
	}

	bodyPointer := &output.Data{Data: account}
	body, _ := json.Marshal(bodyPointer)

	return frameworks.Response{&http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer(body))}}, nil
}

func TestFetchWithoutError(t *testing.T) {
	gateway := Gateway{Client: &MockClientFetch{}}
	account, _ := gateway.Fetch("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")

	assert.Equal(t, "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c", account.ID)
}

func TestFetchStatusCodeNotEqual200(t *testing.T) {
	gateway := Gateway{Client: &MockClientBadRequest{}}
	account, _ := gateway.Fetch("eb0bd6f5-c3f5-44b2-b677-acd23cdde7c")

	assert.Equal(t, "Bad request", account.Error)
}
