package service

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
)

func newTestDBSalesRep(t *testing.T) (*database.Queries, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	return database.New(db), mock
}

func newTestContextSalesRep() context.Context {
	return context.Background()
}

func TestCreateSalesRep_Success(t *testing.T) {
	dbQueries, mock := newTestDBSalesRep(t)
	ctx := newTestContextSalesRep()
	currentTime := time.Now()

	mock.ExpectQuery(`-- name: CreateSalesRep :one`).
		WithArgs(
			"Active",
			"John",
			"Doe",
			"Acme Inc",
			sql.NullString{String: "123 Main St", Valid: true},
			sql.NullString{String: "Suite 100", Valid: true},
			sql.NullString{String: "Metropolis", Valid: true},
			sql.NullString{String: "NY", Valid: true},
			sql.NullString{String: "10001", Valid: true},
		).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "status", "first_name", "last_name", "company", "address_1", "address_2", "city", "state", "zip_code", "email", "phone",
		}).AddRow(
			1, currentTime, currentTime, "active", "John", "Doe", "Acme Inc", "123 Main St", "Suite 100", "Metropolis", "NY", "10001", nil, nil,
		))

	input := CreateOrUpdateSalesRepInput{
		Status:    "Active",
		FirstName: "John",
		LastName:  "Doe",
		Company:   "Acme Inc",
		Address1:  "123 Main St",
		Address2:  "Suite 100",
		City:      "Metropolis",
		State:     "NY",
		ZipCode:   "10001",
	}

	rep, err := CreateSalesRep(ctx, dbQueries, input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if rep.FirstName != "John" || rep.LastName != "Doe" {
		t.Errorf("unexpected sales rep returned: %+v", rep)
	}
}

func TestGetSalesRepByID_NotFound(t *testing.T) {
	dbQueries, mock := newTestDBSalesRep(t)
	ctx := newTestContextSalesRep()

	mock.ExpectQuery(`-- name: GetSalesRep :one`).WithArgs(42).
		WillReturnError(sql.ErrNoRows)

	_, err := GetSalesRepByID(ctx, dbQueries, 42)
	if err == nil || err.Error() != "sales rep not found" {
		t.Errorf("expected not found error, got %v", err)
	}
}

func TestListSalesReps_Empty(t *testing.T) {
	dbQueries, mock := newTestDBSalesRep(t)
	ctx := newTestContextSalesRep()

	mock.ExpectQuery(`-- name: ListSalesReps :many`).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "status", "first_name", "last_name", "company", "address_1", "address_2", "city", "state", "zip_code", "email", "phone",
		}))

	reps, err := ListSalesReps(ctx, dbQueries)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(reps) != 0 {
		t.Errorf("expected 0 sales reps, got %d", len(reps))
	}
}

func TestUpdateSalesRep_NotFound(t *testing.T) {
	dbQueries, mock := newTestDBSalesRep(t)
	ctx := newTestContextSalesRep()

	mock.ExpectQuery(`-- name: UpdateSalesRep :one`).WithArgs(99, "Active", "Jane", "Smith", "Beta LLC", sql.NullString{String: "", Valid: false}, sql.NullString{String: "", Valid: false}, sql.NullString{String: "", Valid: false}, sql.NullString{String: "", Valid: false}, sql.NullString{String: "", Valid: false}).
		WillReturnError(sql.ErrNoRows)

	input := CreateOrUpdateSalesRepInput{
		Status:    "Active",
		FirstName: "Jane",
		LastName:  "Smith",
		Company:   "Beta LLC",
	}

	_, err := UpdateSalesRep(ctx, dbQueries, 99, input)
	if err == nil || err.Error() != "sales rep not found" {
		t.Errorf("expected not found error, got %v", err)
	}
}

func TestDeleteSalesRep_Success(t *testing.T) {
	dbQueries, mock := newTestDBSalesRep(t)
	ctx := newTestContextSalesRep()

	mock.ExpectExec(`-- name: DeleteSalesRep :exec`).WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := DeleteSalesRep(ctx, dbQueries, 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
