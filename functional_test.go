package main

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	Application "rnoblega/client-form3/src"
	Account "rnoblega/client-form3/src/dto"
	"testing"
	"time"
)

func Test_functional_Create(t *testing.T) {

	randomId := uuid.New()
	appClient := Application.New("localhost:8080", 5*time.Second)

	account := Account.AccountData{
		ID:             randomId.String(),
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

	actualCreate, _ := appClient.ApplicationClient.Create(account)

	assert.Equal(t, account, actualCreate)
}

func Test_functional_Fetch(t *testing.T) {

	randomId := uuid.New()
	appClient := Application.New("localhost:8080", 5*time.Second)

	account := Account.AccountData{
		ID:             randomId.String(),
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

	actualCreate, _ := appClient.ApplicationClient.Create(account)

	actualFetch, _ := appClient.ApplicationClient.Fetch(randomId.String())

	assert.Equal(t, account, actualFetch)
	assert.NotNil(t, actualCreate)
}

func Test_functional_Delete(t *testing.T) {

	randomId := uuid.New()
	appClient := Application.New("localhost:8080", 5*time.Second)

	account := Account.AccountData{
		ID:             randomId.String(),
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

	actualCreate, _ := appClient.ApplicationClient.Create(account)

	actualDelete, _ := appClient.ApplicationClient.Delete(randomId.String(), 0)

	actualFetch, _ := appClient.ApplicationClient.Fetch(randomId.String())

	assert.NotNil(t, actualCreate)
	assert.True(t, actualDelete)
	errorMsg := "{\"error_message\":\"record " + randomId.String() + " does not exist\"}"
	assert.Equal(t, Account.AccountData{Error: errorMsg}, actualFetch)
}
