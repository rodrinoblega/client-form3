package response_interpreter

import (
	Account "github.com/rodrinoblega/client-form3/src/useCases/output"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalError(t *testing.T) {
	accountData, err := Interpreter(nil)

	assert.Equal(t, Account.AccountData{Error: "unexpected end of JSON input"}, accountData)
	assert.Equal(t, "unexpected end of JSON input", err.Error())
}
