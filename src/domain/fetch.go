package domain

import (
	"errors"
	"net/http"
	ErrorHandler "rnoblega/client-form3/src/domain/handler_error"
	PathBuilder "rnoblega/client-form3/src/domain/path_builder"
	RequestBuilder "rnoblega/client-form3/src/domain/request_builder"
	ResponseInterpreter "rnoblega/client-form3/src/domain/response_interpreter"
	Account "rnoblega/client-form3/src/dto"
)

func (ac *Gateway) Fetch(id string) (Account.AccountData, error) {
	var err error
	path := PathBuilder.ObtainFetchPath(id)
	request := RequestBuilder.BuildRequest(ac.Host, path, http.MethodGet)

	response, err := ac.Client.Execute(request)

	if err != nil {
		ErrorHandler.Handle(err)
		return Account.AccountData{Error: err.Error()}, err
	}

	content, err := response.ReadBody()

	if response.StatusCode() != 200 {
		err = errors.New(string(content))
		ErrorHandler.Handle(err)
		return Account.AccountData{Error: err.Error()}, err
	}

	return ResponseInterpreter.Interpreter(content)
}
