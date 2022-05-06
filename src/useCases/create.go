package useCases

import (
	"errors"
	"github.com/rodrinoblega/client-form3/src/useCases/output"
	"github.com/rodrinoblega/client-form3/src/useCases/pathBuilder"
	"github.com/rodrinoblega/client-form3/src/useCases/requestBuilder"
	"github.com/rodrinoblega/client-form3/src/useCases/responseInterpreter"
	"net/http"
)

func (g *Gateway) Create(account output.AccountData) (output.AccountData, error) {
	var err error
	path := pathBuilder.ObtainCreatePath()
	request := requestBuilder.BuildRequestWithBody(account, g.Host, path, http.MethodPost)

	response, err := g.Client.Execute(request)

	if err != nil {
		trackError(err)
		return output.AccountData{Error: err.Error()}, err
	}

	content, err := response.ReadBody()

	if response.StatusCode() != 201 {
		err = errors.New(string(content))
		trackError(err)
		return output.AccountData{Error: err.Error()}, err
	}

	return responseInterpreter.Interpreter(content)
}
