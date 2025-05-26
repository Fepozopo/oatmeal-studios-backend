package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
	"github.com/google/uuid"
)

func CreateProduct(ctx context.Context, db *database.Queries, input CreateProductInput) (*database.Product, error) {
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

	prod, err := db.CreateProduct(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}
	return &prod, nil
}

func GetProductByID(ctx context.Context, db *database.Queries, id uuid.UUID) (*database.Product, error) {
	prod, err := db.GetProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get product by ID: %w", err)
	}
	return &prod, nil
}

func GetProductBySKU(ctx context.Context, db *database.Queries, sku string) (*database.Product, error) {
	prod, err := db.GetProductBySKU(ctx, sku)
	if err != nil {
		return nil, fmt.Errorf("failed to get product by SKU: %w", err)
	}
	return &prod, nil
}

func ListProductsByType(ctx context.Context, db *database.Queries, typ string) ([]database.Product, error) {
	prod, err := db.ListProductsByType(ctx, typ)
	if err != nil {
		return nil, fmt.Errorf("failed to list products by type: %w", err)
	}
	return prod, nil
}

func ListProductsByCategory(ctx context.Context, db *database.Queries, category string) ([]database.Product, error) {
	prod, err := db.ListProductsByCategory(ctx, sql.NullString{String: category, Valid: category != ""})
	if err != nil {
		return nil, fmt.Errorf("failed to list products by category: %w", err)
	}
	return prod, nil
}

func ListProductsByArtist(ctx context.Context, db *database.Queries, artist string) ([]database.Product, error) {
	prod, err := db.ListProductsByArtist(ctx, sql.NullString{String: artist, Valid: artist != ""})
	if err != nil {
		return nil, fmt.Errorf("failed to list products by artist: %w", err)
	}
	return prod, nil
}

func ListProductsByStatus(ctx context.Context, db *database.Queries, status string) ([]database.Product, error) {
	prod, err := db.ListProductsByStatus(ctx, status)
	if err != nil {
		return nil, fmt.Errorf("failed to list products by status: %w", err)
	}
	return prod, nil
}

func UpdateProduct(ctx context.Context, db *database.Queries, input UpdateProductInput) (*database.Product, error) {
	params := database.UpdateProductParams{
		ID:             input.ID,
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
	prod, err := db.UpdateProduct(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}
	return &prod, nil
}

func DeleteProduct(ctx context.Context, db *database.Queries, id uuid.UUID) error {
	return db.DeleteProduct(ctx, id)
}
