package applicationclient

import (
	Instrumentation "rnoblega/client-form3/cmd/configuration"
)

func handleError(err error) {
	Instrumentation.TrackError(err.Error())
}
