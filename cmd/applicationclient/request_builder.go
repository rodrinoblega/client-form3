package applicationclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	Configuration "rnoblega/client-form3/cmd/configuration"
	Instrumentation "rnoblega/client-form3/cmd/configuration"
	Account "rnoblega/client-form3/cmd/dto"
)

func buildRequestWithBody(account Account.AccountData, host string, path string, method string) *http.Request {
	bodyPointer := &Account.Data{Data: account}
	body, _ := json.Marshal(bodyPointer)

	url := Configuration.PROTOCOL + host + "/" + path
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	Instrumentation.TrackError(method + " api call to " + url + " with body:" + string(body))
	request.Header.Add(Configuration.CONTENT_TYPE, Configuration.CONTENT_VALUE)

	return request
}

func buildRequest(host string, path string, method string) *http.Request {
	url := Configuration.PROTOCOL + host + "/" + path
	request, _ := http.NewRequest(method, url, nil)
	Instrumentation.TrackError(method + " api call to " + url)
	request.Header.Add(Configuration.CONTENT_TYPE, Configuration.CONTENT_VALUE)

	return request
}
