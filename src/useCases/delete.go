package useCases

import (
	"errors"
	"net/http"

	"strconv"
)

func (g *Gateway) Delete(id string, version int64) (bool, error) {
	var err error
	path := ObtainDeletePath(id, version)
	request := BuildRequest(g.Host, path, http.MethodDelete)

	response, err := g.Client.Execute(request)

	if err != nil {
		Handle(err)
		return false, err
	}

	if response.StatusCode() != 204 {
		err = errors.New("Invalid status code: " + string(strconv.FormatInt(int64(response.StatusCode()), 10)))
		Handle(err)
		return false, err
	}

	return true, nil
}
