package service

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
)

func TestGetPlanogram_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	now := time.Now()
	mock.ExpectQuery(`SELECT id, name, num_pockets, notes, created_at, updated_at FROM planograms WHERE id = \$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "num_pockets", "notes", "created_at", "updated_at"}).
			AddRow(1, "Test Planogram", 5, sql.NullString{String: "", Valid: false}, now, now))
	planogram, err := GetPlanogram(ctx, dbQueries, 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if planogram == nil || planogram.ID != 1 {
		t.Errorf("expected planogram ID 1, got %+v", planogram)
	}
}

func TestGetPlanogram_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	mock.ExpectQuery(`SELECT id, name, num_pockets, notes, created_at, updated_at FROM planograms WHERE id = \$1`).
		WithArgs(2).
		WillReturnError(sql.ErrNoRows)
	planogram, err := GetPlanogram(ctx, dbQueries, 2)
	if err == nil || planogram != nil {
		t.Errorf("expected error for not found, got %v, %+v", err, planogram)
	}
}

func TestListPlanograms_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	now := time.Now()
	mock.ExpectQuery(`SELECT id, name, num_pockets, notes, created_at, updated_at FROM planograms ORDER BY created_at DESC`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "num_pockets", "notes", "created_at", "updated_at"}).
			AddRow(1, "Test Planogram", 5, sql.NullString{String: "", Valid: false}, now, now))
	planograms, err := ListPlanograms(ctx, dbQueries)
	if err != nil || len(planograms) != 1 {
		t.Errorf("expected 1 planogram, got %v, %d", err, len(planograms))
	}
}

func TestListPlanograms_Failure(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	mock.ExpectQuery(`SELECT id, name, num_pockets, notes, created_at, updated_at FROM planograms ORDER BY created_at DESC`).
		WillReturnError(sql.ErrConnDone)
	planograms, err := ListPlanograms(ctx, dbQueries)
	if err == nil || planograms != nil {
		t.Errorf("expected error, got %v, %+v", err, planograms)
	}
}

func TestCreatePlanogram_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	now := time.Now()
	input := CreatePlanogramInput{Name: "Test Planogram", NumPockets: 5, Notes: "Some notes"}
	mock.ExpectQuery(`INSERT INTO planograms \(name, num_pockets, notes\) VALUES \(\$1, \$2, \$3\) RETURNING id, name, num_pockets, notes, created_at, updated_at`).
		WithArgs(input.Name, input.NumPockets, sql.NullString{String: input.Notes, Valid: input.Notes != ""}).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "num_pockets", "notes", "created_at", "updated_at"}).
			AddRow(1, input.Name, input.NumPockets, sql.NullString{String: input.Notes, Valid: input.Notes != ""}, now, now))
	planogram, err := CreatePlanogram(ctx, dbQueries, input)
	if err != nil || planogram == nil || planogram.Name != input.Name {
		t.Errorf("expected success, got %v, %+v", err, planogram)
	}
}

func TestCreatePlanogram_InvalidInput(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := context.Background()
	input := CreatePlanogramInput{Name: "", NumPockets: 0, Notes: ""}
	planogram, err := CreatePlanogram(ctx, dbQueries, input)
	if err == nil || planogram != nil {
		t.Errorf("expected error for invalid input, got %v, %+v", err, planogram)
	}
}

func TestUpdatePlanogram_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	now := time.Now()
	input := UpdatePlanogramInput{ID: 1, Name: "Updated", NumPockets: 10, Notes: "Updated notes"}
	mock.ExpectQuery(`UPDATE planograms SET name = \$2, num_pockets = \$3, notes = \$4, updated_at = NOW\(\) WHERE id = \$1 RETURNING id, name, num_pockets, notes, created_at, updated_at`).
		WithArgs(input.ID, input.Name, input.NumPockets, sql.NullString{String: input.Notes, Valid: input.Notes != ""}).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "num_pockets", "notes", "created_at", "updated_at"}).
			AddRow(input.ID, input.Name, input.NumPockets, sql.NullString{String: input.Notes, Valid: input.Notes != ""}, now, now))
	planogram, err := UpdatePlanogram(ctx, dbQueries, input)
	if err != nil || planogram == nil || planogram.Name != input.Name {
		t.Errorf("expected success, got %v, %+v", err, planogram)
	}
}

func TestUpdatePlanogram_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	input := UpdatePlanogramInput{ID: 2, Name: "Updated", NumPockets: 10, Notes: "Updated notes"}
	mock.ExpectQuery(`UPDATE planograms SET name = \$2, num_pockets = \$3, notes = \$4, updated_at = NOW\(\) WHERE id = \$1 RETURNING id, name, num_pockets, notes, created_at, updated_at`).
		WithArgs(input.ID, input.Name, input.NumPockets, sql.NullString{String: input.Notes, Valid: input.Notes != ""}).
		WillReturnError(sql.ErrNoRows)
	planogram, err := UpdatePlanogram(ctx, dbQueries, input)
	if err == nil || planogram != nil {
		t.Errorf("expected error for not found, got %v, %+v", err, planogram)
	}
}

func TestDeletePlanogram_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	mock.ExpectExec(`DELETE FROM planograms WHERE id = \$1`).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	err := DeletePlanogram(ctx, dbQueries, 1)
	if err != nil {
		t.Errorf("expected success, got %v", err)
	}
}

func TestDeletePlanogram_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	mock.ExpectExec(`DELETE FROM planograms WHERE id = \$1`).WithArgs(2).WillReturnError(sql.ErrNoRows)
	err := DeletePlanogram(ctx, dbQueries, 2)
	if err == nil {
		t.Errorf("expected error for not found, got %v", err)
	}
}

func TestAssignPlanogramToLocation_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	planogramID := int32(1)
	customerID := int32(2)
	mock.ExpectQuery(`INSERT INTO planogram_customer_locations \(planogram_id, customer_location_id\) VALUES \(\$1, \$2\) ON CONFLICT \(planogram_id, customer_location_id\) DO NOTHING RETURNING id, planogram_id, customer_location_id`).
		WithArgs(planogramID, customerID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "planogram_id", "customer_location_id"}).AddRow(1, planogramID, customerID))
	pcl, err := AssignPlanogramToLocation(ctx, dbQueries, planogramID, customerID)
	if err != nil || pcl == nil || pcl.PlanogramID != planogramID {
		t.Errorf("expected success, got %v, %+v", err, pcl)
	}
}

func TestAssignPlanogramToLocation_InvalidInput(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := context.Background()
	pcl, err := AssignPlanogramToLocation(ctx, dbQueries, 0, 0)
	if err == nil || pcl != nil {
		t.Errorf("expected error for invalid input, got %v, %+v", err, pcl)
	}
}

func TestRemovePlanogramFromLocation_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	planogramID := int32(1)
	customerID := int32(2)
	mock.ExpectExec(`DELETE FROM planogram_customer_locations WHERE planogram_id = \$1 AND customer_location_id = \$2`).
		WithArgs(planogramID, customerID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := RemovePlanogramFromLocation(ctx, dbQueries, planogramID, customerID)
	if err != nil {
		t.Errorf("expected success, got %v", err)
	}
}

func TestRemovePlanogramFromLocation_InvalidInput(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := context.Background()
	err := RemovePlanogramFromLocation(ctx, dbQueries, 0, 0)
	if err == nil {
		t.Errorf("expected error for invalid input, got %v", err)
	}
}

func TestGetPlanogramsByLocation_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	now := time.Now()
	mock.ExpectQuery(`SELECT p.id, p.name, p.num_pockets, p.notes, p.created_at, p.updated_at FROM planograms p JOIN planogram_customer_locations pcl ON p.id = pcl.planogram_id WHERE pcl.customer_location_id = \$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "num_pockets", "notes", "created_at", "updated_at"}).AddRow(1, "Test Planogram", 5, sql.NullString{String: "", Valid: false}, now, now))
	planogram, err := GetPlanogramsByLocation(ctx, dbQueries, 1)
	if err != nil || planogram == nil || planogram.ID != 1 {
		t.Errorf("expected success, got %v, %+v", err, planogram)
	}
}

func TestGetPlanogramsByLocation_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	mock.ExpectQuery(`SELECT p.id, p.name, p.num_pockets, p.notes, p.created_at, p.updated_at FROM planograms p JOIN planogram_customer_locations pcl ON p.id = pcl.planogram_id WHERE pcl.customer_location_id = \$1`).
		WithArgs(2).
		WillReturnError(sql.ErrNoRows)
	planogram, err := GetPlanogramsByLocation(ctx, dbQueries, 2)
	if err == nil || planogram != nil {
		t.Errorf("expected error for not found, got %v, %+v", err, planogram)
	}
}

func TestListLocationsByPlanogram_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	now := time.Now()
	mock.ExpectQuery(`SELECT cl.id, cl.customer_id, cl.address_1, cl.address_2, cl.city, cl.state, cl.zip_code, cl.phone, cl.notes, cl.created_at, cl.updated_at FROM customer_locations cl JOIN planogram_customer_locations pcl ON cl.id = pcl.customer_location_id WHERE pcl.planogram_id = \$1 ORDER BY cl.id`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "customer_id", "address_1", "address_2", "city", "state", "zip_code", "phone", "notes", "created_at", "updated_at"}).
			AddRow(1, 1, "123 St", sql.NullString{String: "", Valid: false}, "City", "State", "12345", sql.NullString{String: "", Valid: false}, sql.NullString{String: "", Valid: false}, now, now))
	locations, err := ListLocationsByPlanogram(ctx, dbQueries, 1)
	if err != nil || len(locations) != 1 {
		t.Errorf("expected 1 location, got %v, %d", err, len(locations))
	}
}

func TestListLocationsByPlanogram_InvalidInput(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := context.Background()
	locations, err := ListLocationsByPlanogram(ctx, dbQueries, 0)
	if err == nil || locations != nil {
		t.Errorf("expected error for invalid input, got %v, %+v", err, locations)
	}
}

func TestListPocketsForPlanogram_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	mock.ExpectQuery(`SELECT id, planogram_id, pocket_number, category, product_id FROM planogram_pockets WHERE planogram_id = \$1 ORDER BY pocket_number`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "planogram_id", "pocket_number", "category", "product_id"}).
			AddRow(1, 1, 1, "A", sql.NullInt32{Int32: 0, Valid: false}))
	pockets, err := ListPocketsForPlanogram(ctx, dbQueries, 1)
	if err != nil || len(pockets) != 1 {
		t.Errorf("expected 1 pocket, got %v, %d", err, len(pockets))
	}
}

func TestListPocketsForPlanogram_InvalidInput(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := context.Background()
	pockets, err := ListPocketsForPlanogram(ctx, dbQueries, 0)
	if err == nil || pockets != nil {
		t.Errorf("expected error for invalid input, got %v, %+v", err, pockets)
	}
}

func TestGetPlanogramPocket_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	mock.ExpectQuery(`SELECT id, planogram_id, pocket_number, category, product_id FROM planogram_pockets WHERE id = \$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "planogram_id", "pocket_number", "category", "product_id"}).AddRow(1, 1, 1, "A", sql.NullInt32{Int32: 0, Valid: false}))
	pocket, err := GetPlanogramPocket(ctx, dbQueries, 1)
	if err != nil || pocket == nil || pocket.ID != 1 {
		t.Errorf("expected success, got %v, %+v", err, pocket)
	}
}

func TestGetPlanogramPocket_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	mock.ExpectQuery(`SELECT id, planogram_id, pocket_number, category, product_id FROM planogram_pockets WHERE id = \$1`).
		WithArgs(2).
		WillReturnError(sql.ErrNoRows)
	pocket, err := GetPlanogramPocket(ctx, dbQueries, 2)
	if err == nil || pocket != nil {
		t.Errorf("expected error for not found, got %v, %+v", err, pocket)
	}
}

func TestCreatePlanogramPocket_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	input := CreatePlanogramPocketInput{PlanogramID: 1, PocketNumber: 1, Category: "A", ProductID: 5}
	mock.ExpectQuery(`INSERT INTO planogram_pockets \(planogram_id, pocket_number, category, product_id\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING id, planogram_id, pocket_number, category, product_id`).
		WithArgs(input.PlanogramID, input.PocketNumber, input.Category, sql.NullInt32{Int32: input.ProductID, Valid: input.ProductID > 0}).
		WillReturnRows(sqlmock.NewRows([]string{"id", "planogram_id", "pocket_number", "category", "product_id"}).AddRow(1, 1, 1, "A", sql.NullInt32{Int32: input.ProductID, Valid: input.ProductID > 0}))
	pocket, err := CreatePlanogramPocket(ctx, dbQueries, input)
	if err != nil || pocket == nil || pocket.ID != 1 {
		t.Errorf("expected success, got %v, %+v", err, pocket)
	}
}

func TestCreatePlanogramPocket_InvalidInput(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := context.Background()
	input := CreatePlanogramPocketInput{PlanogramID: 0, PocketNumber: 0, Category: "", ProductID: 0}
	pocket, err := CreatePlanogramPocket(ctx, dbQueries, input)
	if err == nil || pocket != nil {
		t.Errorf("expected error for invalid input, got %v, %+v", err, pocket)
	}
}

func TestUpdatePlanogramPocket_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	input := UpdatePlanogramPocketInput{ID: 1, Category: "B", ProductID: 7}
	mock.ExpectQuery(`UPDATE planogram_pockets SET category = \$2, product_id = \$3 WHERE id = \$1 RETURNING id, planogram_id, pocket_number, category, product_id`).
		WithArgs(input.ID, input.Category, sql.NullInt32{Int32: input.ProductID, Valid: input.ProductID > 0}).
		WillReturnRows(sqlmock.NewRows([]string{"id", "planogram_id", "pocket_number", "category", "product_id"}).AddRow(1, 1, 1, "B", sql.NullInt32{Int32: input.ProductID, Valid: input.ProductID > 0}))
	pocket, err := UpdatePlanogramPocket(ctx, dbQueries, input)
	if err != nil || pocket == nil || pocket.ID != 1 {
		t.Errorf("expected success, got %v, %+v", err, pocket)
	}
}

func TestUpdatePlanogramPocket_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	input := UpdatePlanogramPocketInput{ID: 2, Category: "B", ProductID: 7}
	mock.ExpectQuery(`UPDATE planogram_pockets SET category = \$2, product_id = \$3 WHERE id = \$1 RETURNING id, planogram_id, pocket_number, category, product_id`).
		WithArgs(input.ID, input.Category, sql.NullInt32{Int32: input.ProductID, Valid: input.ProductID > 0}).
		WillReturnError(sql.ErrNoRows)
	pocket, err := UpdatePlanogramPocket(ctx, dbQueries, input)
	if err == nil || pocket != nil {
		t.Errorf("expected error for not found, got %v, %+v", err, pocket)
	}
}

func TestDeletePlanogramPocket_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	mock.ExpectExec(`DELETE FROM planogram_pockets WHERE id = \$1`).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	err := DeletePlanogramPocket(ctx, dbQueries, 1)
	if err != nil {
		t.Errorf("expected success, got %v", err)
	}
}

func TestDeletePlanogramPocket_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	mock.ExpectExec(`DELETE FROM planogram_pockets WHERE id = \$1`).WithArgs(2).WillReturnError(sql.ErrNoRows)
	err := DeletePlanogramPocket(ctx, dbQueries, 2)
	if err == nil {
		t.Errorf("expected error for not found, got %v", err)
	}
}

func TestGetPlanogramPocketByNumber_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	input := database.GetPlanogramPocketByNumberParams{PlanogramID: 1, PocketNumber: 1}
	mock.ExpectQuery(`SELECT id, planogram_id, pocket_number, category, product_id FROM planogram_pockets WHERE planogram_id = \$1 AND pocket_number = \$2`).
		WithArgs(input.PlanogramID, input.PocketNumber).
		WillReturnRows(sqlmock.NewRows([]string{"id", "planogram_id", "pocket_number", "category", "product_id"}).AddRow(1, 1, 1, "A", sql.NullInt32{Int32: 0, Valid: false}))
	pocket, err := GetPlanogramPocketByNumber(ctx, dbQueries, input.PlanogramID, input.PocketNumber)
	if err != nil || pocket == nil || pocket.ID != 1 {
		t.Errorf("expected success, got %v, %+v", err, pocket)
	}
}

func TestGetPlanogramPocketByNumber_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := context.Background()
	input := database.GetPlanogramPocketByNumberParams{PlanogramID: 2, PocketNumber: 2}
	mock.ExpectQuery(`SELECT id, planogram_id, pocket_number, category, product_id FROM planogram_pockets WHERE planogram_id = \$1 AND pocket_number = \$2`).
		WithArgs(input.PlanogramID, input.PocketNumber).
		WillReturnError(sql.ErrNoRows)
	pocket, err := GetPlanogramPocketByNumber(ctx, dbQueries, input.PlanogramID, input.PocketNumber)
	if err == nil || pocket != nil {
		t.Errorf("expected error for not found, got %v, %+v", err, pocket)
	}
}
