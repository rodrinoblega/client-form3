package applicationclient

import (
	"errors"
	"io/ioutil"
	"net/http"
	Account "rnoblega/client-form3/cmd/dto"
)

func (ac *ApplicationClient) Fetch(id string) (Account.AccountData, error) {
	var err error
	path := obtainFetchPath(id)
	request := buildRequest(ac.Host, path, http.MethodGet)

	response, err := ac.Client.Do(request)

	if err != nil {
		handleError(err)
		return Account.AccountData{Error: err.Error()}, err
	}

	content, err := ioutil.ReadAll(response.Body)

	if response.StatusCode != 200 {
		err = errors.New(string(content))
		handleError(err)
		return Account.AccountData{Error: err.Error()}, err
	}

	return responseInterpreter(content)
}
