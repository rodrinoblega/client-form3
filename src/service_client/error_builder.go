package service_client

import (
	Instrumentation "rnoblega/client-form3/src/configuration"
)

func handleError(err error) {
	Instrumentation.TrackError(err.Error())
}
