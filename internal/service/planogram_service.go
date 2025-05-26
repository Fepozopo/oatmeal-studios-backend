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

func (cfg *apiConfig) GetPlanogram(ctx context.Context, id int32) (*database.Planogram, error) {
	planogram, err := cfg.DbQueries.GetPlanogram(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("planogram not found")
		}
		return nil, fmt.Errorf("failed to get planogram: %w", err)
	}
	return &planogram, nil
}

// ListPlanograms retrieves all planograms.
func (cfg *apiConfig) ListPlanograms(ctx context.Context) ([]database.Planogram, error) {
	planograms, err := cfg.DbQueries.ListPlanograms(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list planograms: %w", err)
	}
	return planograms, nil
}

// CreatePlanogram creates a new planogram.
func (cfg *apiConfig) CreatePlanogram(ctx context.Context, input CreatePlanogramInput) (*database.Planogram, error) {
	if input.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if input.NumPockets <= 0 {
		return nil, fmt.Errorf("num_pockets must be positive")
	}
	planogram, err := cfg.DbQueries.CreatePlanogram(ctx, database.CreatePlanogramParams{
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
func (cfg *apiConfig) UpdatePlanogram(ctx context.Context, input UpdatePlanogramInput) (*database.Planogram, error) {
	if input.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if input.NumPockets <= 0 {
		return nil, fmt.Errorf("num_pockets must be positive")
	}
	planogram, err := cfg.DbQueries.UpdatePlanogram(ctx, database.UpdatePlanogramParams{
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
func (cfg *apiConfig) DeletePlanogram(ctx context.Context, id int32) error {
	err := cfg.DbQueries.DeletePlanogram(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("planogram not found")
		}
		return fmt.Errorf("failed to delete planogram: %w", err)
	}
	return nil
}

// AssignPlanogramToLocation assigns a planogram to a customer location.
func (cfg *apiConfig) AssignPlanogramToLocation(ctx context.Context, planogramID, customerID int32) (*database.PlanogramCustomerLocation, error) {
	if planogramID <= 0 || customerID <= 0 {
		return nil, fmt.Errorf("invalid planogram_id or customer_id")
	}
	pcl, err := cfg.DbQueries.AssignPlanogramToLocation(ctx, database.AssignPlanogramToLocationParams{
		PlanogramID:        planogramID,
		CustomerLocationID: customerID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to assign planogram to location: %w", err)
	}
	return &pcl, nil
}

// RemovePlanogramFromLocation removes a planogram from a customer location.
func (cfg *apiConfig) RemovePlanogramFromLocation(ctx context.Context, planogramID, customerID int32) error {
	if planogramID <= 0 || customerID <= 0 {
		return fmt.Errorf("invalid planogram_id or customer_id")
	}
	err := cfg.DbQueries.RemovePlanogramFromLocation(ctx, database.RemovePlanogramFromLocationParams{
		PlanogramID:        planogramID,
		CustomerLocationID: customerID,
	})
	if err != nil {
		return fmt.Errorf("failed to remove planogram from location: %w", err)
	}
	return nil
}

// GetPlanogramsByLocation retrieves a planogram by customer location ID.
func (cfg *apiConfig) GetPlanogramsByLocation(ctx context.Context, customerLocationID int32) (*database.Planogram, error) {
	if customerLocationID <= 0 {
		return nil, fmt.Errorf("invalid customer_location_id")
	}
	planogram, err := cfg.DbQueries.GetPlanogramsByLocation(ctx, customerLocationID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("planogram not found for location")
		}
		return nil, fmt.Errorf("failed to get planogram by location: %w", err)
	}
	return &planogram, nil
}

// ListLocationsByPlanogram lists all customer locations for a planogram.
func (cfg *apiConfig) ListLocationsByPlanogram(ctx context.Context, planogramID int32) ([]database.CustomerLocation, error) {
	if planogramID <= 0 {
		return nil, fmt.Errorf("invalid planogram_id")
	}
	locations, err := cfg.DbQueries.ListLocationsByPlanogram(ctx, planogramID)
	if err != nil {
		return nil, fmt.Errorf("failed to list locations by planogram: %w", err)
	}
	return locations, nil
}

// ListPocketsForPlanogram lists all pockets for a planogram.
func (cfg *apiConfig) ListPocketsForPlanogram(ctx context.Context, planogramID int32) ([]database.PlanogramPocket, error) {
	if planogramID <= 0 {
		return nil, fmt.Errorf("invalid planogram_id")
	}
	pockets, err := cfg.DbQueries.ListPocketsForPlanogram(ctx, planogramID)
	if err != nil {
		return nil, fmt.Errorf("failed to list pockets for planogram: %w", err)
	}
	return pockets, nil
}

// GetPlanogramPocket retrieves a planogram pocket by its ID.
func (cfg *apiConfig) GetPlanogramPocket(ctx context.Context, id int32) (*database.PlanogramPocket, error) {
	pocket, err := cfg.DbQueries.GetPlanogramPocket(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("planogram pocket not found")
		}
		return nil, fmt.Errorf("failed to get planogram pocket: %w", err)
	}
	return &pocket, nil
}

// CreatePlanogramPocket creates a new planogram pocket.
func (cfg *apiConfig) CreatePlanogramPocket(ctx context.Context, input CreatePlanogramPocketInput) (*database.PlanogramPocket, error) {
	if input.PlanogramID <= 0 || input.PocketNumber <= 0 || input.Category == "" {
		return nil, fmt.Errorf("invalid input for creating planogram pocket")
	}
	pocket, err := cfg.DbQueries.CreatePlanogramPocket(ctx, database.CreatePlanogramPocketParams{
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
func (cfg *apiConfig) UpdatePlanogramPocket(ctx context.Context, input UpdatePlanogramPocketInput) (*database.PlanogramPocket, error) {
	if input.ID <= 0 || input.Category == "" {
		return nil, fmt.Errorf("invalid input for updating planogram pocket")
	}
	pocket, err := cfg.DbQueries.UpdatePlanogramPocket(ctx, database.UpdatePlanogramPocketParams{
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
func (cfg *apiConfig) DeletePlanogramPocket(ctx context.Context, id int32) error {
	err := cfg.DbQueries.DeletePlanogramPocket(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("planogram pocket not found")
		}
		return fmt.Errorf("failed to delete planogram pocket: %w", err)
	}
	return nil
}

// GetPlanogramPocketByNumber retrieves a planogram pocket by planogram ID and pocket number.
func (cfg *apiConfig) GetPlanogramPocketByNumber(ctx context.Context, planogramID, pocketNumber int32) (*database.PlanogramPocket, error) {
	if planogramID <= 0 || pocketNumber <= 0 {
		return nil, fmt.Errorf("invalid planogram_id or pocket_number")
	}
	pocket, err := cfg.DbQueries.GetPlanogramPocketByNumber(ctx, database.GetPlanogramPocketByNumberParams{
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
