package application

import (
	Service "rnoblega/client-form3/src/service_client"
)

type Application struct {
	Service Service.GatewayInterface
}

func NewService(host string, timeout int64) *Application {
	return &Application{
		Service: Service.NewGateway(host, timeout),
	}
}
