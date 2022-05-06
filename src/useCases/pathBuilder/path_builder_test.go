package pathBuilder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObtainFetchPath(t *testing.T) {
	path := ObtainFetchPath("id")

	assert.Equal(t, "v1/organisation/accounts/id", path)
}

func TestObtainCreatePath(t *testing.T) {
	path := ObtainCreatePath()

	assert.Equal(t, "v1/organisation/accounts", path)
}

func TestObtainDeletePath(t *testing.T) {
	path := ObtainDeletePath("id", 0)

	assert.Equal(t, "v1/organisation/accounts/id?version=0", path)
}
