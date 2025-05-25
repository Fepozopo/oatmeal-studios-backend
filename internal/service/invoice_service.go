package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
)

type CreateOrUpdateInvoiceInput struct {
	InvoiceDate        time.Time `json:"invoice_date"`
	OrderID            int32     `json:"order_id"`
	CustomerID         int32     `json:"customer_id"`
	CustomerLocationID int32     `json:"customer_location_id,omitempty"`
	DueDate            time.Time `json:"due_date"`
	Status             string    `json:"status"`
	Total              float64   `json:"total"`
}

func CreateInvoice(ctx context.Context, db *database.Queries, input CreateOrUpdateInvoiceInput) (*database.Invoice, error) {
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
	invoice, err := db.CreateInvoice(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create invoice: %w", err)
	}
	return &invoice, nil
}

func GetInvoice(ctx context.Context, db *database.Queries, id int32) (*database.Invoice, error) {
	invoice, err := db.GetInvoice(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invoice not found")
		}
		return nil, fmt.Errorf("failed to get invoice: %w", err)
	}
	return &invoice, nil
}

func GetInvoicesByOrder(ctx context.Context, db *database.Queries, orderID int32) (*database.Invoice, error) {
	invoice, err := db.GetInvoicesByOrder(ctx, orderID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invoice not found for order")
		}
		return nil, fmt.Errorf("failed to get invoice by order: %w", err)
	}
	return &invoice, nil
}

func ListInvoicesByCustomer(ctx context.Context, db *database.Queries, customerID int32) ([]database.Invoice, error) {
	if customerID == 0 {
		return nil, errors.New("customer_id is required")
	}
	invoices, err := db.ListInvoicesByCustomer(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to list invoices by customer: %w", err)
	}
	return invoices, nil
}

func ListInvoicesByCustomerLocation(ctx context.Context, db *database.Queries, customerLocationID int32) ([]database.Invoice, error) {
	if customerLocationID == 0 {
		return nil, errors.New("customer_location_id is required")
	}
	invoices, err := db.ListInvoicesByCustomerLocation(ctx, customerLocationID)
	if err != nil {
		return nil, fmt.Errorf("failed to list invoices by customer location: %w", err)
	}
	return invoices, nil
}

func UpdateInvoice(ctx context.Context, db *database.Queries, id int32, input CreateOrUpdateInvoiceInput) (*database.Invoice, error) {
	if id == 0 {
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
		ID:                 id,
		InvoiceDate:        input.InvoiceDate,
		OrderID:            input.OrderID,
		CustomerID:         input.CustomerID,
		CustomerLocationID: input.CustomerLocationID,
		DueDate:            input.DueDate,
		Status:             input.Status,
		Total:              input.Total,
	}
	invoice, err := db.UpdateInvoice(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invoice not found")
		}
		return nil, fmt.Errorf("failed to update invoice: %w", err)
	}
	return &invoice, nil
}

func DeleteInvoice(ctx context.Context, db *database.Queries, id int32) error {
	err := db.DeleteInvoice(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invoice not found")
		}
		return fmt.Errorf("failed to delete invoice: %w", err)
	}
	return nil
}
