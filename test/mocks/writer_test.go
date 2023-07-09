package mocks_test

import (
	"bytes"
	"testing"

	"github.com/damasdev/fiber/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMockWriter(t *testing.T) {
	// Create a new instance of bytes.Buffer
	buffer := bytes.NewBuffer(nil)

	// Test NewMockWriter function
	writer := mocks.NewMockWriter(buffer)
	assert.NotNil(t, writer)

	// Test Write method of the returned writer
	message := []byte("Hello, World!")
	n, err := writer.Write(message)
	assert.NoError(t, err)
	assert.Equal(t, len(message), n)
	assert.Equal(t, message, buffer.Bytes())
}
