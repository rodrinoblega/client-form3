package useCases

import Client "github.com/rodrinoblega/client-form3/src/frameworks"

type ClientInterface interface {
	Execute(req Client.Request) (Client.Response, error)
}
