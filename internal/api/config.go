package api

import (
	"sync/atomic"
	"time"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
	"github.com/google/uuid"
)

type ApiConfig struct {
	FileserverHits atomic.Int32
	DbQueries      *database.Queries
	Platform       string `env:"PLATFORM"`
	TokenSecret    string `env:"TOKEN_SECRET"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}
