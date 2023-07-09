package log_test

import (
	"errors"
	"testing"

	"github.com/damasdev/fiber/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	// Create a new instance of Options
	opts := log.Options{}

	// Test initial values
	assert.Nil(t, opts.GetData())
	assert.Nil(t, opts.GetError())

	// Test WithData function
	data := map[string]interface{}{"key": "value"}
	log.WithData(data)(&opts)
	assert.Equal(t, data, *opts.GetData())

	// Test WithError function
	err := errors.New("some error")
	log.WithError(err)(&opts)
	assert.Equal(t, err, *opts.GetError())
}
