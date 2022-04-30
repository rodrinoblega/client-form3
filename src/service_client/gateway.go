package service_client

import (
	Account "rnoblega/client-form3/src/dto"
)

type Gateway struct {
	Host   string
	Client ClientInterface
}

type GatewayInterface interface {
	Create(account Account.AccountData) (Account.AccountData, error)
	Fetch(id string) (Account.AccountData, error)
	Delete(id string, version int64) (bool, error)
}

func NewGateway(host string, timeout int64) *Gateway {
	return &Gateway{
		Host:   host,
		Client: NewClient(timeout),
	}
}
