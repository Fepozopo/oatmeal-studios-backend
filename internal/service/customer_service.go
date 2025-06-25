package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/auth"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
)

// CreateCustomer creates a new customer and returns the created customer.
func CreateCustomer(ctx context.Context, db *database.Queries, input CreateCustomerInput) (*database.Customer, error) {
	// Check for required fields
	if input.BusinessName == "" {
		return nil, errors.New("business name is required")
	}
	if input.Address1 == "" {
		return nil, errors.New("address1 is required")
	}
	if input.City == "" {
		return nil, errors.New("city is required")
	}
	if input.State == "" {
		return nil, errors.New("state is required")
	}
	if input.ZipCode == "" {
		return nil, errors.New("zip code is required")
	}
	// Validate discount and commission values
	if input.Discount < 0 || input.Discount > 100 {
		return nil, errors.New("discount must be between 0 and 100")
	}
	if input.Commission < 0 || input.Commission > 100 {
		return nil, errors.New("commission must be between 0 and 100")
	}
	// Validate email format if provided
	if input.Email != "" {
		if err := auth.IsValidEmail(input.Email); err != nil {
			return nil, fmt.Errorf("invalid email format: %w", err)
		}
	}
	// Validate phone format if provided
	if input.Phone != "" {
		if err := auth.IsValidPhone(input.Phone); err != nil {
			return nil, fmt.Errorf("invalid phone format: %w", err)
		}
	}

	// Prepare parameters for database insertion
	params := database.CreateCustomerParams{
		BusinessName: input.BusinessName,
		ContactName:  sql.NullString{String: input.ContactName, Valid: input.ContactName != ""},
		Email:        sql.NullString{String: input.Email, Valid: input.Email != ""},
		Phone:        sql.NullString{String: input.Phone, Valid: input.Phone != ""},
		Address1:     input.Address1,
		Address2:     sql.NullString{String: input.Address2, Valid: input.Address2 != ""},
		City:         input.City,
		State:        input.State,
		ZipCode:      input.ZipCode,
		Country:      input.Country,
		Terms:        input.Terms,
		Discount:     input.Discount,
		Commission:   input.Commission,
		Notes:        sql.NullString{String: input.Notes, Valid: input.Notes != ""},
	}

	customer, err := db.CreateCustomer(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create customer: %w", err)
	}

	return &customer, nil
}

// GetCustomerByID retrieves a customer by their ID using the GetCustomer query.
func GetCustomerByID(ctx context.Context, db *database.Queries, id int32) (*database.Customer, error) {
	customer, err := db.GetCustomer(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}
		return nil, fmt.Errorf("failed to get customer: %w", err)
	}
	return &customer, nil
}

// GetAllCustomers retrieves all customers from the database.
func ListCustomers(ctx context.Context, db *database.Queries) ([]database.Customer, error) {
	customers, err := db.ListCustomers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get customers: %w", err)
	}
	return customers, nil
}

// UpdateCustomer updates an existing customer's details by ID.
func UpdateCustomer(ctx context.Context, db *database.Queries, input UpdateCustomerInput) (*database.Customer, error) {
	// Check for required fields
	if input.BusinessName == "" {
		return nil, errors.New("business name is required")
	}
	if input.Address1 == "" {
		return nil, errors.New("address1 is required")
	}
	if input.City == "" {
		return nil, errors.New("city is required")
	}
	if input.State == "" {
		return nil, errors.New("state is required")
	}
	if input.ZipCode == "" {
		return nil, errors.New("zip code is required")
	}
	// Validate discount and commission values
	if input.Discount < 0 || input.Discount > 100 {
		return nil, errors.New("discount must be between 0 and 100")
	}
	if input.Commission < 0 || input.Commission > 100 {
		return nil, errors.New("commission must be between 0 and 100")
	}
	// Validate email format if provided
	if input.Email != "" {
		if err := auth.IsValidEmail(input.Email); err != nil {
			return nil, fmt.Errorf("invalid email format: %w", err)
		}
	}
	// Validate phone format if provided
	if input.Phone != "" {
		if err := auth.IsValidPhone(input.Phone); err != nil {
			return nil, fmt.Errorf("invalid phone format: %w", err)
		}
	}

	params := database.UpdateCustomerParams{
		ID:           input.ID,
		BusinessName: input.BusinessName,
		ContactName:  sql.NullString{String: input.ContactName, Valid: input.ContactName != ""},
		Email:        sql.NullString{String: input.Email, Valid: input.Email != ""},
		Phone:        sql.NullString{String: input.Phone, Valid: input.Phone != ""},
		Address1:     input.Address1,
		Address2:     sql.NullString{String: input.Address2, Valid: input.Address2 != ""},
		City:         input.City,
		State:        input.State,
		ZipCode:      input.ZipCode,
		Country:      input.Country,
		Terms:        input.Terms,
		Discount:     input.Discount,
		Commission:   input.Commission,
		Notes:        sql.NullString{String: input.Notes, Valid: input.Notes != ""},
	}

	customer, err := db.UpdateCustomer(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}
		return nil, fmt.Errorf("failed to update customer: %w", err)
	}
	return &customer, nil
}

// DeleteCustomer deletes a customer by their ID.
func DeleteCustomer(ctx context.Context, db *database.Queries, id int32) error {
	err := db.DeleteCustomer(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("customer not found")
		}
		return fmt.Errorf("failed to delete customer: %w", err)
	}
	return nil
}
