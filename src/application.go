package application

import (
	Service "rnoblega/client-form3/src/service_client"
	"time"
)

type Application struct {
	Service Service.GatewayInterface
}

func NewService(host string, timeout time.Duration) *Application {
	return &Application{
		Service: Service.NewGateway(host, timeout),
	}
}
