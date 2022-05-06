package useCases

import (
	"errors"
	ErrorHandler "github.com/rodrinoblega/client-form3/src/useCases/handler_error"
	PathBuilder "github.com/rodrinoblega/client-form3/src/useCases/path_builder"
	RequestBuilder "github.com/rodrinoblega/client-form3/src/useCases/request_builder"
	"net/http"

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
