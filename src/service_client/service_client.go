package service_client

import (
	"net/http"
	Account "rnoblega/client-form3/src/dto"
	"time"
)

type Client struct {
	Host   string
	Client *http.Client
}

type ClientInterface interface {
	Create(account Account.AccountData) (Account.AccountData, error)
	Fetch(id string) (Account.AccountData, error)
	Delete(id string, version int64) (bool, error)
}

func NewClient(host string, timeout time.Duration) *Client {
	return &Client{
		Host:   host,
		Client: &http.Client{Timeout: timeout},
	}
}
