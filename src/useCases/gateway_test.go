package useCases

import (
	"errors"
	"github.com/rodrinoblega/client-form3/src/frameworks"
	Account "github.com/rodrinoblega/client-form3/src/useCases/inputoutput"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockClient struct{}

func (mc *MockClient) Execute(req frameworks.Request) (frameworks.Response, error) {
	return frameworks.Response{nil}, errors.New("error mocked")
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
