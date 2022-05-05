package application

import (
	Gateway "rnoblega/client-form3/src/adapters/gateway"
	Service "rnoblega/client-form3/src/adapters/service"
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
