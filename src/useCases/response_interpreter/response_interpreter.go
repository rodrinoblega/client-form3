package response_interpreter

import (
	"encoding/json"
	Instrumentation "github.com/rodrinoblega/client-form3/src/configuration"
	Account "github.com/rodrinoblega/client-form3/src/useCases/output"
)

func Interpreter(content []byte) (Account.AccountData, error) {
	var err error
	realResponse := &Account.Data{}

	err = json.Unmarshal(content, realResponse)

	if err != nil {
		Instrumentation.TrackError(err.Error())
		return Account.AccountData{Error: err.Error()}, err
	}

	return realResponse.Data, nil
}
