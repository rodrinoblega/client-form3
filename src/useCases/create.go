package useCases

import (
	"errors"
	"net/http"
)

func (g *Gateway) Create(account AccountData) (AccountData, error) {
	var err error
	path := ObtainCreatePath()
	request := BuildRequestWithBody(account, g.Host, path, http.MethodPost)

	response, err := g.Client.Execute(request)

	if err != nil {
		handle(err)
		return AccountData{Error: err.Error()}, err
	}

	content, err := response.ReadBody()

	if response.StatusCode() != 201 {
		err = errors.New(string(content))
		handle(err)
		return AccountData{Error: err.Error()}, err
	}

	return Interpreter(content)
}
