package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
)

type CreatePlanogramInput struct {
	Name       string `json:"name"`
	NumPockets int32  `json:"num_pockets"`
	Notes      string `json:"notes"`
}

type UpdatePlanogramInput struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	NumPockets int32  `json:"num_pockets"`
	Notes      string `json:"notes"`
}

type CreatePlanogramPocketInput struct {
	PlanogramID  int32  `json:"planogram_id"`
	PocketNumber int32  `json:"pocket_number"`
	Category     string `json:"category"`
	ProductID    int32  `json:"product_id"`
}

type UpdatePlanogramPocketInput struct {
	ID        int32  `json:"id"`
	Category  string `json:"category"`
	ProductID int32  `json:"product_id"`
}

func GetPlanogram(ctx context.Context, db *database.Queries, id int32) (*database.Planogram, error) {
	planogram, err := db.GetPlanogram(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("planogram not found")
		}
		return nil, fmt.Errorf("failed to get planogram: %w", err)
	}
	return &planogram, nil
}

// ListPlanograms retrieves all planograms.
func ListPlanograms(ctx context.Context, db *database.Queries) ([]database.Planogram, error) {
	planograms, err := db.ListPlanograms(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list planograms: %w", err)
	}
	return planograms, nil
}

// CreatePlanogram creates a new planogram.
func CreatePlanogram(ctx context.Context, db *database.Queries, input CreatePlanogramInput) (*database.Planogram, error) {
	if input.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if input.NumPockets <= 0 {
		return nil, fmt.Errorf("num_pockets must be positive")
	}
	planogram, err := db.CreatePlanogram(ctx, database.CreatePlanogramParams{
		Name:       input.Name,
		NumPockets: input.NumPockets,
		Notes:      sql.NullString{String: input.Notes, Valid: input.Notes != ""},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create planogram: %w", err)
	}
	return &planogram, nil
}

// UpdatePlanogram updates an existing planogram.
func UpdatePlanogram(ctx context.Context, db *database.Queries, input UpdatePlanogramInput) (*database.Planogram, error) {
	if input.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if input.NumPockets <= 0 {
		return nil, fmt.Errorf("num_pockets must be positive")
	}
	planogram, err := db.UpdatePlanogram(ctx, database.UpdatePlanogramParams{
		ID:         input.ID,
		Name:       input.Name,
		NumPockets: input.NumPockets,
		Notes:      sql.NullString{String: input.Notes, Valid: input.Notes != ""},
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("planogram not found")
		}
		return nil, fmt.Errorf("failed to update planogram: %w", err)
	}
	return &planogram, nil
}

// DeletePlanogram deletes a planogram by its ID.
func DeletePlanogram(ctx context.Context, db *database.Queries, id int32) error {
	err := db.DeletePlanogram(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("planogram not found")
		}
		return fmt.Errorf("failed to delete planogram: %w", err)
	}
	return nil
}

// AssignPlanogramToLocation assigns a planogram to a customer location.
func AssignPlanogramToLocation(ctx context.Context, db *database.Queries, planogramID, customerID int32) (*database.PlanogramCustomerLocation, error) {
	if planogramID <= 0 || customerID <= 0 {
		return nil, fmt.Errorf("invalid planogram_id or customer_id")
	}
	pcl, err := db.AssignPlanogramToLocation(ctx, database.AssignPlanogramToLocationParams{
		PlanogramID:        planogramID,
		CustomerLocationID: customerID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to assign planogram to location: %w", err)
	}
	return &pcl, nil
}

// RemovePlanogramFromLocation removes a planogram from a customer location.
func RemovePlanogramFromLocation(ctx context.Context, db *database.Queries, planogramID, customerID int32) error {
	if planogramID <= 0 || customerID <= 0 {
		return fmt.Errorf("invalid planogram_id or customer_id")
	}
	err := db.RemovePlanogramFromLocation(ctx, database.RemovePlanogramFromLocationParams{
		PlanogramID:        planogramID,
		CustomerLocationID: customerID,
	})
	if err != nil {
		return fmt.Errorf("failed to remove planogram from location: %w", err)
	}
	return nil
}

// GetPlanogramsByLocation retrieves a planogram by customer location ID.
func GetPlanogramsByLocation(ctx context.Context, db *database.Queries, customerLocationID int32) (*database.Planogram, error) {
	if customerLocationID <= 0 {
		return nil, fmt.Errorf("invalid customer_location_id")
	}
	planogram, err := db.GetPlanogramsByLocation(ctx, customerLocationID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("planogram not found for location")
		}
		return nil, fmt.Errorf("failed to get planogram by location: %w", err)
	}
	return &planogram, nil
}

// ListLocationsByPlanogram lists all customer locations for a planogram.
func ListLocationsByPlanogram(ctx context.Context, db *database.Queries, planogramID int32) ([]database.CustomerLocation, error) {
	if planogramID <= 0 {
		return nil, fmt.Errorf("invalid planogram_id")
	}
	locations, err := db.ListLocationsByPlanogram(ctx, planogramID)
	if err != nil {
		return nil, fmt.Errorf("failed to list locations by planogram: %w", err)
	}
	return locations, nil
}

// ListPocketsForPlanogram lists all pockets for a planogram.
func ListPocketsForPlanogram(ctx context.Context, db *database.Queries, planogramID int32) ([]database.PlanogramPocket, error) {
	if planogramID <= 0 {
		return nil, fmt.Errorf("invalid planogram_id")
	}
	pockets, err := db.ListPocketsForPlanogram(ctx, planogramID)
	if err != nil {
		return nil, fmt.Errorf("failed to list pockets for planogram: %w", err)
	}
	return pockets, nil
}

// GetPlanogramPocket retrieves a planogram pocket by its ID.
func GetPlanogramPocket(ctx context.Context, db *database.Queries, id int32) (*database.PlanogramPocket, error) {
	pocket, err := db.GetPlanogramPocket(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("planogram pocket not found")
		}
		return nil, fmt.Errorf("failed to get planogram pocket: %w", err)
	}
	return &pocket, nil
}

// CreatePlanogramPocket creates a new planogram pocket.
func CreatePlanogramPocket(ctx context.Context, db *database.Queries, input CreatePlanogramPocketInput) (*database.PlanogramPocket, error) {
	if input.PlanogramID <= 0 || input.PocketNumber <= 0 || input.Category == "" {
		return nil, fmt.Errorf("invalid input for creating planogram pocket")
	}
	pocket, err := db.CreatePlanogramPocket(ctx, database.CreatePlanogramPocketParams{
		PlanogramID:  input.PlanogramID,
		PocketNumber: input.PocketNumber,
		Category:     input.Category,
		ProductID:    sql.NullInt32{Int32: input.ProductID, Valid: input.ProductID > 0},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create planogram pocket: %w", err)
	}
	return &pocket, nil
}

// UpdatePlanogramPocket updates a planogram pocket.
func UpdatePlanogramPocket(ctx context.Context, db *database.Queries, input UpdatePlanogramPocketInput) (*database.PlanogramPocket, error) {
	if input.ID <= 0 || input.Category == "" {
		return nil, fmt.Errorf("invalid input for updating planogram pocket")
	}
	pocket, err := db.UpdatePlanogramPocket(ctx, database.UpdatePlanogramPocketParams{
		ID:        input.ID,
		Category:  input.Category,
		ProductID: sql.NullInt32{Int32: input.ProductID, Valid: input.ProductID > 0},
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("planogram pocket not found")
		}
		return nil, fmt.Errorf("failed to update planogram pocket: %w", err)
	}
	return &pocket, nil
}

// DeletePlanogramPocket deletes a planogram pocket by its ID.
func DeletePlanogramPocket(ctx context.Context, db *database.Queries, id int32) error {
	err := db.DeletePlanogramPocket(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("planogram pocket not found")
		}
		return fmt.Errorf("failed to delete planogram pocket: %w", err)
	}
	return nil
}

// GetPlanogramPocketByNumber retrieves a planogram pocket by planogram ID and pocket number.
func GetPlanogramPocketByNumber(ctx context.Context, db *database.Queries, planogramID, pocketNumber int32) (*database.PlanogramPocket, error) {
	if planogramID <= 0 || pocketNumber <= 0 {
		return nil, fmt.Errorf("invalid planogram_id or pocket_number")
	}
	pocket, err := db.GetPlanogramPocketByNumber(ctx, database.GetPlanogramPocketByNumberParams{
		PlanogramID:  planogramID,
		PocketNumber: pocketNumber,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("planogram pocket not found")
		}
		return nil, fmt.Errorf("failed to get planogram pocket by number: %w", err)
	}
	return &pocket, nil
}
