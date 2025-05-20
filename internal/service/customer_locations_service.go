package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/auth"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
)

// AddCustomerLocation adds a new customer location and returns the created location.
func AddCustomerLocation(ctx context.Context, db *database.Queries, input AddCustomerLocationInput) (*database.CustomerLocation, error) {
	// Check if the input is valid
	if input.CustomerID == 0 {
		return nil, errors.New("customer_id is required")
	}
	// Check for required fields
	if input.Address1 == "" {
		return nil, errors.New("address_1 is required")
	}
	if input.City == "" {
		return nil, errors.New("city is required")
	}
	if input.State == "" {
		return nil, errors.New("state is required")
	}
	if input.ZipCode == "" {
		return nil, errors.New("zip_code is required")
	}
	// Validate phone number if provided
	if input.Phone != "" {
		if err := auth.IsValidPhone(input.Phone); err != nil {
			return nil, fmt.Errorf("invalid phone format: %w", err)
		}
	}

	params := database.CreateCustomerLocationParams{
		CustomerID: input.CustomerID,
		Address1:   input.Address1,
		Address2:   sql.NullString{String: input.Address2, Valid: input.Address2 != ""},
		City:       input.City,
		State:      input.State,
		ZipCode:    input.ZipCode,
		Phone:      sql.NullString{String: input.Phone, Valid: input.Phone != ""},
		Notes:      sql.NullString{String: input.Notes, Valid: input.Notes != ""},
	}

	location, err := db.CreateCustomerLocation(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create customer location: %w", err)
	}
	return &location, nil
}
