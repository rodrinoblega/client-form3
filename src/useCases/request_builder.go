package useCases

import (
	"bytes"
	"encoding/json"
	Configuration "github.com/rodrinoblega/client-form3/src/configuration"
	Client "github.com/rodrinoblega/client-form3/src/frameworks"
	"log"
	"net/http"
)

func BuildRequestWithBody(account AccountData, host string, path string, method string) Client.Request {
	bodyPointer := &Data{Data: account}
	body, _ := json.Marshal(bodyPointer)

	url := Configuration.PROTOCOL + host + "/" + path
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	log.Println(method + " api call to " + url + " with body:" + string(body))
	request.Header.Add(Configuration.CONTENT_TYPE, Configuration.CONTENT_VALUE)

	return Client.Request{request}
}

func BuildRequest(host string, path string, method string) Client.Request {
	url := Configuration.PROTOCOL + host + "/" + path
	request, _ := http.NewRequest(method, url, nil)
	log.Println(method + " api call to " + url)
	request.Header.Add(Configuration.CONTENT_TYPE, Configuration.CONTENT_VALUE)

	return Client.Request{request}
}
