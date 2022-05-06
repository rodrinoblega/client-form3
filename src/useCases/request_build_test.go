package useCases

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBuildRequest(t *testing.T) {
	request := BuildRequest("localhost:8080", "account", http.MethodGet)

	assert.Equal(t, "localhost:8080", request.Request.URL.Host)
	assert.Equal(t, "/account", request.Request.URL.Path)
}

func TestBuildRequestWithBody(t *testing.T) {
	request := BuildRequestWithBody(AccountData{}, "localhost:8080", "account", http.MethodGet)

	bodyPointer := &Data{Data: AccountData{}}
	body, _ := json.Marshal(bodyPointer)
	assert.Equal(t, "localhost:8080", request.Request.URL.Host)
	assert.Equal(t, "/account", request.Request.URL.Path)
	assert.Equal(t, ioutil.NopCloser(bytes.NewBuffer(body)), request.Request.Body)
}
