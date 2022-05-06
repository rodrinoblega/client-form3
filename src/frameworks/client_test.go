package frameworks

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type Data struct {
	Data string
}

func TestStatusCode(t *testing.T) {
	bodyPointer := &Data{Data: "account"}
	body, _ := json.Marshal(bodyPointer)
	response := Response{Response: &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer(body))}}
	statusCode := response.StatusCode()

	assert.Equal(t, 200, statusCode)
}

func TestReadBody(t *testing.T) {
	bodyPointer := &Data{Data: "account"}
	body, _ := json.Marshal(bodyPointer)
	response := Response{Response: &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer(body))}}

	content, _ := response.ReadBody()
	realResponse := &Data{}

	_ = json.Unmarshal(content, realResponse)
	assert.Equal(t, &Data{Data: "account"}, realResponse)
}

func TestExecuteWithBadRequestUri(t *testing.T) {
	client := NewClient(1 * time.Second)

	response, _ := client.Execute(Request{&http.Request{RequestURI: "local"}})

	assert.Equal(t, Response{nil}, response)
}
