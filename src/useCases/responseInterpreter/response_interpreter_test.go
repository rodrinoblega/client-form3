package responseInterpreter

import (
	"encoding/json"
	"github.com/rodrinoblega/client-form3/src/useCases/output"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalError(t *testing.T) {
	accountData, err := Interpreter(nil)

	assert.Equal(t, output.AccountData{Error: "unexpected end of JSON input"}, accountData)
	assert.Equal(t, "unexpected end of JSON input", err.Error())
}

func TestUnmarshal(t *testing.T) {
	accountEntity := output.Data{Data: output.AccountData{ID: "id"}}
	accountBytes, _ := json.Marshal(accountEntity)

	accountData, _ := Interpreter(accountBytes)

	assert.Equal(t, "id", accountData.ID)
}
