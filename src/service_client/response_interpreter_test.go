package service_client

import (
	"github.com/stretchr/testify/assert"
	Account "rnoblega/client-form3/src/dto"
	"testing"
)

func TestUnmarshalError(t *testing.T) {
	accountData, err := responseInterpreter(nil)

	assert.Equal(t, Account.AccountData{Error: "unexpected end of JSON input"}, accountData)
	assert.Equal(t, "unexpected end of JSON input", err.Error())
}
