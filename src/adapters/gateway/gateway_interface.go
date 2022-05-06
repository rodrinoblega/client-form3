package gateway

import (
	"github.com/rodrinoblega/client-form3/src/frameworks"
	UseCases "github.com/rodrinoblega/client-form3/src/useCases"
	"time"
)

type GatewayInterface interface {
	Create(account UseCases.AccountData) (UseCases.AccountData, error)
	Fetch(id string) (UseCases.AccountData, error)
	Delete(id string, version int64) (bool, error)
}

func NewGateway(host string, timeout time.Duration) *UseCases.Gateway {
	return &UseCases.Gateway{
		Host:   host,
		Client: frameworks.NewClient(timeout),
	}
}
