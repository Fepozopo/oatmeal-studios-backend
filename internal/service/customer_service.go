package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/auth"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
)

type CreateOrUpdateCustomerInput struct {
	BusinessName string  `json:"business_name"`
	ContactName  string  `json:"contact_name,omitempty"`
	Email        string  `json:"email,omitempty"`
	Phone        string  `json:"phone,omitempty"`
	Address1     string  `json:"address1,omitempty"`
	Address2     string  `json:"address2,omitempty"`
	City         string  `json:"city,omitempty"`
	State        string  `json:"state,omitempty"`
	ZipCode      string  `json:"zip_code,omitempty"`
	Terms        string  `json:"terms,omitempty"`
	Discount     float64 `json:"discount,omitempty"`
	Commission   float64 `json:"commission,omitempty"`
	SalesRep     string  `json:"sales_rep,omitempty"`
	Notes        string  `json:"notes,omitempty"`
}

// CreateCustomer creates a new customer and returns the created customer.
func (cfg *apiConfig) CreateCustomer(ctx context.Context, input CreateOrUpdateCustomerInput) (*database.Customer, error) {
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
		Terms:        input.Terms,
		Discount:     input.Discount,
		Commission:   input.Commission,
		SalesRep:     sql.NullString{String: input.SalesRep, Valid: input.SalesRep != ""},
		Notes:        sql.NullString{String: input.Notes, Valid: input.Notes != ""},
	}

	customer, err := cfg.DbQueries.CreateCustomer(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create customer: %w", err)
	}

	return &customer, nil
}

// GetCustomerByID retrieves a customer by their ID using the GetCustomer query.
func (cfg *apiConfig) GetCustomerByID(ctx context.Context, id int32) (*database.Customer, error) {
	customer, err := cfg.DbQueries.GetCustomer(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}
		return nil, fmt.Errorf("failed to get customer: %w", err)
	}
	return &customer, nil
}

// GetAllCustomers retrieves all customers from the database.
func (cfg *apiConfig) ListCustomers(ctx context.Context) ([]database.Customer, error) {
	customers, err := cfg.DbQueries.ListCustomers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get customers: %w", err)
	}
	return customers, nil
}

// UpdateCustomer updates an existing customer's details by ID.
func (cfg *apiConfig) UpdateCustomer(ctx context.Context, id int32, input CreateOrUpdateCustomerInput) (*database.Customer, error) {
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
		ID:           id,
		BusinessName: input.BusinessName,
		ContactName:  sql.NullString{String: input.ContactName, Valid: input.ContactName != ""},
		Email:        sql.NullString{String: input.Email, Valid: input.Email != ""},
		Phone:        sql.NullString{String: input.Phone, Valid: input.Phone != ""},
		Address1:     input.Address1,
		Address2:     sql.NullString{String: input.Address2, Valid: input.Address2 != ""},
		City:         input.City,
		State:        input.State,
		ZipCode:      input.ZipCode,
		Terms:        input.Terms,
		Discount:     input.Discount,
		Commission:   input.Commission,
		SalesRep:     sql.NullString{String: input.SalesRep, Valid: input.SalesRep != ""},
		Notes:        sql.NullString{String: input.Notes, Valid: input.Notes != ""},
	}

	customer, err := cfg.DbQueries.UpdateCustomer(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}
		return nil, fmt.Errorf("failed to update customer: %w", err)
	}
	return &customer, nil
}

// DeleteCustomer deletes a customer by their ID.
func (cfg *apiConfig) DeleteCustomer(ctx context.Context, id int32) error {
	err := cfg.DbQueries.DeleteCustomer(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("customer not found")
		}
		return fmt.Errorf("failed to delete customer: %w", err)
	}
	return nil
}
