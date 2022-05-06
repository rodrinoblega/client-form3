package useCases

import (
	"errors"
	ErrorHandler "github.com/rodrinoblega/client-form3/src/useCases/handler_error"
	Account "github.com/rodrinoblega/client-form3/src/useCases/output"
	PathBuilder "github.com/rodrinoblega/client-form3/src/useCases/path_builder"
	RequestBuilder "github.com/rodrinoblega/client-form3/src/useCases/request_builder"
	ResponseInterpreter "github.com/rodrinoblega/client-form3/src/useCases/response_interpreter"
	"net/http"
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
