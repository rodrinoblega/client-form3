package useCases

import (
	"bytes"
	"encoding/json"
	"github.com/rodrinoblega/client-form3/src/frameworks"
	Account "github.com/rodrinoblega/client-form3/src/useCases/output"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockClientCreate struct{}

func (mc *MockClientCreate) Execute(req frameworks.Request) (frameworks.Response, error) {
	account := Account.AccountData{
		ID:             "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
		Attributes: &Account.AccountAttributes{
			Country:      "AR",
			BaseCurrency: "ARS",
			BankID:       "400301",
			BankIDCode:   "GAL",
			Bic:          "NWBKGB22",
			Name:         []string{"Rodrigo"},
		},
	}

	bodyPointer := &Account.Data{Data: account}
	body, _ := json.Marshal(bodyPointer)

	return frameworks.Response{&http.Response{StatusCode: 201, Body: ioutil.NopCloser(bytes.NewBuffer(body))}}, nil
}

func TestCreateWithoutError(t *testing.T) {
	account := Account.AccountData{
		ID:             "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
		Attributes: &Account.AccountAttributes{
			Country:      "AR",
			BaseCurrency: "ARS",
			BankID:       "400301",
			BankIDCode:   "GAL",
			Bic:          "NWBKGB22",
			Name:         []string{"Rodrigo"},
		},
	}

	gateway := Gateway{Client: &MockClientCreate{}}
	accountCreated, _ := gateway.Create(account)

	assert.Equal(t, account, accountCreated)
}

func TestFetchStatusCodeNotEqual201(t *testing.T) {
	account := Account.AccountData{
		ID:             "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
		Attributes: &Account.AccountAttributes{
			Country:      "AR",
			BaseCurrency: "ARS",
			BankID:       "400301",
			BankIDCode:   "GAL",
			Bic:          "NWBKGB22",
			Name:         []string{"Rodrigo"},
		},
	}

	gateway := Gateway{Client: &MockClientBadRequest{}}
	accountCreated, _ := gateway.Create(account)

	assert.Equal(t, "Bad request", accountCreated.Error)
}
