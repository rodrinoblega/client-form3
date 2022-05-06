package useCases

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalError(t *testing.T) {
	accountData, err := Interpreter(nil)

	assert.Equal(t, AccountData{Error: "unexpected end of JSON input"}, accountData)
	assert.Equal(t, "unexpected end of JSON input", err.Error())
}

func TestUnmarshal(t *testing.T) {
	accountEntity := Data{Data: AccountData{ID: "id"}}
	accountBytes, _ := json.Marshal(accountEntity)

	accountData, _ := Interpreter(accountBytes)

	assert.Equal(t, "id", accountData.ID)
}
