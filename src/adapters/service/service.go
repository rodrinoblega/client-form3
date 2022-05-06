package service

import (
	Service "github.com/rodrinoblega/client-form3/src/adapters/gateway"
	UseCases "github.com/rodrinoblega/client-form3/src/useCases"
	"time"
)

func CreateService(host string, time time.Duration) *UseCases.Gateway {
	return Service.NewGateway(host, time)
}
