package api

import (
	"sync/atomic"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
)

type ApiConfig struct {
	FileserverHits atomic.Int32
	DbQueries      *database.Queries
	Platform       string `env:"PLATFORM"`
	TokenSecret    string `env:"TOKEN_SECRET"`
}
