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
	if input.BusinessName == "" {
		return nil, errors.New("business_name is required")
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
		CustomerID:   input.CustomerID,
		BusinessName: input.BusinessName,
		ContactName:  sql.NullString{String: input.ContactName, Valid: input.ContactName != ""},
		Address1:     input.Address1,
		Address2:     sql.NullString{String: input.Address2, Valid: input.Address2 != ""},
		City:         input.City,
		State:        input.State,
		ZipCode:      input.ZipCode,
		Country:      input.Country,
		Phone:        sql.NullString{String: input.Phone, Valid: input.Phone != ""},
		Notes:        sql.NullString{String: input.Notes, Valid: input.Notes != ""},
	}

	location, err := db.CreateCustomerLocation(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create customer location: %w", err)
	}
	return &location, nil
}

// DeleteCustomerLocation deletes a customer location by its ID.
// It returns nil if successful, or an error if not.
func DeleteCustomerLocation(ctx context.Context, db *database.Queries, id int32) error {
	if id == 0 {
		return errors.New("id is required")
	}
	err := db.DeleteCustomerLocation(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("customer location not found")
		}
		return fmt.Errorf("failed to delete customer location: %w", err)
	}
	return nil
}

// UpdateCustomerLocation updates an existing customer location and returns the updated location.
func UpdateCustomerLocation(ctx context.Context, db *database.Queries, input UpdateCustomerLocationInput) (*database.CustomerLocation, error) {
	// Check if the input is valid
	if input.ID == 0 {
		return nil, errors.New("id is required")
	}
	if input.BusinessName == "" {
		return nil, errors.New("business_name is required")
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
	if input.Phone != "" {
		if err := auth.IsValidPhone(input.Phone); err != nil {
			return nil, fmt.Errorf("invalid phone format: %w", err)
		}
	}

	params := database.UpdateCustomerLocationParams{
		ID:           input.ID,
		BusinessName: input.BusinessName,
		ContactName:  sql.NullString{String: input.ContactName, Valid: input.ContactName != ""},
		Address1:     input.Address1,
		Address2:     sql.NullString{String: input.Address2, Valid: input.Address2 != ""},
		City:         input.City,
		State:        input.State,
		ZipCode:      input.ZipCode,
		Country:      input.Country,
		Phone:        sql.NullString{String: input.Phone, Valid: input.Phone != ""},
		Notes:        sql.NullString{String: input.Notes, Valid: input.Notes != ""},
	}

	location, err := db.UpdateCustomerLocation(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update customer location: %w", err)
	}
	return &location, nil
}

// GetCustomerLocationByID retrieves a customer location by its ID.
// It returns the location if found, or an error if not.
func GetCustomerLocationByID(ctx context.Context, db *database.Queries, id int32) (*database.CustomerLocation, error) {
	location, err := db.GetCustomerLocationByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer location not found")
		}
		return nil, fmt.Errorf("failed to get customer location: %w", err)
	}
	return &location, nil
}

// ListCustomerLocations retrieves all locations for a given customer ID.
// It returns a slice of locations if found, or an error if not.
func ListCustomerLocations(ctx context.Context, db *database.Queries, customerID int32) ([]database.CustomerLocation, error) {
	if customerID == 0 {
		return nil, errors.New("customer_id is required")
	}
	locations, err := db.ListCustomerLocationsByCustomer(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to list customer locations: %w", err)
	}
	if len(locations) == 0 {
		return nil, errors.New("no customer locations found")
	}
	return locations, nil
}
