package gateway

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewGateway(t *testing.T) {
	gateway := NewGateway("localhost:8080", 4*time.Second)

	assert.Equal(t, "localhost:8080", gateway.Host)
}
