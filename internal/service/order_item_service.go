package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
)

func GetOrderItem(ctx context.Context, db *database.Queries, id int32) (*database.OrderItem, error) {
	item, err := db.GetOrderItem(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order item not found")
		}
		return nil, fmt.Errorf("failed to get order item: %w", err)
	}
	return &item, nil
}

func ListOrderItemsBySKU(ctx context.Context, db *database.Queries, sku string) ([]database.OrderItem, error) {
	if sku == "" {
		return nil, errors.New("sku is required")
	}
	items, err := db.ListOrderItemsBySKU(ctx, sku)
	if err != nil {
		return nil, fmt.Errorf("failed to list order items by sku: %w", err)
	}
	return items, nil
}

func CreateOrderItem(ctx context.Context, db *database.Queries, input CreateOrderItemInput) (*database.OrderItem, error) {
	if input.OrderID == 0 {
		return nil, errors.New("order_id is required")
	}
	if input.Sku == "" {
		return nil, errors.New("sku is required")
	}
	params := database.CreateOrderItemParams{
		OrderID:      input.OrderID,
		Sku:          input.Sku,
		Quantity:     input.Quantity,
		Price:        input.Price,
		Discount:     input.Discount,
		ItemTotal:    input.ItemTotal,
		PocketNumber: sql.NullInt32{Int32: input.PocketNumber, Valid: input.PocketNumber != 0},
	}
	item, err := db.CreateOrderItem(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create order item: %w", err)
	}
	return &item, nil
}

func UpdateOrderItem(ctx context.Context, db *database.Queries, input UpdateOrderItemInput) (*database.OrderItem, error) {
	if input.ID == 0 {
		return nil, errors.New("id is required")
	}
	if input.Sku == "" {
		return nil, errors.New("sku is required")
	}
	params := database.UpdateOrderItemParams{
		ID:           input.ID,
		Sku:          input.Sku,
		Quantity:     input.Quantity,
		Price:        input.Price,
		Discount:     input.Discount,
		ItemTotal:    input.ItemTotal,
		PocketNumber: sql.NullInt32{Int32: input.PocketNumber, Valid: input.PocketNumber != 0},
	}
	item, err := db.UpdateOrderItem(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order item not found")
		}
		return nil, fmt.Errorf("failed to update order item: %w", err)
	}
	return &item, nil
}

func DeleteOrderItem(ctx context.Context, db *database.Queries, id int32) error {
	if id == 0 {
		return errors.New("id is required")
	}
	err := db.DeleteOrderItem(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("order item not found")
		}
		return fmt.Errorf("failed to delete order item: %w", err)
	}
	return nil
}
