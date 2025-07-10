package api

import (
	"database/sql"
	"sync/atomic"
	"time"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
	"github.com/google/uuid"
)

type ApiConfig struct {
	FileserverHits atomic.Int32
	DbQueries      *database.Queries
	DB             *sql.DB
	Platform       string        `env:"PLATFORM"`
	TokenSecret    string        `env:"TOKEN_SECRET"`
	TokenExpiry    time.Duration `env:"TOKEN_EXPIRY"` // Duration for which the JWT token is valid
}

type UserResponse struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Token        string    `json:"token,omitempty"`         // Optional token field for authenticated users
	RefreshToken string    `json:"refresh_token,omitempty"` // Optional refresh token field
}

type NewJWT struct {
	Token     string `json:"token"`
	ExpiresIn int64  `json:"expires_in"`
}
