package useCases

import "log"

func trackError(err error) {
	log.Println(err.Error())
}
