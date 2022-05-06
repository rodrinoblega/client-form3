package useCases

import (
	"errors"
	"github.com/rodrinoblega/client-form3/src/useCases/pathBuilder"
	"github.com/rodrinoblega/client-form3/src/useCases/requestBuilder"
	"net/http"

	"strconv"
)

func (g *Gateway) Delete(id string, version int64) (bool, error) {
	var err error
	path := pathBuilder.ObtainDeletePath(id, version)
	request := requestBuilder.BuildRequest(g.Host, path, http.MethodDelete)

	response, err := g.Client.Execute(request)

	if err != nil {
		trackError(err)
		return false, err
	}

	if response.StatusCode() != 204 {
		err = errors.New("Invalid status code: " + string(strconv.FormatInt(int64(response.StatusCode()), 10)))
		trackError(err)
		return false, err
	}

	return true, nil
}
