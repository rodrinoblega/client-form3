package useCases

import (
	"errors"
	"net/http"
)

func (g *Gateway) Fetch(id string) (AccountData, error) {
	var err error
	path := ObtainFetchPath(id)
	request := BuildRequest(g.Host, path, http.MethodGet)

	response, err := g.Client.Execute(request)

	if err != nil {
		Handle(err)
		return AccountData{Error: err.Error()}, err
	}

	content, err := response.ReadBody()

	if response.StatusCode() != 200 {
		err = errors.New(string(content))
		Handle(err)
		return AccountData{Error: err.Error()}, err
	}

	return Interpreter(content)
}
