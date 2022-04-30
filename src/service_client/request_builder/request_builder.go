package request_builder

import (
	"bytes"
	"encoding/json"
	"net/http"
	Configuration "rnoblega/client-form3/src/configuration"
	Instrumentation "rnoblega/client-form3/src/configuration"
	Account "rnoblega/client-form3/src/dto"
	"rnoblega/client-form3/src/service_client"
)

func BuildRequestWithBody(account Account.AccountData, host string, path string, method string) service_client.Request {
	bodyPointer := &Account.Data{Data: account}
	body, _ := json.Marshal(bodyPointer)

	url := Configuration.PROTOCOL + host + "/" + path
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	Instrumentation.TrackError(method + " api call to " + url + " with body:" + string(body))
	request.Header.Add(Configuration.CONTENT_TYPE, Configuration.CONTENT_VALUE)

	return service_client.Request{request}
}

func BuildRequest(host string, path string, method string) service_client.Request {
	url := Configuration.PROTOCOL + host + "/" + path
	request, _ := http.NewRequest(method, url, nil)
	Instrumentation.TrackError(method + " api call to " + url)
	request.Header.Add(Configuration.CONTENT_TYPE, Configuration.CONTENT_VALUE)

	return service_client.Request{request}
}
