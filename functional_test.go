package client_form3

import (
	"fmt"
	"github.com/google/uuid"
	Account "github.com/rodrinoblega/client-form3/src/useCases/output"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestFunctionalCreateExistingAccount(t *testing.T) {

	url := os.Getenv("BASE_URL")

	if url == "" {
		url = "localhost:8080"
	}

	fmt.Println(url)
	randomId := uuid.New()
	app := NewService(url, 4*time.Second)

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

	accountCreated, _ := app.Service.Create(account)
	existingAccountCreated, _ := app.Service.Create(account)

	assert.Equal(t, account, accountCreated)
	errorMsg := "{\"error_message\":\"Account cannot be created as it violates a duplicate constraint\"}"
	assert.Equal(t, Account.AccountData{Error: errorMsg}, existingAccountCreated)
}

func TestFunctionalFetch(t *testing.T) {
	url := os.Getenv("BASE_URL")

	if url == "" {
		url = "localhost:8080"
	}

	randomId := uuid.New()
	app := NewService(url, 4*time.Second)

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

	actualCreate, _ := app.Service.Create(account)

	actualFetch, _ := app.Service.Fetch(randomId.String())

	assert.Equal(t, account, actualFetch)
	assert.NotNil(t, actualCreate)
}

func TestFunctionalDeleteNotExistingAccount(t *testing.T) {
	url := os.Getenv("BASE_URL")

	if url == "" {
		url = "localhost:8080"
	}

	app := NewService(url, 4*time.Second)

	actualDelete, _ := app.Service.Delete("rodrigoId", 0)

	assert.False(t, actualDelete)
}

func TestFunctionalDelete(t *testing.T) {
	url := os.Getenv("BASE_URL")

	if url == "" {
		url = "localhost:8080"
	}

	randomId := uuid.New()
	app := NewService(url, 4*time.Second)

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

	actualCreate, _ := app.Service.Create(account)

	actualDelete, _ := app.Service.Delete(randomId.String(), 0)

	actualFetch, _ := app.Service.Fetch(randomId.String())

	assert.NotNil(t, actualCreate)
	assert.True(t, actualDelete)
	errorMsg := "{\"error_message\":\"record " + randomId.String() + " does not exist\"}"
	assert.Equal(t, Account.AccountData{Error: errorMsg}, actualFetch)
}
