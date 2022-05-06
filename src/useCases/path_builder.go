package useCases

import (
	Configuration "github.com/rodrinoblega/client-form3/src/configuration"
	"strconv"
)

func ObtainFetchPath(id string) string {
	return Configuration.URI + "/" + id
}

func ObtainCreatePath() string {
	return Configuration.URI
}

func ObtainDeletePath(id string, version int64) string {
	return Configuration.URI + "/" + id + "?version=" + strconv.FormatInt(version, 10)
}
