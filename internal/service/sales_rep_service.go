package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
)

// CreateSalesRep creates a new sales rep and returns the created sales rep.
func CreateSalesRep(ctx context.Context, db *database.Queries, input CreateSalesRepInput) (*database.SalesRep, error) {
	// Required fields
	if input.Status == "" {
		input.Status = "Active"
	}
	if input.FirstName == "" {
		return nil, errors.New("first name is required")
	}
	if input.LastName == "" {
		return nil, errors.New("last name is required")
	}
	if input.Company == "" {
		return nil, errors.New("company is required")
	}

	params := database.CreateSalesRepParams{
		Status:    input.Status,
		RepCode:   input.RepCode,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Company:   sql.NullString{String: input.Company, Valid: input.Company != ""},
		Address1:  sql.NullString{String: input.Address1, Valid: input.Address1 != ""},
		Address2:  sql.NullString{String: input.Address2, Valid: input.Address2 != ""},
		City:      sql.NullString{String: input.City, Valid: input.City != ""},
		State:     sql.NullString{String: input.State, Valid: input.State != ""},
		ZipCode:   sql.NullString{String: input.ZipCode, Valid: input.ZipCode != ""},
		Country:   sql.NullString{String: input.Country, Valid: input.Country != ""},
		Phone:     sql.NullString{String: input.Phone, Valid: input.Phone != ""},
		Email:     sql.NullString{String: input.Email, Valid: input.Email != ""},
	}

	rep, err := db.CreateSalesRep(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create sales rep: %w", err)
	}
	return &rep, nil
}

// GetSalesRepByID retrieves a sales rep by their ID.
func GetSalesRepByID(ctx context.Context, db *database.Queries, id int32) (*database.SalesRep, error) {
	rep, err := db.GetSalesRep(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("sales rep not found")
		}
		return nil, fmt.Errorf("failed to get sales rep: %w", err)
	}
	return &rep, nil
}

// ListSalesReps retrieves all sales reps from the database.
func ListSalesReps(ctx context.Context, db *database.Queries) ([]database.SalesRep, error) {
	reps, err := db.ListSalesReps(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list sales reps: %w", err)
	}
	return reps, nil
}

// UpdateSalesRep updates an existing sales rep's details by ID.
func UpdateSalesRep(ctx context.Context, db *database.Queries, input UpdateSalesRepInput) (*database.SalesRep, error) {
	// Required fields
	if input.Status == "" {
		input.Status = "Active"
	}
	if input.FirstName == "" {
		return nil, errors.New("first name is required")
	}
	if input.LastName == "" {
		return nil, errors.New("last name is required")
	}
	if input.Company == "" {
		return nil, errors.New("company is required")
	}

	params := database.UpdateSalesRepParams{
		Status:    input.Status,
		RepCode:   input.RepCode,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Company:   sql.NullString{String: input.Company, Valid: input.Company != ""},
		Address1:  sql.NullString{String: input.Address1, Valid: input.Address1 != ""},
		Address2:  sql.NullString{String: input.Address2, Valid: input.Address2 != ""},
		City:      sql.NullString{String: input.City, Valid: input.City != ""},
		State:     sql.NullString{String: input.State, Valid: input.State != ""},
		ZipCode:   sql.NullString{String: input.ZipCode, Valid: input.ZipCode != ""},
		Country:   sql.NullString{String: input.Country, Valid: input.Country != ""},
		Phone:     sql.NullString{String: input.Phone, Valid: input.Phone != ""},
		Email:     sql.NullString{String: input.Email, Valid: input.Email != ""},
	}

	rep, err := db.UpdateSalesRep(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("sales rep not found")
		}
		return nil, fmt.Errorf("failed to update sales rep: %w", err)
	}
	return &rep, nil
}

// DeleteSalesRep deletes a sales rep by their ID.
func DeleteSalesRep(ctx context.Context, db *database.Queries, id int32) error {
	err := db.DeleteSalesRep(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("sales rep not found")
		}
		return fmt.Errorf("failed to delete sales rep: %w", err)
	}
	return nil
}
