package domain

import (
	"errors"
	"net/http"
	ErrorHandler "rnoblega/client-form3/src/domain/handler_error"
	PathBuilder "rnoblega/client-form3/src/domain/path_builder"
	RequestBuilder "rnoblega/client-form3/src/domain/request_builder"

	"strconv"
)

func (ac *Gateway) Delete(id string, version int64) (bool, error) {
	var err error
	path := PathBuilder.ObtainDeletePath(id, version)
	request := RequestBuilder.BuildRequest(ac.Host, path, http.MethodDelete)

	response, err := ac.Client.Execute(request)

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