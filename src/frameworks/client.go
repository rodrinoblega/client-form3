package frameworks

import (
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	HttpClientInterface *http.Client
	Request             *Request
	Response            Response
	Timeout             int64
}

type Request struct {
	Request *http.Request
}

type Response struct {
	Response *http.Response
}

func (r Response) StatusCode() int {
	return r.Response.StatusCode
}

func (r Response) ReadBody() ([]byte, error) {
	return ioutil.ReadAll(r.Response.Body)
}

func (c *Client) Execute(req Request) (Response, error) {
	response, err := c.HttpClientInterface.Do(req.Request)
	return Response{response}, err
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		HttpClientInterface: &http.Client{Timeout: timeout},
	}
}
