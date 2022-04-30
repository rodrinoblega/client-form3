package service_client

import (
	"errors"
	"net/http"
	"strconv"
)

func (ac *Gateway) Delete(id string, version int64) (bool, error) {
	var err error
	path := obtainDeletePath(id, version)
	request := buildRequest(ac.Host, path, http.MethodDelete)

	response, err := ac.Client.Execute(request)

	if err != nil {
		handleError(err)
		return false, err
	}

	if response.statusCode() != 204 {
		err = errors.New("Invalid status code: " + string(strconv.FormatInt(int64(response.statusCode()), 10)))
		handleError(err)
		return false, err
	}

	return true, nil
}
