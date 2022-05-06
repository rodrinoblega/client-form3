package useCases

import (
	"errors"
	"github.com/rodrinoblega/client-form3/src/useCases/output"
	"github.com/rodrinoblega/client-form3/src/useCases/pathBuilder"
	"github.com/rodrinoblega/client-form3/src/useCases/requestBuilder"
	"github.com/rodrinoblega/client-form3/src/useCases/responseInterpreter"
	"net/http"
)

func (g *Gateway) Fetch(id string) (output.AccountData, error) {
	var err error
	path := pathBuilder.ObtainFetchPath(id)
	request := requestBuilder.BuildRequest(g.Host, path, http.MethodGet)

	response, err := g.Client.Execute(request)

	if err != nil {
		trackError(err)
		return output.AccountData{Error: err.Error()}, err
	}

	content, err := response.ReadBody()

	if response.StatusCode() != 200 {
		err = errors.New(string(content))
		trackError(err)
		return output.AccountData{Error: err.Error()}, err
	}

	return responseInterpreter.Interpreter(content)
}
