package useCases

func Handle(err error) {
	TrackError(err.Error())
}
