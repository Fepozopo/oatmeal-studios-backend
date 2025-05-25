package service

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func TestCreateRefreshToken_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	token := "test-refresh-token"
	userID := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO refresh_tokens`)).
		WithArgs(token, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := CreateRefreshToken(ctx, dbQueries, token, userID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateRefreshToken_InvalidParams(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := newTestContext()

	err := CreateRefreshToken(ctx, dbQueries, "", uuid.Nil)
	if err == nil {
		t.Errorf("expected error for invalid params, got nil")
	}
}

func TestRevokeRefreshToken_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	token := "test-refresh-token"

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE refresh_tokens`)).
		WithArgs(token).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := RevokeRefreshToken(ctx, dbQueries, token)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRevokeRefreshToken_EmptyToken(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := newTestContext()

	err := RevokeRefreshToken(ctx, dbQueries, "")
	if err == nil {
		t.Errorf("expected error for empty token, got nil")
	}
}

func TestGetRefreshToken_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	token := "test-refresh-token"
	now := time.Now()
	userID := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT token, created_at, updated_at, user_id, expires_at, revoked_at FROM refresh_tokens WHERE token = $1`)).
		WithArgs(token).
		WillReturnRows(sqlmock.NewRows([]string{"token", "created_at", "updated_at", "user_id", "expires_at", "revoked_at"}).
			AddRow(token, now, now, userID, now.Add(60*24*time.Hour), sql.NullTime{}))

	rt, err := GetRefreshToken(ctx, dbQueries, token)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if rt == nil || rt.Token != token || rt.UserID != userID {
		t.Errorf("unexpected refresh token result: %+v", rt)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRefreshToken_EmptyToken(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := newTestContext()

	rt, err := GetRefreshToken(ctx, dbQueries, "")
	if err == nil || rt != nil {
		t.Errorf("expected error and nil result for empty token")
	}
}

func TestGetRefreshToken_NotFound(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	token := "not-found-token"

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT token, created_at, updated_at, user_id, expires_at, revoked_at FROM refresh_tokens WHERE token = $1`)).
		WithArgs(token).
		WillReturnError(sql.ErrNoRows)

	rt, err := GetRefreshToken(ctx, dbQueries, token)
	if err == nil || rt != nil {
		t.Errorf("expected error and nil result for not found token")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
