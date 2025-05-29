package api

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func idFromURLAsInt32(r *http.Request) (int32, error) {
	id := r.URL.Query().Get("id")
	if id == "" {
		return 0, fmt.Errorf("ID is required")
	}

	var idInt int32
	_, err := fmt.Sscanf(id, "%d", &idInt)
	if err != nil {
		return 0, fmt.Errorf("Invalid ID format: %w", err)
	}

	return idInt, nil
}

func idFromURLAsUUID(r *http.Request) (uuid.UUID, error) {
	id := r.URL.Query().Get("id")
	if id == "" {
		return uuid.Nil, fmt.Errorf("ID is required")
	}

	// Validate the UUID format
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Invalid ID format: %w", err)
	}

	return uid, nil
}
