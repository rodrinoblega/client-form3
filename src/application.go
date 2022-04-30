package application

import (
	Service "rnoblega/client-form3/src/service_client"
	"time"
)

type Application struct {
	Service Service.GatewayInterface
}

func NewService(host string, time time.Duration) *Application {
	return &Application{
		Service: Service.NewGateway(host, time),
	}
}
