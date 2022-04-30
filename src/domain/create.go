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

func (ac *Gateway) Create(account Account.AccountData) (Account.AccountData, error) {
	var err error
	path := PathBuilder.ObtainCreatePath()
	request := RequestBuilder.BuildRequestWithBody(account, ac.Host, path, http.MethodPost)

	response, err := ac.Client.Execute(request)

	if err != nil {
		ErrorHandler.Handle(err)
		return Account.AccountData{Error: err.Error()}, err
	}

	content, err := response.ReadBody()

	if response.StatusCode() != 201 {
		err = errors.New(string(content))
		ErrorHandler.Handle(err)
		return Account.AccountData{Error: err.Error()}, err
	}

	return ResponseInterpreter.Interpreter(content)
}
