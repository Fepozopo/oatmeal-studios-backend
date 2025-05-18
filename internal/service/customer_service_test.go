package service

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
	"github.com/stretchr/testify/assert"
)

// TestCreateCustomer tests the CreateCustomer function.
func TestCreateCustomer(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	ctx := context.Background()
	currentTime := time.Now()

	// Mock customer data
	mock.ExpectQuery(`
		-- name: CreateCustomer :one
		INSERT INTO customers (.+)
		VALUES (.+)
		RETURNING id, created_at, updated_at, business_name, contact_name, email, phone, address_1, address_2, city, state, zip_code, terms, discount, commission, sales_rep, notes
	`).
		WithArgs(
			"Test Biz",
			sql.NullString{String: "", Valid: false},
			sql.NullString{String: "", Valid: false},
			sql.NullString{String: "", Valid: false},
			sql.NullString{String: "", Valid: false},
			sql.NullString{String: "", Valid: false},
			sql.NullString{String: "", Valid: false},
			sql.NullString{String: "", Valid: false},
			sql.NullString{String: "", Valid: false},
			sql.NullString{String: "", Valid: false},
			0.0,
			0.0,
			sql.NullString{String: "", Valid: false},
			sql.NullString{String: "", Valid: false},
		).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "business_name", "contact_name", "email", "phone", "address_1", "address_2", "city", "state", "zip_code", "terms", "discount", "commission", "sales_rep", "notes",
		}).AddRow(
			1, currentTime, currentTime, "Test Biz", nil, nil, nil, nil, nil, nil, nil, nil, nil, 0.0, 0.0, nil, nil,
		))

	customer, err := CreateCustomer(ctx, dbQueries, CreateCustomerInput{
		BusinessName: "Test Biz",
	})
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Test Biz", customer.BusinessName)

	// Test for missing business name
	customer, err = CreateCustomer(ctx, dbQueries, CreateCustomerInput{
		BusinessName: "",
	})
	assert.Error(t, err)
	assert.Nil(t, customer)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// TestGetCustomerByID tests the GetCustomerByID function.
func TestGetCustomerByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	ctx := context.Background()

	// Mock customer data
	customerID := int32(1)
	mock.ExpectQuery(`SELECT id, created_at, updated_at, business_name, contact_name, email, phone, address_1, address_2, city, state, zip_code, terms, discount, commission, sales_rep, notes FROM customers WHERE id = \$1`).
		WithArgs(customerID).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "business_name", "contact_name", "email", "phone", "address_1", "address_2", "city", "state", "zip_code", "terms", "discount", "commission", "sales_rep", "notes",
		}).
			AddRow(1, time.Now(), time.Now(), "Test Biz", sql.NullString{}, sql.NullString{}, sql.NullString{}, sql.NullString{}, sql.NullString{}, sql.NullString{}, sql.NullString{}, sql.NullString{}, sql.NullString{}, 0.0, 0.0, sql.NullString{}, sql.NullString{}))

	customer, err := GetCustomerByID(ctx, dbQueries, customerID)
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, customerID, customer.ID)

	// Not found case
	notFoundID := int32(999)
	mock.ExpectQuery(`SELECT id, created_at, updated_at, business_name, contact_name, email, phone, address_1, address_2, city, state, zip_code, terms, discount, commission, sales_rep, notes FROM customers WHERE id = \$1`).
		WithArgs(notFoundID).
		WillReturnError(sql.ErrNoRows)

	customer, err = GetCustomerByID(ctx, dbQueries, notFoundID)
	assert.Error(t, err)
	assert.EqualError(t, err, "customer not found")
	assert.Nil(t, customer)
}
