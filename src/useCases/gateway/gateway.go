package gateway

import "github.com/rodrinoblega/client-form3/src/useCases"

type Gateway struct {
	Host   string
	Client useCases.ClientInterface
}
