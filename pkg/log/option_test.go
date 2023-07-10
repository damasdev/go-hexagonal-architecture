package log_test

import (
	"errors"
	"testing"

	"github.com/damasdev/fiber/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	// Create a new instance of Option
	opt := log.Option{}

	// Test initial values
	assert.Nil(t, opt.GetData())
	assert.Nil(t, opt.GetError())

	// Test WithData function
	data := map[string]interface{}{"key": "value"}
	log.WithData(data)(&opt)
	assert.Equal(t, data, *opt.GetData())

	// Test WithError function
	err := errors.New("some error")
	log.WithError(err)(&opt)
	assert.Equal(t, err, *opt.GetError())
}
