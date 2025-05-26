package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
	"github.com/google/uuid"
)

type CreateOrUpdateProductInput struct {
	Type           string    `json:"type"`
	Sku            string    `json:"sku"`
	Upc            string    `json:"upc"`
	Status         string    `json:"status"`
	Cost           float64   `json:"cost"`
	Price          float64   `json:"price"`
	Envelope       string    `json:"envelope,omitempty"`
	Artist         string    `json:"artist,omitempty"`
	Category       string    `json:"category,omitempty"`
	ReleaseDate    time.Time `json:"release_date,omitempty"`
	LastBoughtDate time.Time `json:"last_bought_date,omitempty"`
	Description    string    `json:"description,omitempty"`
	TextFront      string    `json:"text_front,omitempty"`
	TextInside     string    `json:"text_inside,omitempty"`
}

func (cfg *apiConfig) CreateProduct(ctx context.Context, input CreateOrUpdateProductInput) (*database.Product, error) {
	if input.Type == "" || input.Sku == "" || input.Upc == "" || input.Status == "" {
		return nil, errors.New("type, sku, upc, and status are required")
	}
	params := database.CreateProductParams{
		Type:           input.Type,
		Sku:            input.Sku,
		Upc:            input.Upc,
		Status:         input.Status,
		Cost:           input.Cost,
		Price:          input.Price,
		Envelope:       sql.NullString{String: input.Envelope, Valid: input.Envelope != ""},
		Artist:         sql.NullString{String: input.Artist, Valid: input.Artist != ""},
		Category:       sql.NullString{String: input.Category, Valid: input.Category != ""},
		ReleaseDate:    sql.NullTime{Time: input.ReleaseDate, Valid: !input.ReleaseDate.IsZero()},
		LastBoughtDate: sql.NullTime{Time: input.LastBoughtDate, Valid: !input.LastBoughtDate.IsZero()},
		Description:    sql.NullString{String: input.Description, Valid: input.Description != ""},
		TextFront:      sql.NullString{String: input.TextFront, Valid: input.TextFront != ""},
		TextInside:     sql.NullString{String: input.TextInside, Valid: input.TextInside != ""},
	}

	prod, err := cfg.DbQueries.CreateProduct(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}
	return &prod, nil
}

func (cfg *apiConfig) GetProductByID(ctx context.Context, id uuid.UUID) (*database.Product, error) {
	prod, err := cfg.DbQueries.GetProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get product by ID: %w", err)
	}
	return &prod, nil
}

func (cfg *apiConfig) GetProductBySKU(ctx context.Context, sku string) (*database.Product, error) {
	prod, err := cfg.DbQueries.GetProductBySKU(ctx, sku)
	if err != nil {
		return nil, fmt.Errorf("failed to get product by SKU: %w", err)
	}
	return &prod, nil
}

func (cfg *apiConfig) ListProductsByType(ctx context.Context, typ string) ([]database.Product, error) {
	prod, err := cfg.DbQueries.ListProductsByType(ctx, typ)
	if err != nil {
		return nil, fmt.Errorf("failed to list products by type: %w", err)
	}
	return prod, nil
}

func (cfg *apiConfig) ListProductsByCategory(ctx context.Context, category string) ([]database.Product, error) {
	prod, err := cfg.DbQueries.ListProductsByCategory(ctx, sql.NullString{String: category, Valid: category != ""})
	if err != nil {
		return nil, fmt.Errorf("failed to list products by category: %w", err)
	}
	return prod, nil
}

func (cfg *apiConfig) ListProductsByArtist(ctx context.Context, artist string) ([]database.Product, error) {
	prod, err := cfg.DbQueries.ListProductsByArtist(ctx, sql.NullString{String: artist, Valid: artist != ""})
	if err != nil {
		return nil, fmt.Errorf("failed to list products by artist: %w", err)
	}
	return prod, nil
}

func (cfg *apiConfig) ListProductsByStatus(ctx context.Context, status string) ([]database.Product, error) {
	prod, err := cfg.DbQueries.ListProductsByStatus(ctx, status)
	if err != nil {
		return nil, fmt.Errorf("failed to list products by status: %w", err)
	}
	return prod, nil
}

func (cfg *apiConfig) UpdateProduct(ctx context.Context, id uuid.UUID, input CreateOrUpdateProductInput) (*database.Product, error) {
	params := database.UpdateProductParams{
		ID:             id,
		Type:           input.Type,
		Sku:            input.Sku,
		Upc:            input.Upc,
		Status:         input.Status,
		Cost:           input.Cost,
		Price:          input.Price,
		Envelope:       sql.NullString{String: input.Envelope, Valid: input.Envelope != ""},
		Artist:         sql.NullString{String: input.Artist, Valid: input.Artist != ""},
		Category:       sql.NullString{String: input.Category, Valid: input.Category != ""},
		ReleaseDate:    sql.NullTime{Time: input.ReleaseDate, Valid: !input.ReleaseDate.IsZero()},
		LastBoughtDate: sql.NullTime{Time: input.LastBoughtDate, Valid: !input.LastBoughtDate.IsZero()},
		Description:    sql.NullString{String: input.Description, Valid: input.Description != ""},
		TextFront:      sql.NullString{String: input.TextFront, Valid: input.TextFront != ""},
		TextInside:     sql.NullString{String: input.TextInside, Valid: input.TextInside != ""},
	}
	prod, err := cfg.DbQueries.UpdateProduct(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}
	return &prod, nil
}

func (cfg *apiConfig) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	return cfg.DbQueries.DeleteProduct(ctx, id)
}
