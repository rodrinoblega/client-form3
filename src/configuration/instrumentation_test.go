package configuration

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestTrackError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	err := errors.New("mocked error")
	TrackError(err)

	assert.Contains(t, buf.String(), "mocked error")
}
