package useCases

import (
	"github.com/rodrinoblega/client-form3/src/frameworks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type MockClientDelete struct{}

func (mc *MockClientDelete) Execute(req frameworks.Request) (frameworks.Response, error) {
	return frameworks.Response{&http.Response{StatusCode: 204}}, nil
}

func TestDeleteWithoutError(t *testing.T) {
	gateway := Gateway{Client: &MockClientDelete{}}
	deleted, _ := gateway.Delete("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c", 0)

	assert.Equal(t, true, deleted)
}

func TestDeleteStatusCodeNotEqual204(t *testing.T) {
	gateway := Gateway{Client: &MockClientBadRequest{}}
	deleted, _ := gateway.Delete("eb0bd6f5-c3f5-44b2-b677-acd23cdde7c", 0)

	assert.Equal(t, false, deleted)
}
