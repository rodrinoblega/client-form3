package client_form3

import (
	Gateway "github.com/rodrinoblega/client-form3/src/adapters/gateway"
	Service "github.com/rodrinoblega/client-form3/src/adapters/service"
	"time"
)

type Application struct {
	Service Gateway.GatewayInterface
}

func NewService(host string, time time.Duration) *Application {
	return &Application{
		Service: Service.CreateService(host, time),
	}
}
