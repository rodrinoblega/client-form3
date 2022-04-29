package service_client

import (
	"errors"
	"io/ioutil"
	"net/http"
	Account "rnoblega/client-form3/src/dto"
)

func (ac *Gateway) Create(account Account.AccountData) (Account.AccountData, error) {
	var err error
	path := obtainCreatePath()
	request := buildRequestWithBody(account, ac.Host, path, http.MethodPost)

	response, err := ac.Client.Do(request)

	if err != nil {
		handleError(err)
		return Account.AccountData{Error: err.Error()}, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 201 {
		err = errors.New(string(content))
		handleError(err)
		return Account.AccountData{Error: err.Error()}, err
	}

	return responseInterpreter(content)
}
