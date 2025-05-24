package service

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateOrderItem_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	input := CreateOrderItemInput{
		OrderID:      1,
		Sku:          "SKU123",
		Quantity:     10,
		Price:        2.5,
		Discount:     0.0,
		ItemTotal:    25.0,
		PocketNumber: 3,
	}
	mock.ExpectQuery(`-- name: CreateOrderItem :one`).
		WithArgs(
			input.OrderID,
			input.Sku,
			input.Quantity,
			input.Price,
			input.Discount,
			input.ItemTotal,
			input.PocketNumber,
		).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "order_id", "sku", "quantity", "price", "discount", "item_total", "pocket_number",
		}).AddRow(
			1, input.OrderID, input.Sku, input.Quantity, input.Price, input.Discount, input.ItemTotal, input.PocketNumber,
		))
	item, err := CreateOrderItem(ctx, dbQueries, input)
	if err != nil {
		t.Errorf("CreateOrderItem returned error: %v", err)
	}
	if item == nil || item.OrderID != input.OrderID {
		t.Errorf("CreateOrderItem returned wrong item: %+v", item)
	}
}

func TestCreateOrderItem_InvalidInput(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := newTestContext()
	input := CreateOrderItemInput{}
	item, err := CreateOrderItem(ctx, dbQueries, input)
	if err == nil {
		t.Errorf("CreateOrderItem should have returned an error for missing required fields")
	}
	if item != nil {
		t.Errorf("CreateOrderItem should have returned nil for invalid input")
	}
}

func TestGetOrderItem_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	mock.ExpectQuery(`SELECT id, order_id, sku, quantity, price, discount, item_total, pocket_number FROM order_items WHERE id = \$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "order_id", "sku", "quantity", "price", "discount", "item_total", "pocket_number",
		}).AddRow(
			1, 1, "SKU123", 10, 2.5, 0.0, 25.0, 3,
		))
	item, err := GetOrderItem(ctx, dbQueries, 1)
	if err != nil {
		t.Errorf("GetOrderItem returned error: %v", err)
	}
	if item == nil || item.ID != 1 {
		t.Errorf("GetOrderItem returned wrong item: %+v", item)
	}
}

func TestGetOrderItem_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	mock.ExpectQuery(`SELECT id, order_id, sku, quantity, price, discount, item_total, pocket_number FROM order_items WHERE id = \$1`).
		WithArgs(999).
		WillReturnError(sql.ErrNoRows)
	item, err := GetOrderItem(ctx, dbQueries, 999)
	if err == nil {
		t.Errorf("GetOrderItem should have returned an error for not found item")
	}
	if err != nil && err.Error() != "order item not found" {
		t.Errorf("GetOrderItem should have returned 'order item not found', got: %v", err)
	}
	if item != nil {
		t.Errorf("GetOrderItem should have returned nil for not found")
	}
}

func TestListOrderItemsBySKU_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	sku := "SKU123"
	mock.ExpectQuery(`SELECT id, order_id, sku, quantity, price, discount, item_total, pocket_number FROM order_items WHERE sku = \$1`).
		WithArgs(sku).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "order_id", "sku", "quantity", "price", "discount", "item_total", "pocket_number",
		}).AddRow(
			1, 1, sku, 10, 2.5, 0.0, 25.0, 3,
		))
	items, err := ListOrderItemsBySKU(ctx, dbQueries, sku)
	if err != nil {
		t.Errorf("ListOrderItemsBySKU returned error: %v", err)
	}
	if len(items) != 1 {
		t.Errorf("expected 1 item, got %d", len(items))
	}
}

func TestListOrderItemsBySKU_InvalidInput(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := newTestContext()
	items, err := ListOrderItemsBySKU(ctx, dbQueries, "")
	if err == nil {
		t.Errorf("ListOrderItemsBySKU should have returned an error for missing sku")
	}
	if items != nil {
		t.Errorf("ListOrderItemsBySKU should have returned nil for invalid input")
	}
}

func TestUpdateOrderItem_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	input := UpdateOrderItemInput{
		Sku:          "SKU123",
		Quantity:     20,
		Price:        3.0,
		Discount:     0.1,
		ItemTotal:    54.0,
		PocketNumber: 4,
	}
	mock.ExpectQuery(`-- name: UpdateOrderItem :one`).
		WithArgs(
			1,
			input.Sku,
			input.Quantity,
			input.Price,
			input.Discount,
			input.ItemTotal,
			input.PocketNumber,
		).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "order_id", "sku", "quantity", "price", "discount", "item_total", "pocket_number",
		}).AddRow(
			1, 1, input.Sku, input.Quantity, input.Price, input.Discount, input.ItemTotal, input.PocketNumber,
		))
	item, err := UpdateOrderItem(ctx, dbQueries, 1, input)
	if err != nil {
		t.Errorf("UpdateOrderItem returned error: %v", err)
	}
	if item == nil || item.ID != 1 {
		t.Errorf("UpdateOrderItem returned wrong item: %+v", item)
	}
}

func TestUpdateOrderItem_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	input := UpdateOrderItemInput{
		Sku:          "SKU123",
		Quantity:     20,
		Price:        3.0,
		Discount:     0.1,
		ItemTotal:    54.0,
		PocketNumber: 4,
	}
	mock.ExpectQuery(`-- name: UpdateOrderItem :one`).
		WithArgs(999, input.Sku, input.Quantity, input.Price, input.Discount, input.ItemTotal, input.PocketNumber).
		WillReturnError(sql.ErrNoRows)
	item, err := UpdateOrderItem(ctx, dbQueries, 999, input)
	if err == nil {
		t.Errorf("UpdateOrderItem should have returned an error for not found item")
	}
	if err != nil && err.Error() != "order item not found" {
		t.Errorf("UpdateOrderItem should have returned 'order item not found', got: %v", err)
	}
	if item != nil {
		t.Errorf("UpdateOrderItem should have returned nil for not found")
	}
}

func TestDeleteOrderItem_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	mock.ExpectExec(`-- name: DeleteOrderItem :exec`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := DeleteOrderItem(ctx, dbQueries, 1)
	if err != nil {
		t.Errorf("DeleteOrderItem returned error: %v", err)
	}
}

func TestDeleteOrderItem_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	mock.ExpectExec(`-- name: DeleteOrderItem :exec`).
		WithArgs(999).
		WillReturnError(sql.ErrNoRows)
	err := DeleteOrderItem(ctx, dbQueries, 999)
	if err == nil {
		t.Errorf("DeleteOrderItem should have returned an error for not found item")
	}
	if err != nil && err.Error() != "order item not found" {
		t.Errorf("DeleteOrderItem should have returned 'order item not found', got: %v", err)
	}
}
