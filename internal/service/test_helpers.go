package service

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
	"github.com/google/uuid"
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

func makeProductRow(id uuid.UUID, input CreateOrUpdateProductInput, now time.Time) *sqlmock.Rows {
	return sqlmock.NewRows([]string{
		"id", "created_at", "updated_at", "type", "sku", "upc", "status", "cost", "price", "envelope", "artist", "category", "release_date", "last_bought_date", "description", "text_front", "text_inside",
	}).AddRow(
		id, now, now, input.Type, input.Sku, input.Upc, input.Status, input.Cost, input.Price,
		sql.NullString{String: input.Envelope, Valid: input.Envelope != ""}, sql.NullString{String: input.Artist, Valid: input.Artist != ""}, sql.NullString{String: input.Category, Valid: input.Category != ""}, sql.NullTime{Time: input.ReleaseDate, Valid: !input.ReleaseDate.IsZero()}, sql.NullTime{Time: input.LastBoughtDate, Valid: !input.LastBoughtDate.IsZero()}, sql.NullString{String: input.Description, Valid: input.Description != ""}, sql.NullString{String: input.TextFront, Valid: input.TextFront != ""}, sql.NullString{String: input.TextInside, Valid: input.TextInside != ""},
	)
}
