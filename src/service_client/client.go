package service_client

import (
	"io/ioutil"
	"net/http"
	"time"
)

type ClientInterface interface {
	Execute(req Request) (Response, error)
}

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

func (r Response) statusCode() int {
	return r.Response.StatusCode
}

func (r Response) readBody() ([]byte, error) {
	return ioutil.ReadAll(r.Response.Body)
}

func (c *Client) Execute(req Request) (Response, error) {
	var err error
	asd := c.HttpClientInterface
	response, _ := asd.Do(req.Request)

	if err != nil {
		handleError(err)
		return Response{}, err
	}

	return Response{response}, err
}

func NewClient(timeout int64) *Client {
	return &Client{
		HttpClientInterface: &http.Client{Timeout: time.Duration(timeout)},
	}
}
