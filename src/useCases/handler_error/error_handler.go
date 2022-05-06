package handler_error

import (
	Instrumentation "github.com/rodrinoblega/client-form3/src/configuration"
)

func Handle(err error) {
	Instrumentation.TrackError(err.Error())
}
