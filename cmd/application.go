package application

import (
	ApplicationClient "rnoblega/client-form3/cmd/applicationclient"
	"time"
)

type Application struct {
	ApplicationClient ApplicationClient.ApplicationClientInterface
}

func New(host string, timeout time.Duration) *Application {
	return &Application{
		ApplicationClient: ApplicationClient.NewAppClient(host, timeout),
	}
}
