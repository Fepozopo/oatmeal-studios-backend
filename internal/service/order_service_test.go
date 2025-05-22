package service

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateOrder_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()

	input := CreateOrUpdateOrderInput{
		CustomerID:         1,
		CustomerLocationID: 2,
		OrderDate:          now,
		Status:             "Open",
		Type:               "reorder",
		Method:             "online",
		ShipDate:           now,
		PoNumber:           "PO123",
		ShippingCost:       5.0,
		FreeShipping:       false,
		ApplyToCommission:  true,
		SalesRep:           "Rep1",
		Notes:              "Test order",
	}

	mock.ExpectQuery(`-- name: CreateOrder :one`).
		WithArgs(
			input.CustomerID,
			input.CustomerLocationID,
			input.OrderDate,
			input.Status,
			input.Type,
			input.Method,
			input.ShipDate,
			input.PoNumber,
			input.ShippingCost,
			input.FreeShipping,
			input.ApplyToCommission,
			input.SalesRep,
			input.Notes,
		).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "customer_id", "customer_location_id", "order_date", "ship_date", "status", "type", "method", "po_number", "shipping_cost", "free_shipping", "apply_to_commission", "sales_rep", "notes",
		}).AddRow(
			1, now, now, 1, 2, now, now, "Open", "reorder", "online", "PO123", 5.0, false, true, "Rep1", "Test order",
		))

	order, err := CreateOrder(ctx, dbQueries, input)
	if err != nil {
		t.Errorf("CreateOrder returned error: %v", err)
	}
	if order == nil {
		t.Errorf("CreateOrder should have returned a non-nil order")
	}
	if order != nil && order.CustomerID != input.CustomerID {
		t.Errorf("expected customer_id %d, got %d", input.CustomerID, order.CustomerID)
	}
}

func TestCreateOrder_InvalidInput(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := newTestContext()
	input := CreateOrUpdateOrderInput{}
	order, err := CreateOrder(ctx, dbQueries, input)
	if err == nil {
		t.Errorf("CreateOrder should have returned an error for missing required fields")
	}
	if order != nil {
		t.Errorf("CreateOrder should have returned nil for invalid input")
	}
}

func TestGetOrder_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	mock.ExpectQuery(`SELECT id, created_at, updated_at, customer_id, customer_location_id, order_date, ship_date, status, type, method, po_number, shipping_cost, free_shipping, apply_to_commission, sales_rep, notes FROM orders WHERE id = \$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "customer_id", "customer_location_id", "order_date", "ship_date", "status", "type", "method", "po_number", "shipping_cost", "free_shipping", "apply_to_commission", "sales_rep", "notes",
		}).AddRow(
			1, now, now, 1, 2, now, now, "Open", "reorder", "online", "PO123", 5.0, false, true, "Rep1", "Test order",
		))
	order, err := GetOrder(ctx, dbQueries, 1)
	if err != nil {
		t.Errorf("GetOrder returned error: %v", err)
	}
	if order == nil {
		t.Errorf("GetOrder should have returned a non-nil order")
	}
	if order != nil && order.ID != 1 {
		t.Errorf("expected order ID 1, got %d", order.ID)
	}
}

func TestGetOrder_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	mock.ExpectQuery(`SELECT id, created_at, updated_at, customer_id, customer_location_id, order_date, ship_date, status, type, method, po_number, shipping_cost, free_shipping, apply_to_commission, sales_rep, notes FROM orders WHERE id = \$1`).
		WithArgs(999).
		WillReturnError(sql.ErrNoRows)
	order, err := GetOrder(ctx, dbQueries, 999)
	if err == nil {
		t.Errorf("GetOrder should have returned an error for not found order")
	}
	if err != nil && err.Error() != "order not found" {
		t.Errorf("GetOrder should have returned 'order not found', got: %v", err)
	}
	if order != nil {
		t.Errorf("GetOrder should have returned nil for not found")
	}
}

func TestUpdateOrder_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	input := CreateOrUpdateOrderInput{
		CustomerID:         1,
		CustomerLocationID: 2,
		OrderDate:          now,
		Status:             "Open",
		Type:               "reorder",
		Method:             "online",
		ShipDate:           now,
		PoNumber:           "PO123",
		ShippingCost:       5.0,
		FreeShipping:       false,
		ApplyToCommission:  true,
		SalesRep:           "Rep1",
		Notes:              "Test order",
	}
	mock.ExpectQuery(`-- name: UpdateOrder :one`).
		WithArgs(
			1,
			input.CustomerID,
			input.CustomerLocationID,
			input.OrderDate,
			input.Status,
			input.Type,
			input.Method,
			input.ShipDate,
			input.PoNumber,
			input.ShippingCost,
			input.FreeShipping,
			input.ApplyToCommission,
			input.SalesRep,
			input.Notes,
		).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "customer_id", "customer_location_id", "order_date", "ship_date", "status", "type", "method", "po_number", "shipping_cost", "free_shipping", "apply_to_commission", "sales_rep", "notes",
		}).AddRow(
			1, now, now, 1, 2, now, now, "Open", "reorder", "online", "PO123", 5.0, false, true, "Rep1", "Test order",
		))
	order, err := UpdateOrder(ctx, dbQueries, 1, input)
	if err != nil {
		t.Errorf("UpdateOrder returned error: %v", err)
	}
	if order == nil {
		t.Errorf("UpdateOrder should have returned a non-nil order")
	}
	if order != nil && order.ID != 1 {
		t.Errorf("expected order ID 1, got %d", order.ID)
	}
}

func TestUpdateOrder_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	input := CreateOrUpdateOrderInput{
		CustomerID:         1,
		CustomerLocationID: 2,
		OrderDate:          now,
		Status:             "Open",
		Type:               "reorder",
		Method:             "online",
		ShipDate:           now,
		PoNumber:           "PO123",
		ShippingCost:       5.0,
		FreeShipping:       false,
		ApplyToCommission:  true,
		SalesRep:           "Rep1",
		Notes:              "Test order",
	}
	mock.ExpectQuery(`-- name: UpdateOrder :one`).
		WithArgs(999, input.CustomerID, input.CustomerLocationID, input.OrderDate, input.Status, input.Type, input.Method, input.ShipDate, input.PoNumber, input.ShippingCost, input.FreeShipping, input.ApplyToCommission, input.SalesRep, input.Notes).
		WillReturnError(sql.ErrNoRows)
	order, err := UpdateOrder(ctx, dbQueries, 999, input)
	if err == nil {
		t.Errorf("UpdateOrder should have returned an error for not found order")
	}
	if err != nil && err.Error() != "order not found" {
		t.Errorf("UpdateOrder should have returned 'order not found', got: %v", err)
	}
	if order != nil {
		t.Errorf("UpdateOrder should have returned nil for not found")
	}
}

func TestDeleteOrder_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	mock.ExpectExec(`-- name: DeleteOrder :exec`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := DeleteOrder(ctx, dbQueries, 1)
	if err != nil {
		t.Errorf("DeleteOrder returned error: %v", err)
	}
}

func TestDeleteOrder_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	mock.ExpectExec(`-- name: DeleteOrder :exec`).
		WithArgs(999).
		WillReturnError(sql.ErrNoRows)
	err := DeleteOrder(ctx, dbQueries, 999)
	if err == nil {
		t.Errorf("DeleteOrder should have returned an error for not found order")
	}
	if err != nil && err.Error() != "order not found" {
		t.Errorf("DeleteOrder should have returned 'order not found', got: %v", err)
	}
}

func TestListOrdersOpen_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	mock.ExpectQuery(`SELECT id, created_at, updated_at, customer_id, customer_location_id, order_date, ship_date, status, type, method, po_number, shipping_cost, free_shipping, apply_to_commission, sales_rep, notes FROM orders WHERE status = 'Open' ORDER BY order_date DESC`).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "customer_id", "customer_location_id", "order_date", "ship_date", "status", "type", "method", "po_number", "shipping_cost", "free_shipping", "apply_to_commission", "sales_rep", "notes",
		}).AddRow(
			1, now, now, 1, 2, now, now, "Open", "reorder", "online", "PO123", 5.0, false, true, "Rep1", "Test order",
		))
	orders, err := ListOrdersOpen(ctx, dbQueries)
	if err != nil {
		t.Errorf("ListOrdersOpen returned error: %v", err)
	}
	if len(orders) != 1 {
		t.Errorf("expected 1 open order, got %d", len(orders))
	}
}

func TestListOrdersOpen_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	mock.ExpectQuery(`SELECT id, created_at, updated_at, customer_id, customer_location_id, order_date, ship_date, status, type, method, po_number, shipping_cost, free_shipping, apply_to_commission, sales_rep, notes FROM orders WHERE status = 'Open' ORDER BY order_date DESC`).
		WillReturnError(errors.New("db error"))
	orders, err := ListOrdersOpen(ctx, dbQueries)
	if err == nil {
		t.Errorf("ListOrdersOpen should have returned an error")
	}
	if orders != nil {
		t.Errorf("ListOrdersOpen should have returned nil on error")
	}
}

func TestListOrdersByCustomer_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	mock.ExpectQuery(`SELECT id, created_at, updated_at, customer_id, customer_location_id, order_date, ship_date, status, type, method, po_number, shipping_cost, free_shipping, apply_to_commission, sales_rep, notes FROM orders WHERE customer_id = \$1 ORDER BY order_date DESC`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "customer_id", "customer_location_id", "order_date", "ship_date", "status", "type", "method", "po_number", "shipping_cost", "free_shipping", "apply_to_commission", "sales_rep", "notes",
		}).AddRow(
			1, now, now, 1, 2, now, now, "Open", "reorder", "online", "PO123", 5.0, false, true, "Rep1", "Test order",
		))
	orders, err := ListOrdersByCustomer(ctx, dbQueries, 1)
	if err != nil {
		t.Errorf("ListOrdersByCustomer returned error: %v", err)
	}
	if len(orders) != 1 {
		t.Errorf("expected 1 order, got %d", len(orders))
	}
}

func TestListOrdersByCustomer_InvalidInput(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := newTestContext()
	orders, err := ListOrdersByCustomer(ctx, dbQueries, 0)
	if err == nil {
		t.Errorf("ListOrdersByCustomer should have returned an error for missing customer_id")
	}
	if orders != nil {
		t.Errorf("ListOrdersByCustomer should have returned nil for invalid input")
	}
}
