package applicationclient

import (
	"encoding/json"
	Instrumentation "rnoblega/client-form3/src/configuration"
	Account "rnoblega/client-form3/src/dto"
)

func responseInterpreter(content []byte) (Account.AccountData, error) {
	var err error
	realResponse := &Account.Data{}

	err = json.Unmarshal(content, realResponse)

	if err != nil {
		Instrumentation.TrackError(err.Error())
		return Account.AccountData{}, err
	}

	return realResponse.Data, nil
}
