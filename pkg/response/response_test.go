package response_test

import (
	"testing"

	"github.com/damasdev/fiber/pkg/response"
	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	// Create a new instance of Response
	response := response.New()

	// Set values for the response fields
	response.Status.Code = 200
	response.Status.Message = "Success"
	response.Data = "Some data"
	response.Meta = map[string]interface{}{"key": "value"}
	response.Errors = nil

	// Assert the properties of the response
	assert.Equal(t, 200, response.Status.Code)
	assert.Equal(t, "Success", response.Status.Message)
	assert.Equal(t, "Some data", response.Data)
	assert.Equal(t, map[string]interface{}{"key": "value"}, response.Meta)
	assert.Nil(t, response.Errors)
}
