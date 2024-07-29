package user_test

import (
	"encoding/json"
	"testing"

	domain "github.com/damasdev/go-hexagonal-architecture/internal/core/domain/user"
)

// TestUser tests the User domain model
func TestUser(t *testing.T) {
	user := domain.User{
		ID:   1,
		Name: "John Doe",
	}

	// Test JSON marshaling
	data, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Failed to marshal user: %v", err)
	}

	expected := `{"id":1,"name":"John Doe"}`
	if string(data) != expected {
		t.Errorf("Unexpected marshaled JSON. Got: %s, want: %s", string(data), expected)
	}

	// Test JSON unmarshaling
	var decodedUser domain.User
	err = json.Unmarshal([]byte(expected), &decodedUser)
	if err != nil {
		t.Errorf("Failed to unmarshal user: %v", err)
	}

	if decodedUser != user {
		t.Errorf("Unexpected unmarshaled user. Got: %+v, want: %+v", decodedUser, user)
	}
}
