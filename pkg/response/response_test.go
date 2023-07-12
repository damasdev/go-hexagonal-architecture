package response_test

import (
	"testing"

	"go-hexagonal-architecture/pkg/response"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	// Create a new instance of Response
	resp := response.New()

	// Set values for the response fields
	resp.Status.Code = 200
	resp.Status.Message = "Success"
	resp.Data = "Some data"
	resp.Meta = map[string]interface{}{"key": "value"}
	resp.Errors = nil

	// Assert the properties of the response
	assert.Equal(t, 200, resp.Status.Code)
	assert.Equal(t, "Success", resp.Status.Message)
	assert.Equal(t, "Some data", resp.Data)
	assert.Equal(t, map[string]interface{}{"key": "value"}, resp.Meta)
	assert.Nil(t, resp.Errors)
}
