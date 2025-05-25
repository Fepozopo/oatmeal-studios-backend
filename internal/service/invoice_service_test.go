package service

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUpdateInvoice_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	input := CreateOrUpdateInvoiceInput{
		InvoiceDate:        now,
		OrderID:            1,
		CustomerID:         2,
		CustomerLocationID: 3,
		DueDate:            now.Add(30 * 24 * time.Hour),
		Status:             "Paid",
		Total:              200.0,
	}
	mock.ExpectQuery(`-- name: UpdateInvoice :one`).
		WithArgs(
			1,
			input.InvoiceDate,
			input.OrderID,
			input.CustomerID,
			input.CustomerLocationID,
			input.DueDate,
			input.Status,
			input.Total,
		).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "invoice_date", "order_id", "customer_id", "customer_location_id", "due_date", "status", "total",
		}).AddRow(
			1, now, now, now, 1, 2, 3, now.Add(30*24*time.Hour), "Paid", 200.0,
		))
	invoice, err := UpdateInvoice(ctx, dbQueries, 1, input)
	if err != nil {
		t.Errorf("UpdateInvoice returned error: %v", err)
	}
	if invoice == nil {
		t.Errorf("UpdateInvoice should have returned a non-nil invoice")
	}
	if invoice != nil && invoice.Status != "Paid" {
		t.Errorf("expected status 'Paid', got %s", invoice.Status)
	}
}

func TestDeleteInvoice_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	mock.ExpectExec(`-- name: DeleteInvoice :exec`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := DeleteInvoice(ctx, dbQueries, 1)
	if err != nil {
		t.Errorf("DeleteInvoice returned error: %v", err)
	}
}

func TestListInvoicesByCustomer_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	mock.ExpectQuery(`-- name: ListInvoicesByCustomer :many`).
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "invoice_date", "order_id", "customer_id", "customer_location_id", "due_date", "status", "total",
		}).AddRow(
			1, now, now, now, 1, 2, 3, now.Add(30*24*time.Hour), "Open", 100.0,
		))
	invoices, err := ListInvoicesByCustomer(ctx, dbQueries, 2)
	if err != nil {
		t.Errorf("ListInvoicesByCustomer returned error: %v", err)
	}
	if len(invoices) != 1 {
		t.Errorf("expected 1 invoice, got %d", len(invoices))
	}
}

func TestListInvoicesByCustomerLocation_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	mock.ExpectQuery(`-- name: ListInvoicesByCustomerLocation :many`).
		WithArgs(3).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "invoice_date", "order_id", "customer_id", "customer_location_id", "due_date", "status", "total",
		}).AddRow(
			1, now, now, now, 1, 2, 3, now.Add(30*24*time.Hour), "Open", 100.0,
		))
	invoices, err := ListInvoicesByCustomerLocation(ctx, dbQueries, 3)
	if err != nil {
		t.Errorf("ListInvoicesByCustomerLocation returned error: %v", err)
	}
	if len(invoices) != 1 {
		t.Errorf("expected 1 invoice, got %d", len(invoices))
	}
}
