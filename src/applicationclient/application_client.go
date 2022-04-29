package applicationclient

import (
	"net/http"
	Account "rnoblega/client-form3/src/dto"
	"time"
)

type ApplicationClient struct {
	Host   string
	Client *http.Client
}

type ApplicationClientInterface interface {
	Create(account Account.AccountData) (Account.AccountData, error)
	Fetch(id string) (Account.AccountData, error)
	Delete(id string, version int64) (bool, error)
}

func NewAppClient(host string, timeout time.Duration) *ApplicationClient {
	return &ApplicationClient{
		Host:   host,
		Client: &http.Client{Timeout: timeout},
	}
}
