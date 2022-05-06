package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewGateway(t *testing.T) {
	service := CreateService("localhost:8080", 4*time.Second)

	assert.Equal(t, "localhost:8080", service.Host)
}
