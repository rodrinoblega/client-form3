package service_client

import (
	"errors"
	"net/http"
	Account "rnoblega/client-form3/src/dto"
	ErrorHandler "rnoblega/client-form3/src/service_client/handler_error"
	PathBuilder "rnoblega/client-form3/src/service_client/path_builder"
	RequestBuilder "rnoblega/client-form3/src/service_client/request_builder"
	ResponseInterpreter "rnoblega/client-form3/src/service_client/response_interpreter"
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

	content, err := response.readBody()

	if response.statusCode() != 200 {
		err = errors.New(string(content))
		ErrorHandler.Handle(err)
		return Account.AccountData{Error: err.Error()}, err
	}

	return ResponseInterpreter.Interpreter(content)
}
