package service

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

// --- TestCreateCustomer ---
func TestCreateCustomer_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
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
			"Test Contact",
			"test@example.com",
			"1234567890",
			"123 Test St",
			"Suite 100",
			"Test City",
			"Test State",
			"12345",
			"Test Terms",
			10.0,
			5.0,
			"Test Sales Rep",
			"Test Notes",
		).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "business_name", "contact_name", "email", "phone", "address_1", "address_2", "city", "state", "zip_code", "terms", "discount", "commission", "sales_rep", "notes",
		}).AddRow(
			1, currentTime, currentTime, "Test Biz", "Test Contact", "test@example.com", "1234567890", "123 Test St", "Suite 100", "Test City", "Test State", "12345", "Test Terms", 10.0, 5.0, "Test Sales Rep", "Test Notes",
		))

	customer, err := CreateCustomer(ctx, dbQueries, CreateCustomerInput{
		BusinessName: "Test Biz",
		ContactName:  "Test Contact",
		Email:        "test@example.com",
		Phone:        "1234567890",
		Address1:     "123 Test St",
		Address2:     "Suite 100",
		City:         "Test City",
		State:        "Test State",
		ZipCode:      "12345",
		Terms:        "Test Terms",
		Discount:     10.0,
		Commission:   5.0,
		SalesRep:     "Test Sales Rep",
		Notes:        "Test Notes",
	})
	if err != nil {
		t.Errorf("CreateCustomer returned error: %v", err)
	}
	if customer == nil {
		t.Errorf("CreateCustomer should have returned a non-nil customer")
	}
	if customer != nil && customer.BusinessName != "Test Biz" {
		t.Errorf("expected business name 'Test Biz', got '%s'", customer.BusinessName)
	}
}

func TestCreateCustomer_InvalidInput(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	customer, err := CreateCustomer(ctx, dbQueries, CreateCustomerInput{
		BusinessName: "",
	})
	if err == nil {
		t.Errorf("CreateCustomer should have returned an error for missing business name")
	}
	if customer != nil {
		t.Errorf("CreateCustomer should have returned nil for missing business name, got: %v", customer)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// --- TestGetCustomerByID ---
func TestGetCustomerByID_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()

	// Mock customer data
	customerID := int32(1)
	mock.ExpectQuery(`SELECT id, created_at, updated_at, business_name, contact_name, email, phone, address_1, address_2, city, state, zip_code, terms, discount, commission, sales_rep, notes FROM customers WHERE id = \$1`).
		WithArgs(customerID).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "business_name", "contact_name", "email", "phone", "address_1", "address_2", "city", "state", "zip_code", "terms", "discount", "commission", "sales_rep", "notes",
		}).
			AddRow(1, time.Now(), time.Now(), "Test Biz", "Test Contact", "test@example.com", "1234567890", "123 Test St", "Suite 100", "Test City", "Test State", "12345", "Test Terms", 10.0, 5.0, "Test Sales Rep", "Test Notes"))

	customer, err := GetCustomerByID(ctx, dbQueries, customerID)
	if err != nil {
		t.Errorf("GetCustomerByID returned error: %v", err)
	}
	if customer == nil {
		t.Errorf("GetCustomerByID should have returned a non-nil customer")
	}
	if customer != nil && customer.ID != customerID {
		t.Errorf("expected customer ID '%d', got '%d'", customerID, customer.ID)
	}
}

func TestGetCustomerByID_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()

	// Mock not found scenario
	customer, err := GetCustomerByID(ctx, dbQueries, 999)
	if err == nil {
		t.Errorf("GetCustomerByID should have returned an error for not found")
	}
	if customer != nil {
		t.Errorf("GetCustomerByID should have returned nil for not found, got: %v", customer)
	}

	// Mock the database query for not found
	notFoundID := int32(999)
	mock.ExpectQuery(`SELECT id, created_at, updated_at, business_name, contact_name, email, phone, address_1, address_2, city, state, zip_code, terms, discount, commission, sales_rep, notes FROM customers WHERE id = \$1`).
		WithArgs(notFoundID).
		WillReturnError(sql.ErrNoRows)

	customer, err = GetCustomerByID(ctx, dbQueries, notFoundID)
	if err == nil {
		t.Errorf("GetCustomerByID should have returned an error for not found")
	}
	if err != nil && err.Error() != "customer not found" {
		t.Errorf("expected error 'customer not found', got '%s'", err.Error())
	}
	if customer != nil {
		t.Errorf("GetCustomerByID should have returned nil for not found, got: %v", customer)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// --- TestListCustomers ---
func TestListCustomers_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()

	mock.ExpectQuery(`SELECT id, created_at, updated_at, business_name, contact_name, email, phone, address_1, address_2, city, state, zip_code, terms, discount, commission, sales_rep, notes FROM customers`).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "business_name", "contact_name", "email", "phone", "address_1", "address_2", "city", "state", "zip_code", "terms", "discount", "commission", "sales_rep", "notes",
		}).AddRow(
			1, now, now, "Biz1", "Contact1", "email1@example.com", "1111111111", "Addr1", "Apt1", "City1", "State1", "11111", "Terms1", 10.0, 5.0, "Rep1", "Notes1",
		).AddRow(
			2, now, now, "Biz2", "Contact2", "email2@example.com", "2222222222", "Addr2", "Apt2", "City2", "State2", "22222", "Terms2", 20.0, 10.0, "Rep2", "Notes2",
		))

	customers, err := ListCustomers(ctx, dbQueries)
	if err != nil {
		t.Errorf("ListCustomers returned error: %v", err)
	}
	if customers == nil {
		t.Errorf("ListCustomers should have returned a non-nil slice")
	}
	if len(customers) != 2 {
		t.Errorf("expected 2 customers, got %d", len(customers))
	}
	if customers[0].BusinessName != "Biz1" || customers[1].BusinessName != "Biz2" {
		t.Errorf("unexpected business names: got %v, %v", customers[0].BusinessName, customers[1].BusinessName)
	}
}

func TestListCustomers_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()

	mock.ExpectQuery(`SELECT id, created_at, updated_at, business_name, contact_name, email, phone, address_1, address_2, city, state, zip_code, terms, discount, commission, sales_rep, notes FROM customers`).
		WillReturnError(sql.ErrConnDone)

	customers, err := ListCustomers(ctx, dbQueries)
	if err == nil {
		t.Errorf("ListCustomers should have returned an error on DB failure")
	}
	if customers != nil {
		t.Errorf("ListCustomers should have returned nil on DB failure, got: %v", customers)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
