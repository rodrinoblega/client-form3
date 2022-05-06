package request_builder

import (
	"bytes"
	"encoding/json"
	Configuration "github.com/rodrinoblega/client-form3/src/configuration"
	Instrumentation "github.com/rodrinoblega/client-form3/src/configuration"
	Client "github.com/rodrinoblega/client-form3/src/frameworks"
	Account "github.com/rodrinoblega/client-form3/src/useCases/inputoutput"
	"net/http"
)

func BuildRequestWithBody(account Account.AccountData, host string, path string, method string) Client.Request {
	bodyPointer := &Account.Data{Data: account}
	body, _ := json.Marshal(bodyPointer)

	url := Configuration.PROTOCOL + host + "/" + path
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	Instrumentation.TrackError(method + " api call to " + url + " with body:" + string(body))
	request.Header.Add(Configuration.CONTENT_TYPE, Configuration.CONTENT_VALUE)

	return Client.Request{request}
}

func BuildRequest(host string, path string, method string) Client.Request {
	url := Configuration.PROTOCOL + host + "/" + path
	request, _ := http.NewRequest(method, url, nil)
	Instrumentation.TrackError(method + " api call to " + url)
	request.Header.Add(Configuration.CONTENT_TYPE, Configuration.CONTENT_VALUE)

	return Client.Request{request}
}
