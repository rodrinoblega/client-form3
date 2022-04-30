package service_client

import (
	"errors"
	"net/http"
	Account "rnoblega/client-form3/src/dto"
)

func (ac *Gateway) Fetch(id string) (Account.AccountData, error) {
	var err error
	path := obtainFetchPath(id)
	request := buildRequest(ac.Host, path, http.MethodGet)

	response, err := ac.Client.Execute(request)

	if err != nil {
		handleError(err)
		return Account.AccountData{Error: err.Error()}, err
	}

	content, err := response.readBody()

	if response.statusCode() != 200 {
		err = errors.New(string(content))
		handleError(err)
		return Account.AccountData{Error: err.Error()}, err
	}

	return responseInterpreter(content)
}
