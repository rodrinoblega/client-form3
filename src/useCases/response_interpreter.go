package useCases

import (
	"encoding/json"
	"github.com/rodrinoblega/client-form3/src/configuration"
)

func Interpreter(content []byte) (AccountData, error) {
	var err error
	realResponse := &Data{}

	err = json.Unmarshal(content, realResponse)

	if err != nil {
		configuration.TrackError(err.Error())
		return AccountData{Error: err.Error()}, err
	}

	return realResponse.Data, nil
}
