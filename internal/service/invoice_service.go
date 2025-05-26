package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
)

func (cfg *apiConfig) CreateInvoice(ctx context.Context, input CreateInvoiceInput) (*database.Invoice, error) {
	if input.OrderID == 0 {
		return nil, errors.New("order_id is required")
	}
	if input.CustomerID == 0 {
		return nil, errors.New("customer_id is required")
	}
	if input.Status == "" {
		return nil, errors.New("status is required")
	}
	params := database.CreateInvoiceParams{
		InvoiceDate:        input.InvoiceDate,
		OrderID:            input.OrderID,
		CustomerID:         input.CustomerID,
		CustomerLocationID: input.CustomerLocationID,
		DueDate:            input.DueDate,
		Status:             input.Status,
		Total:              input.Total,
	}
	invoice, err := cfg.DbQueries.CreateInvoice(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create invoice: %w", err)
	}
	return &invoice, nil
}

func (cfg *apiConfig) GetInvoice(ctx context.Context, id int32) (*database.Invoice, error) {
	invoice, err := cfg.DbQueries.GetInvoice(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invoice not found")
		}
		return nil, fmt.Errorf("failed to get invoice: %w", err)
	}
	return &invoice, nil
}

func (cfg *apiConfig) GetInvoicesByOrder(ctx context.Context, orderID int32) (*database.Invoice, error) {
	invoice, err := cfg.DbQueries.GetInvoicesByOrder(ctx, orderID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invoice not found for order")
		}
		return nil, fmt.Errorf("failed to get invoice by order: %w", err)
	}
	return &invoice, nil
}

func (cfg *apiConfig) ListInvoicesByCustomer(ctx context.Context, customerID int32) ([]database.Invoice, error) {
	if customerID == 0 {
		return nil, errors.New("customer_id is required")
	}
	invoices, err := cfg.DbQueries.ListInvoicesByCustomer(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to list invoices by customer: %w", err)
	}
	return invoices, nil
}

func (cfg *apiConfig) ListInvoicesByCustomerLocation(ctx context.Context, customerLocationID int32) ([]database.Invoice, error) {
	if customerLocationID == 0 {
		return nil, errors.New("customer_location_id is required")
	}
	invoices, err := cfg.DbQueries.ListInvoicesByCustomerLocation(ctx, customerLocationID)
	if err != nil {
		return nil, fmt.Errorf("failed to list invoices by customer location: %w", err)
	}
	return invoices, nil
}

func (cfg *apiConfig) UpdateInvoice(ctx context.Context, input UpdateInvoiceInput) (*database.Invoice, error) {
	if input.ID == 0 {
		return nil, errors.New("id is required")
	}
	if input.OrderID == 0 {
		return nil, errors.New("order_id is required")
	}
	if input.CustomerID == 0 {
		return nil, errors.New("customer_id is required")
	}
	if input.Status == "" {
		return nil, errors.New("status is required")
	}
	params := database.UpdateInvoiceParams{
		ID:                 input.ID,
		InvoiceDate:        input.InvoiceDate,
		OrderID:            input.OrderID,
		CustomerID:         input.CustomerID,
		CustomerLocationID: input.CustomerLocationID,
		DueDate:            input.DueDate,
		Status:             input.Status,
		Total:              input.Total,
	}
	invoice, err := cfg.DbQueries.UpdateInvoice(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invoice not found")
		}
		return nil, fmt.Errorf("failed to update invoice: %w", err)
	}
	return &invoice, nil
}

func (cfg *apiConfig) DeleteInvoice(ctx context.Context, id int32) error {
	err := cfg.DbQueries.DeleteInvoice(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invoice not found")
		}
		return fmt.Errorf("failed to delete invoice: %w", err)
	}
	return nil
}
