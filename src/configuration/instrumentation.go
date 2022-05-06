package configuration

import "log"

func TrackError(err error) {
	log.Println(err.Error())
}
