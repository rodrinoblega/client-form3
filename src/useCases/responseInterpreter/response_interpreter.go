package responseInterpreter

import (
	"encoding/json"
	"github.com/rodrinoblega/client-form3/src/useCases"
	"github.com/rodrinoblega/client-form3/src/useCases/output"
)

func Interpreter(content []byte) (output.AccountData, error) {
	var err error
	realResponse := &output.Data{}

	err = json.Unmarshal(content, realResponse)

	if err != nil {
		useCases.trackError(err)
		return output.AccountData{Error: err.Error()}, err
	}

	return realResponse.Data, nil
}
