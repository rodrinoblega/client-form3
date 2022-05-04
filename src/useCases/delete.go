package useCases

import (
	"errors"
	"net/http"
	ErrorHandler "rnoblega/client-form3/src/useCases/handler_error"
	PathBuilder "rnoblega/client-form3/src/useCases/path_builder"
	RequestBuilder "rnoblega/client-form3/src/useCases/request_builder"

	"strconv"
)

func (g *Gateway) Delete(id string, version int64) (bool, error) {
	var err error
	path := PathBuilder.ObtainDeletePath(id, version)
	request := RequestBuilder.BuildRequest(g.Host, path, http.MethodDelete)

	response, err := g.Client.Execute(request)

	if err != nil {
		ErrorHandler.Handle(err)
		return false, err
	}

	if response.StatusCode() != 204 {
		err = errors.New("Invalid status code: " + string(strconv.FormatInt(int64(response.StatusCode()), 10)))
		ErrorHandler.Handle(err)
		return false, err
	}

	return true, nil
}
