package handler_error

import (
	Instrumentation "rnoblega/client-form3/src/configuration"
)

func Handle(err error) {
	Instrumentation.TrackError(err.Error())
}
