package domain

import (
	"rnoblega/client-form3/src/domain/client"
	Account "rnoblega/client-form3/src/dto"
	"time"
)

type Gateway struct {
	Host   string
	Client client.ClientInterface
}

type GatewayInterface interface {
	Create(account Account.AccountData) (Account.AccountData, error)
	Fetch(id string) (Account.AccountData, error)
	Delete(id string, version int64) (bool, error)
}

func NewGateway(host string, timeout time.Duration) *Gateway {
	return &Gateway{
		Host:   host,
		Client: client.NewClient(timeout),
	}
}
