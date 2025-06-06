package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
)

func CreateOrder(ctx context.Context, db *database.Queries, input CreateOrderInput) (*database.Order, error) {
	if input.CustomerID == 0 {
		return nil, errors.New("customer_id is required")
	}
	if input.Status == "" {
		return nil, errors.New("status is required")
	}
	if input.Type == "" {
		return nil, errors.New("type is required")
	}
	params := database.CreateOrderParams{
		CustomerID:         input.CustomerID,
		CustomerLocationID: input.CustomerLocationID,
		OrderDate:          input.OrderDate,
		Status:             input.Status,
		Type:               input.Type,
		Method:             sql.NullString{String: input.Method, Valid: input.Method != ""},
		ShipDate:           input.ShipDate,
		PoNumber:           sql.NullString{String: input.PoNumber, Valid: input.PoNumber != ""},
		ShippingCost:       input.ShippingCost,
		FreeShipping:       input.FreeShipping,
		ApplyToCommission:  input.ApplyToCommission,
		SalesRep:           sql.NullString{String: input.SalesRep, Valid: input.SalesRep != ""},
		Notes:              sql.NullString{String: input.Notes, Valid: input.Notes != ""},
	}
	order, err := db.CreateOrder(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}
	return &order, nil
}

func GetOrder(ctx context.Context, db *database.Queries, id int32) (*database.Order, error) {
	order, err := db.GetOrder(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order not found")
		}
		return nil, fmt.Errorf("failed to get order: %w", err)
	}
	return &order, nil
}

func UpdateOrder(ctx context.Context, db *database.Queries, input UpdateOrderInput) (*database.Order, error) {
	if input.CustomerID == 0 {
		return nil, errors.New("customer_id is required")
	}
	if input.Status == "" {
		return nil, errors.New("status is required")
	}
	if input.Type == "" {
		return nil, errors.New("type is required")
	}
	params := database.UpdateOrderParams{
		ID:                 input.ID,
		CustomerID:         input.CustomerID,
		CustomerLocationID: input.CustomerLocationID,
		OrderDate:          input.OrderDate,
		Status:             input.Status,
		Type:               input.Type,
		Method:             sql.NullString{String: input.Method, Valid: input.Method != ""},
		ShipDate:           input.ShipDate,
		PoNumber:           sql.NullString{String: input.PoNumber, Valid: input.PoNumber != ""},
		ShippingCost:       input.ShippingCost,
		FreeShipping:       input.FreeShipping,
		ApplyToCommission:  input.ApplyToCommission,
		SalesRep:           sql.NullString{String: input.SalesRep, Valid: input.SalesRep != ""},
		Notes:              sql.NullString{String: input.Notes, Valid: input.Notes != ""},
	}
	order, err := db.UpdateOrder(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order not found")
		}
		return nil, fmt.Errorf("failed to update order: %w", err)
	}
	return &order, nil
}

func DeleteOrder(ctx context.Context, db *database.Queries, id int32) error {
	err := db.DeleteOrder(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("order not found")
		}
		return fmt.Errorf("failed to delete order: %w", err)
	}
	return nil
}

func ListOrdersOpen(ctx context.Context, db *database.Queries) ([]database.Order, error) {
	orders, err := db.ListOrdersOpen(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list open orders: %w", err)
	}
	return orders, nil
}

func ListOrdersByCustomer(ctx context.Context, db *database.Queries, customerID int32) ([]database.Order, error) {
	if customerID == 0 {
		return nil, errors.New("customer_id is required")
	}
	orders, err := db.ListOrdersByCustomer(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to list orders by customer: %w", err)
	}
	return orders, nil
}
