package service

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
)

// --- Helper Functions ---
func newTestDB(t *testing.T) (*database.Queries, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	dbQueries := database.New(db)

	return dbQueries, mock
}

func newTestContext() context.Context {
	return context.Background()
}
