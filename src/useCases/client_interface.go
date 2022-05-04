package useCases

import Client "rnoblega/client-form3/src/frameworks"

type ClientInterface interface {
	Execute(req Client.Request) (Client.Response, error)
}
