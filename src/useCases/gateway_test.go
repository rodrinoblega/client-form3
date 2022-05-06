package useCases

import (
	"bytes"
	"errors"
	"github.com/rodrinoblega/client-form3/src/frameworks"
	"github.com/rodrinoblega/client-form3/src/useCases/output"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockClientWithError struct{}

type MockClientBadRequest struct{}

func (mc *MockClientWithError) Execute(req frameworks.Request) (frameworks.Response, error) {
	return frameworks.Response{nil}, errors.New("error mocked")
}

func (mc *MockClientBadRequest) Execute(req frameworks.Request) (frameworks.Response, error) {
	body := "Bad request"
	return frameworks.Response{&http.Response{StatusCode: 400, Body: ioutil.NopCloser(bytes.NewBufferString(body))}}, nil
}

func TestFetchExecuteMethodWith(t *testing.T) {
	gateway := Gateway{Client: &MockClientWithError{}}
	account, err := gateway.Fetch("fakeId")

	assert.Equal(t, errors.New("error mocked"), err)
	assert.Equal(t, output.AccountData{Error: "error mocked"}, account)
}

func TestDeleteExecuteMethodWith(t *testing.T) {
	gateway := Gateway{Client: &MockClientWithError{}}
	bool, err := gateway.Delete("fakeId", 0)

	assert.Equal(t, errors.New("error mocked"), err)
	assert.Equal(t, false, bool)
}

func TestCreateExecuteMethodWith(t *testing.T) {
	gateway := Gateway{Client: &MockClientWithError{}}
	account, err := gateway.Create(output.AccountData{})

	assert.Equal(t, errors.New("error mocked"), err)
	assert.Equal(t, output.AccountData{Error: "error mocked"}, account)
}
