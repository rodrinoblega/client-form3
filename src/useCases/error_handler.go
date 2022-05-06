package useCases

import "github.com/rodrinoblega/client-form3/src/configuration"

func Handle(err error) {
	configuration.TrackError(err.Error())
}
