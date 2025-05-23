package service

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

// --- CreateProduct ---
func TestCreateProduct_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	input := CreateOrUpdateProductInput{
		Type:   "card",
		Sku:    "SKU123",
		Upc:    "UPC123",
		Status: "active",
		Cost:   1.23,
		Price:  2.34,
	}
	id := uuid.New()
	mock.ExpectQuery(`-- name: CreateProduct :one`).
		WithArgs(input.Type, input.Sku, input.Upc, input.Status, input.Cost, input.Price, sql.NullString{Valid: false}, sql.NullString{Valid: false}, sql.NullString{Valid: false}, sql.NullTime{Valid: false}, sql.NullTime{Valid: false}, sql.NullString{Valid: false}, sql.NullString{Valid: false}, sql.NullString{Valid: false}).
		WillReturnRows(makeProductRow(id, input, now))
	prod, err := CreateProduct(ctx, dbQueries, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if prod.Sku != input.Sku {
		t.Errorf("expected sku %s, got %s", input.Sku, prod.Sku)
	}
}

func TestCreateProduct_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	input := CreateOrUpdateProductInput{}
	_, err := CreateProduct(ctx, dbQueries, input)
	if err == nil {
		t.Fatal("expected error for missing required fields")
	}
	// Simulate DB error
	input = CreateOrUpdateProductInput{Type: "card", Sku: "SKU123", Upc: "UPC123", Status: "active", Cost: 1, Price: 2}
	mock.ExpectQuery(`-- name: CreateProduct :one`).
		WithArgs(input.Type, input.Sku, input.Upc, input.Status, input.Cost, input.Price, sql.NullString{Valid: false}, sql.NullString{Valid: false}, sql.NullString{Valid: false}, sql.NullTime{Valid: false}, sql.NullTime{Valid: false}, sql.NullString{Valid: false}, sql.NullString{Valid: false}, sql.NullString{Valid: false}).
		WillReturnError(errors.New("db error"))
	_, err = CreateProduct(ctx, dbQueries, input)
	if err == nil {
		t.Fatal("expected db error")
	}
}

// --- GetProductByID ---
func TestGetProductByID_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	id := uuid.New()
	input := CreateOrUpdateProductInput{Type: "card", Sku: "SKU123", Upc: "UPC123", Status: "active", Cost: 1, Price: 2}
	mock.ExpectQuery(`-- name: GetProductByID :one`).WithArgs(id).WillReturnRows(makeProductRow(id, input, now))
	prod, err := GetProductByID(ctx, dbQueries, id)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if prod.ID != id {
		t.Errorf("expected id %v, got %v", id, prod.ID)
	}
}

func TestGetProductByID_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	id := uuid.New()
	mock.ExpectQuery(`-- name: GetProductByID :one`).WithArgs(id).WillReturnError(sql.ErrNoRows)
	_, err := GetProductByID(ctx, dbQueries, id)
	if err == nil {
		t.Fatal("expected error for not found")
	}
}

// --- GetProductBySKU ---
func TestGetProductBySKU_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	id := uuid.New()
	input := CreateOrUpdateProductInput{Type: "card", Sku: "SKU123", Upc: "UPC123", Status: "active", Cost: 1, Price: 2}
	mock.ExpectQuery(`-- name: GetProductBySKU :one`).WithArgs(input.Sku).WillReturnRows(makeProductRow(id, input, now))
	prod, err := GetProductBySKU(ctx, dbQueries, input.Sku)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if prod.Sku != input.Sku {
		t.Errorf("expected sku %s, got %s", input.Sku, prod.Sku)
	}
}

func TestGetProductBySKU_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	sku := "NOTFOUND"
	mock.ExpectQuery(`-- name: GetProductBySKU :one`).WithArgs(sku).WillReturnError(sql.ErrNoRows)
	_, err := GetProductBySKU(ctx, dbQueries, sku)
	if err == nil {
		t.Fatal("expected error for not found")
	}
}

// --- ListProductsByType ---
func TestListProductsByType_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	input := CreateOrUpdateProductInput{Type: "card", Sku: "SKU123", Upc: "UPC123", Status: "active", Cost: 1, Price: 2}
	id := uuid.New()
	mock.ExpectQuery(`-- name: ListProductsByType :many`).WithArgs(input.Type).WillReturnRows(makeProductRow(id, input, now))
	prods, err := ListProductsByType(ctx, dbQueries, input.Type)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(prods) != 1 {
		t.Errorf("expected 1 product, got %d", len(prods))
	}
}

func TestListProductsByType_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	mock.ExpectQuery(`-- name: ListProductsByType :many`).WithArgs("badtype").WillReturnError(errors.New("db error"))
	_, err := ListProductsByType(ctx, dbQueries, "badtype")
	if err == nil {
		t.Fatal("expected db error")
	}
}

// --- ListProductsByCategory ---
func TestListProductsByCategory_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	input := CreateOrUpdateProductInput{Type: "card", Sku: "SKU123", Upc: "UPC123", Status: "active", Cost: 1, Price: 2, Category: "cat"}
	id := uuid.New()
	mock.ExpectQuery(`-- name: ListProductsByCategory :many`).WithArgs(sql.NullString{String: "cat", Valid: true}).WillReturnRows(makeProductRow(id, input, now))
	prods, err := ListProductsByCategory(ctx, dbQueries, "cat")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(prods) != 1 {
		t.Errorf("expected 1 product, got %d", len(prods))
	}
}

func TestListProductsByCategory_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	mock.ExpectQuery(`-- name: ListProductsByCategory :many`).WithArgs(sql.NullString{String: "badcat", Valid: true}).WillReturnError(errors.New("db error"))
	_, err := ListProductsByCategory(ctx, dbQueries, "badcat")
	if err == nil {
		t.Fatal("expected db error")
	}
}

// --- ListProductsByArtist ---
func TestListProductsByArtist_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	input := CreateOrUpdateProductInput{Type: "card", Sku: "SKU123", Upc: "UPC123", Status: "active", Cost: 1, Price: 2, Artist: "artist"}
	id := uuid.New()
	mock.ExpectQuery(`-- name: ListProductsByArtist :many`).WithArgs(sql.NullString{String: "artist", Valid: true}).WillReturnRows(makeProductRow(id, input, now))
	prods, err := ListProductsByArtist(ctx, dbQueries, "artist")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(prods) != 1 {
		t.Errorf("expected 1 product, got %d", len(prods))
	}
}

func TestListProductsByArtist_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	mock.ExpectQuery(`-- name: ListProductsByArtist :many`).WithArgs(sql.NullString{String: "badartist", Valid: true}).WillReturnError(errors.New("db error"))
	_, err := ListProductsByArtist(ctx, dbQueries, "badartist")
	if err == nil {
		t.Fatal("expected db error")
	}
}

// --- ListProductsByStatus ---
func TestListProductsByStatus_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	input := CreateOrUpdateProductInput{Type: "card", Sku: "SKU123", Upc: "UPC123", Status: "active", Cost: 1, Price: 2}
	id := uuid.New()
	mock.ExpectQuery(`-- name: ListProductsByStatus :many`).WithArgs(input.Status).WillReturnRows(makeProductRow(id, input, now))
	prods, err := ListProductsByStatus(ctx, dbQueries, input.Status)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(prods) != 1 {
		t.Errorf("expected 1 product, got %d", len(prods))
	}
}

func TestListProductsByStatus_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	mock.ExpectQuery(`-- name: ListProductsByStatus :many`).WithArgs("badstatus").WillReturnError(errors.New("db error"))
	_, err := ListProductsByStatus(ctx, dbQueries, "badstatus")
	if err == nil {
		t.Fatal("expected db error")
	}
}

// --- UpdateProduct ---
func TestUpdateProduct_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	now := time.Now()
	id := uuid.New()
	input := CreateOrUpdateProductInput{Type: "card", Sku: "SKU123", Upc: "UPC123", Status: "active", Cost: 1, Price: 2}
	mock.ExpectQuery(`-- name: UpdateProduct :one`).
		WithArgs(id, input.Type, input.Sku, input.Upc, input.Status, input.Cost, input.Price, sql.NullString{Valid: false}, sql.NullString{Valid: false}, sql.NullString{Valid: false}, sql.NullTime{Valid: false}, sql.NullTime{Valid: false}, sql.NullString{Valid: false}, sql.NullString{Valid: false}, sql.NullString{Valid: false}).
		WillReturnRows(makeProductRow(id, input, now))
	prod, err := UpdateProduct(ctx, dbQueries, id, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if prod.ID != id {
		t.Errorf("expected id %v, got %v", id, prod.ID)
	}
}

func TestUpdateProduct_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	id := uuid.New()
	input := CreateOrUpdateProductInput{Type: "card", Sku: "SKU123", Upc: "UPC123", Status: "active", Cost: 1, Price: 2}
	mock.ExpectQuery(`-- name: UpdateProduct :one`).
		WithArgs(id, input.Type, input.Sku, input.Upc, input.Status, input.Cost, input.Price, sql.NullString{Valid: false}, sql.NullString{Valid: false}, sql.NullString{Valid: false}, sql.NullTime{Valid: false}, sql.NullTime{Valid: false}, sql.NullString{Valid: false}, sql.NullString{Valid: false}, sql.NullString{Valid: false}).
		WillReturnError(errors.New("db error"))
	_, err := UpdateProduct(ctx, dbQueries, id, input)
	if err == nil {
		t.Fatal("expected db error")
	}
}

// --- DeleteProduct ---
func TestDeleteProduct_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	id := uuid.New()
	mock.ExpectExec(`-- name: DeleteProduct :exec`).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))
	err := DeleteProduct(ctx, dbQueries, id)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDeleteProduct_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	id := uuid.New()
	mock.ExpectExec(`-- name: DeleteProduct :exec`).WithArgs(id).WillReturnError(errors.New("db error"))
	err := DeleteProduct(ctx, dbQueries, id)
	if err == nil {
		t.Fatal("expected db error")
	}
}
