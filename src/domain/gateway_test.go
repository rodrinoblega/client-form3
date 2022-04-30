package domain

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"rnoblega/client-form3/src/domain/client"
	Account "rnoblega/client-form3/src/dto"
	"testing"
)

type MockClient struct{}

func (mc *MockClient) Execute(req client.Request) (client.Response, error) {
	return client.Response{nil}, errors.New("error mocked")
}

func TestFetchExecuteMethodWith(t *testing.T) {
	gateway := Gateway{Client: &MockClient{}}
	account, err := gateway.Fetch("fakeId")

	assert.Equal(t, errors.New("error mocked"), err)
	assert.Equal(t, Account.AccountData{Error: "error mocked"}, account)
}

func TestDeleteExecuteMethodWith(t *testing.T) {
	gateway := Gateway{Client: &MockClient{}}
	bool, err := gateway.Delete("fakeId", 0)

	assert.Equal(t, errors.New("error mocked"), err)
	assert.Equal(t, false, bool)
}

func TestCreateExecuteMethodWith(t *testing.T) {
	gateway := Gateway{Client: &MockClient{}}
	account, err := gateway.Create(Account.AccountData{})

	assert.Equal(t, errors.New("error mocked"), err)
	assert.Equal(t, Account.AccountData{Error: "error mocked"}, account)
}
