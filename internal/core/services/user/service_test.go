package user_test

import (
	"testing"

	port "github.com/damasdev/go-hexagonal-architecture/internal/core/ports/user"
	service "github.com/damasdev/go-hexagonal-architecture/internal/core/services/user"
)

// TestNew tests the New function
func TestNew(t *testing.T) {
	// Define mock dependencies
	repository := &mockRepository{}

	// Define test cases
	testCases := []struct {
		Name   string
		Config service.Config
		IsErr  bool
	}{
		{
			Name: "Test Missing Repository",
			Config: service.Config{
				Repository: nil,
			},
			IsErr: true,
		},
		{
			Name: "Test Valid Repository",
			Config: service.Config{
				Repository: repository,
			},
			IsErr: false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			runNewTestCase(t, tc)
		})
	}
}

// runNewTestCase is a helper function to run test cases for the New function
func runNewTestCase(t *testing.T, tc struct {
	Name   string
	Config service.Config
	IsErr  bool
}) {
	// Call the New function
	s := service.New(tc.Config)

	// Ensure that the returned service satisfies the expected behavior
	if tc.IsErr && s != nil {
		t.Error("Expected nil value, got non-nil")
	}

	// Ensure that the returned service satisfies the expected behavior
	if !tc.IsErr {
		// Ensure that the returned service is not nil
		if s == nil {
			t.Error("Expected non-nil value, got nil")
		}

		// Ensure that the returned service implements the ports.Service interface
		_, ok := any(s).(port.Service)
		if !tc.IsErr && !ok {
			t.Error("New did not return a valid user.Service implementation")
		}
	}
}

type mockRepository struct{}
