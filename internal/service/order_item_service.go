package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
)

func (cfg *apiConfig) GetOrderItem(ctx context.Context, id int32) (*database.OrderItem, error) {
	item, err := cfg.DbQueries.GetOrderItem(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order item not found")
		}
		return nil, fmt.Errorf("failed to get order item: %w", err)
	}
	return &item, nil
}

func (cfg *apiConfig) ListOrderItemsBySKU(ctx context.Context, sku string) ([]database.OrderItem, error) {
	if sku == "" {
		return nil, errors.New("sku is required")
	}
	items, err := cfg.DbQueries.ListOrderItemsBySKU(ctx, sku)
	if err != nil {
		return nil, fmt.Errorf("failed to list order items by sku: %w", err)
	}
	return items, nil
}

func (cfg *apiConfig) CreateOrderItem(ctx context.Context, input CreateOrderItemInput) (*database.OrderItem, error) {
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
	item, err := cfg.DbQueries.CreateOrderItem(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create order item: %w", err)
	}
	return &item, nil
}

func (cfg *apiConfig) UpdateOrderItem(ctx context.Context, input UpdateOrderItemInput) (*database.OrderItem, error) {
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
	item, err := cfg.DbQueries.UpdateOrderItem(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order item not found")
		}
		return nil, fmt.Errorf("failed to update order item: %w", err)
	}
	return &item, nil
}

func (cfg *apiConfig) DeleteOrderItem(ctx context.Context, id int32) error {
	if id == 0 {
		return errors.New("id is required")
	}
	err := cfg.DbQueries.DeleteOrderItem(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("order item not found")
		}
		return fmt.Errorf("failed to delete order item: %w", err)
	}
	return nil
}
