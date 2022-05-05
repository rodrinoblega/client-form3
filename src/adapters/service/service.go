package service

import (
	Service "rnoblega/client-form3/src/adapters/gateway"
	UseCases "rnoblega/client-form3/src/useCases"
	"time"
)

func CreateService(host string, time time.Duration) *UseCases.Gateway {
	return Service.NewGateway(host, time)
}
