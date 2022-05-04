package gateway

import (
	Account "rnoblega/client-form3/src/entities"
	"rnoblega/client-form3/src/frameworks"
	UseCases "rnoblega/client-form3/src/useCases"
	"time"
)

type GatewayInterface interface {
	Create(account Account.AccountData) (Account.AccountData, error)
	Fetch(id string) (Account.AccountData, error)
	Delete(id string, version int64) (bool, error)
}

func NewGateway(host string, timeout time.Duration) *UseCases.Gateway {
	return &UseCases.Gateway{
		Host:   host,
		Client: frameworks.NewClient(timeout),
	}
}
