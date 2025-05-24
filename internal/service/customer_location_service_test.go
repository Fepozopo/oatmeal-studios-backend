package service

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

// --- AddCustomerLocation tests ---
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

// --- DeleteCustomerLocation tests ---
func TestDeleteCustomerLocation_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()

	mock.ExpectExec(`-- name: DeleteCustomerLocation :exec`).
		WithArgs(int32(10)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := DeleteCustomerLocation(ctx, dbQueries, 10)
	if err != nil {
		t.Errorf("DeleteCustomerLocation returned error: %v", err)
	}
}

func TestDeleteCustomerLocation_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()

	// Test missing ID
	err := DeleteCustomerLocation(ctx, dbQueries, 0)
	if err == nil {
		t.Errorf("DeleteCustomerLocation should have returned an error for missing id")
	}

	// Test not found (sql.ErrNoRows)
	mock.ExpectExec(`-- name: DeleteCustomerLocation :exec`).
		WithArgs(int32(999)).
		WillReturnError(sql.ErrNoRows)

	err = DeleteCustomerLocation(ctx, dbQueries, 999)
	if err == nil {
		t.Errorf("DeleteCustomerLocation should have returned an error for not found")
	}
}

// --- UpdateCustomerLocation tests ---
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

// --- GetCustomerLocationByID tests ---
func TestGetCustomerLocationByID_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	currentTime := time.Now()

	mock.ExpectQuery(`-- name: GetCustomerLocationByID :one`).
		WithArgs(int32(10)).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "customer_id", "address_1", "address_2", "city", "state", "zip_code", "phone", "notes", "created_at", "updated_at",
		}).AddRow(
			10, 1, "123 Main St", "Apt 2", "Springfield", "IL", "62704", "+11234567890", "Front entrance", currentTime, currentTime,
		))

	location, err := GetCustomerLocationByID(ctx, dbQueries, 10)
	if err != nil {
		t.Errorf("GetCustomerLocationByID returned error: %v", err)
	}
	if location == nil {
		t.Errorf("GetCustomerLocationByID should have returned a non-nil location")
	}
	if location != nil && location.ID != 10 {
		t.Errorf("expected id 10, got %d", location.ID)
	}
}

func TestGetCustomerLocationByID_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()

	mock.ExpectQuery(`-- name: GetCustomerLocationByID :one`).
		WithArgs(int32(999)).
		WillReturnError(sql.ErrNoRows)

	location, err := GetCustomerLocationByID(ctx, dbQueries, 999)
	if err == nil {
		t.Errorf("GetCustomerLocationByID should have returned an error for not found")
	}
	if location != nil {
		t.Errorf("GetCustomerLocationByID should have returned nil for not found, got: %v", location)
	}
}

// --- ListCustomerLocations tests ---
func TestListCustomerLocations_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	currentTime := time.Now()

	mock.ExpectQuery(`-- name: ListCustomerLocationsByCustomer :many`).
		WithArgs(int32(1)).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "customer_id", "address_1", "address_2", "city", "state", "zip_code", "phone", "notes", "created_at", "updated_at",
		}).AddRow(
			10, 1, "123 Main St", "Apt 2", "Springfield", "IL", "62704", "+11234567890", "Front entrance", currentTime, currentTime,
		).AddRow(
			11, 1, "456 Oak Ave", sql.NullString{String: "", Valid: false}, "Springfield", "IL", "62705", sql.NullString{String: "", Valid: false}, sql.NullString{String: "", Valid: false}, currentTime, currentTime,
		))

	locations, err := ListCustomerLocations(ctx, dbQueries, 1)
	if err != nil {
		t.Errorf("ListCustomerLocations returned error: %v", err)
	}
	if locations == nil || len(locations) != 2 {
		t.Errorf("expected 2 locations, got %v", locations)
	}
	if locations != nil && locations[0].Address1 != "123 Main St" {
		t.Errorf("expected first address_1 '123 Main St', got '%s'", locations[0].Address1)
	}
}

func TestListCustomerLocations_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()

	// Test missing customerID
	locations, err := ListCustomerLocations(ctx, dbQueries, 0)
	if err == nil {
		t.Errorf("ListCustomerLocations should have returned an error for missing customer_id")
	}
	if locations != nil {
		t.Errorf("ListCustomerLocations should have returned nil for missing customer_id, got: %v", locations)
	}

	// Test no locations found
	mock.ExpectQuery(`-- name: ListCustomerLocationsByCustomer :many`).
		WithArgs(int32(999)).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "customer_id", "address_1", "address_2", "city", "state", "zip_code", "phone", "notes", "created_at", "updated_at",
		}))

	locations, err = ListCustomerLocations(ctx, dbQueries, 999)
	if err == nil {
		t.Errorf("ListCustomerLocations should have returned an error for no customer locations found")
	}
	if locations != nil {
		t.Errorf("ListCustomerLocations should have returned nil for no customer locations found, got: %v", locations)
	}
}
