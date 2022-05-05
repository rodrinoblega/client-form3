package useCases

import (
	"errors"
	"net/http"
	ErrorHandler "rnoblega/client-form3/src/useCases/handler_error"
	Account "rnoblega/client-form3/src/useCases/inputoutput"
	PathBuilder "rnoblega/client-form3/src/useCases/path_builder"
	RequestBuilder "rnoblega/client-form3/src/useCases/request_builder"
	ResponseInterpreter "rnoblega/client-form3/src/useCases/response_interpreter"
)

func (g *Gateway) Create(account Account.AccountData) (Account.AccountData, error) {
	var err error
	path := PathBuilder.ObtainCreatePath()
	request := RequestBuilder.BuildRequestWithBody(account, g.Host, path, http.MethodPost)

	response, err := g.Client.Execute(request)

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
