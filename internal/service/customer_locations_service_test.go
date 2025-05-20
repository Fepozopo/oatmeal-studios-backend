package service

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddCustomerLocation_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	currentTime := time.Now()

	mock.ExpectQuery(`-- name: CreateCustomerLocation :one`).
		WithArgs(
			int32(1),
			"123 Main St",
			sql.NullString{String: "Apt 2", Valid: true},
			"Springfield",
			"IL",
			"62704",
			sql.NullString{String: "+11234567890", Valid: true},
			sql.NullString{String: "Front entrance", Valid: true},
		).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "customer_id", "address_1", "address_2", "city", "state", "zip_code", "phone", "notes", "created_at", "updated_at",
		}).AddRow(
			10, 1, "123 Main St", "Apt 2", "Springfield", "IL", "62704", "+11234567890", "Front entrance", currentTime, currentTime,
		))

	input := AddCustomerLocationInput{
		CustomerID: 1,
		Address1:   "123 Main St",
		Address2:   "Apt 2",
		City:       "Springfield",
		State:      "IL",
		ZipCode:    "62704",
		Phone:      "+11234567890",
		Notes:      "Front entrance",
	}

	location, err := AddCustomerLocation(ctx, dbQueries, input)
	if err != nil {
		t.Errorf("AddCustomerLocation returned error: %v", err)
	}
	if location == nil {
		t.Errorf("AddCustomerLocation should have returned a non-nil location")
	}
	if location != nil && location.Address1 != "123 Main St" {
		t.Errorf("expected address_1 '123 Main St', got '%s'", location.Address1)
	}
}

func TestAddCustomerLocation_Failure(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := newTestContext()

	// Missing required fields
	input := AddCustomerLocationInput{
		CustomerID: 0,
		Address1:   "",
		City:       "",
		State:      "",
		ZipCode:    "",
	}
	location, err := AddCustomerLocation(ctx, dbQueries, input)
	if err == nil {
		t.Errorf("AddCustomerLocation should have returned an error for missing required fields")
	}
	if location != nil {
		t.Errorf("AddCustomerLocation should have returned nil for missing required fields, got: %v", location)
	}
}

func TestUpdateCustomerLocation_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	currentTime := time.Now()

	mock.ExpectQuery(`-- name: UpdateCustomerLocation :one`).
		WithArgs(
			int32(10),
			"123 Main St",
			sql.NullString{String: "Apt 2", Valid: true},
			"Springfield",
			"IL",
			"62704",
			sql.NullString{String: "+11234567890", Valid: true},
			sql.NullString{String: "Front entrance", Valid: true},
		).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "customer_id", "address_1", "address_2", "city", "state", "zip_code", "phone", "notes", "created_at", "updated_at",
		}).AddRow(
			10, 1, "123 Main St", "Apt 2", "Springfield", "IL", "62704", "+11234567890", "Front entrance", currentTime, currentTime,
		))

	input := UpdateCustomerLocationInput{
		ID:       10,
		Address1: "123 Main St",
		Address2: "Apt 2",
		City:     "Springfield",
		State:    "IL",
		ZipCode:  "62704",
		Phone:    "+11234567890",
		Notes:    "Front entrance",
	}

	location, err := UpdateCustomerLocation(ctx, dbQueries, input)
	if err != nil {
		t.Errorf("UpdateCustomerLocation returned error: %v", err)
	}
	if location == nil {
		t.Errorf("UpdateCustomerLocation should have returned a non-nil location")
	}
	if location != nil && location.Address1 != "123 Main St" {
		t.Errorf("expected address_1 '123 Main St', got '%s'", location.Address1)
	}
}

func TestUpdateCustomerLocation_Failure(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := newTestContext()

	// Missing required fields
	input := UpdateCustomerLocationInput{
		ID:       0,
		Address1: "",
		City:     "",
		State:    "",
		ZipCode:  "",
	}
	location, err := UpdateCustomerLocation(ctx, dbQueries, input)
	if err == nil {
		t.Errorf("UpdateCustomerLocation should have returned an error for missing required fields")
	}
	if location != nil {
		t.Errorf("UpdateCustomerLocation should have returned nil for missing required fields, got: %v", location)
	}
}
