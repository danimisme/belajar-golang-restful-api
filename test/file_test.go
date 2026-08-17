package test

import (
	"belajar-golang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializeConnection("Database")
	assert.NotNil(t, connection)

	cleanup()
}

