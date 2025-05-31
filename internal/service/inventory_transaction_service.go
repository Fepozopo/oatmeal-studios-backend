package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
	"github.com/google/uuid"
)

// GetCurrentInventory returns the current inventory for a product
func GetCurrentInventory(ctx context.Context, db *database.Queries, productID uuid.UUID) (int64, error) {
	inv, err := db.GetCurrentInventory(ctx, productID)
	if err != nil {
		return 0, fmt.Errorf("failed to get current inventory: %w", err)
	}
	switch v := inv.(type) {
	case int64:
		return v, nil
	case int32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	default:
		return 0, fmt.Errorf("unexpected type for inventory: %T", v)
	}
}

// GetAllCurrentInventory returns the current inventory for all products
func GetAllCurrentInventory(ctx context.Context, db *database.Queries) ([]database.GetAllCurrentInventoryRow, error) {
	invs, err := db.GetAllCurrentInventory(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all current inventory: %w", err)
	}
	return invs, nil
}

// ListInventoryTransactions returns all inventory transactions for a product
func ListInventoryTransactions(ctx context.Context, db *database.Queries, productID uuid.UUID) ([]database.InventoryTransaction, error) {
	transactions, err := db.ListInventoryTransactions(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to list inventory transactions: %w", err)
	}
	return transactions, nil
}

// InsertInventoryTransaction inserts a new inventory transaction
func InsertInventoryTransaction(ctx context.Context, db *database.Queries, input InsertInventoryTransactionInput) (*database.InventoryTransaction, error) {
	if input.ProductID == uuid.Nil {
		return nil, fmt.Errorf("product ID cannot be empty")
	}
	if input.Change == 0 {
		return nil, fmt.Errorf("change cannot be zero")
	}
	if input.Reason == "" {
		return nil, fmt.Errorf("reason cannot be empty")
	}

	params := database.InsertInventoryTransactionParams{
		ProductID: input.ProductID,
		Change:    input.Change,
		Reason:    input.Reason,
		Notes:     sql.NullString{String: input.Notes, Valid: input.Notes != ""},
	}

	transaction, err := db.InsertInventoryTransaction(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to insert inventory transaction: %w", err)
	}

	return &transaction, nil
}

// GetInventoryChangesByDay returns inventory changes for a product grouped by day
func GetInventoryChangesByDay(ctx context.Context, db *database.Queries, productID uuid.UUID) ([]database.GetInventoryChangesByDayRow, error) {
	rows, err := db.GetInventoryChangesByDay(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get inventory changes by day: %w", err)
	}
	return rows, nil
}
