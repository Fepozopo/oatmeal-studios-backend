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
	// Check if BusinessName is provided
	if input.BusinessName == "" {
		return nil, errors.New("business name is required")
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
		Address1:     sql.NullString{String: input.Address1, Valid: input.Address1 != ""},
		Address2:     sql.NullString{String: input.Address2, Valid: input.Address2 != ""},
		City:         sql.NullString{String: input.City, Valid: input.City != ""},
		State:        sql.NullString{String: input.State, Valid: input.State != ""},
		ZipCode:      sql.NullString{String: input.ZipCode, Valid: input.ZipCode != ""},
		Terms:        sql.NullString{String: input.Terms, Valid: input.Terms != ""},
		Discount:     input.Discount,
		Commission:   input.Commission,
		SalesRep:     sql.NullString{String: input.SalesRep, Valid: input.SalesRep != ""},
		Notes:        sql.NullString{String: input.Notes, Valid: input.Notes != ""},
	}

	customer, err := db.CreateCustomer(ctx, params)
	if err != nil {
		return nil, err
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
