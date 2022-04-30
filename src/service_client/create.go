package service_client

import (
	"errors"
	"net/http"
	Account "rnoblega/client-form3/src/dto"
)

func (ac *Gateway) Create(account Account.AccountData) (Account.AccountData, error) {
	var err error
	path := obtainCreatePath()
	request := buildRequestWithBody(account, ac.Host, path, http.MethodPost)

	response, err := ac.Client.Execute(request)

	if err != nil {
		handleError(err)
		return Account.AccountData{Error: err.Error()}, err
	}

	content, err := response.readBody()

	if response.statusCode() != 201 {
		err = errors.New(string(content))
		handleError(err)
		return Account.AccountData{Error: err.Error()}, err
	}

	return responseInterpreter(content)
}
