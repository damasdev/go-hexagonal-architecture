package response_test

import (
	"go-hexagonal-architecture/internal/interfaces/http/response"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusText(t *testing.T) {
	// Test known status codes
	assert.Equal(t, "Continue", response.StatusText(response.StatusContinue))
	assert.Equal(t, "Switching Protocols", response.StatusText(response.StatusSwitchingProtocols))
	assert.Equal(t, "Processing", response.StatusText(response.StatusProcessing))
	assert.Equal(t, "Early Hints", response.StatusText(response.StatusEarlyHints))
	assert.Equal(t, "OK", response.StatusText(response.StatusOK))

	// Test unknown status code
	assert.Equal(t, "", response.StatusText(999))
}
