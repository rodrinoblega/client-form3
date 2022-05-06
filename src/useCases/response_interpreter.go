package useCases

import (
	"encoding/json"
)

func Interpreter(content []byte) (AccountData, error) {
	var err error
	realResponse := &Data{}

	err = json.Unmarshal(content, realResponse)

	if err != nil {
		TrackError(err.Error())
		return AccountData{Error: err.Error()}, err
	}

	return realResponse.Data, nil
}
