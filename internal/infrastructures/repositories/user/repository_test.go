package user_test

import (
	"testing"

	port "github.com/damasdev/go-hexagonal-architecture/internal/core/ports/user"
	repository "github.com/damasdev/go-hexagonal-architecture/internal/infrastructures/repositories/user"
)

// TestNew tests the New function
func TestNew(t *testing.T) {
	r := repository.New()

	// Ensure that the returned repository is not nil
	if r == nil {
		t.Error("Expected non-nil value, got nil")
	}

	// Ensure that the returned repository implements the ports.Repository interface
	_, ok := any(r).(port.Repository)
	if !ok {
		t.Error("New did not return a valid user.Repository implementation")
	}
}

// BenchmarkNew benchmarks the New function
func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = repository.New()
	}
}
