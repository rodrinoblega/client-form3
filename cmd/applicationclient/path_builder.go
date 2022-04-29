package applicationclient

import (
	Configuration "rnoblega/client-form3/cmd/configuration"
	"strconv"
)

func obtainFetchPath(id string) string {
	return Configuration.URI + "/" + id
}

func obtainCreatePath() string {
	return Configuration.URI
}

func obtainDeletePath(id string, version int64) string {
	return Configuration.URI + "/" + id + "?version=" + strconv.FormatInt(version, 10)
}
