package service

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/auth"
	"github.com/google/uuid"
)

// --- RegisterUser ---
func TestRegisterUser_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	defer func() {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	}()

	input := RegisterUserInput{
		Email:     "test@example.com",
		Password:  "Password1!",
		FirstName: "John",
		LastName:  "Doe",
	}

	hashed, err := auth.HashPassword(input.Password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO users`)).
		WithArgs(input.Email, sqlmock.AnyArg(), input.FirstName, input.LastName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "first_name", "last_name", "password"}).
			AddRow(uuid.New(), time.Now(), time.Now(), input.Email, input.FirstName, input.LastName, hashed))

	user, err := RegisterUser(ctx, dbQueries, input)
	if err != nil {
		t.Errorf("RegisterUser returned an error: %v", err)
	}
	if user == nil {
		t.Errorf("RegisterUser should have returned a non-nil user")
	}
	if user != nil && user.Email != input.Email || user.FirstName != input.FirstName || user.LastName != input.LastName {
		t.Errorf("expected user email '%s', first name '%s', last name '%s', got email '%s', first name '%s', last name '%s'",
			input.Email, input.FirstName, input.LastName, user.Email, user.FirstName, user.LastName)
	}
}

func TestRegisterUser_InvalidInput(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := newTestContext()

	input := RegisterUserInput{}
	user, err := RegisterUser(ctx, dbQueries, input)
	if err == nil {
		t.Errorf("RegisterUser should have returned an error for invalid input")
	}
	if user != nil {
		t.Errorf("RegisterUser should have returned a nil user for invalid input, got: %v", user)
	}
}

// --- AuthenticateUser ---
func TestAuthenticateUser_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	defer func() {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	}()

	email := "test@example.com"
	password := "Password1!"

	hashed, err := auth.HashPassword(password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, created_at, updated_at, email, first_name, last_name, password FROM users WHERE email = $1`)).
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "first_name", "last_name", "password"}).
			AddRow(uuid.New(), time.Now(), time.Now(), email, "John", "Doe", hashed))

	user, err := AuthenticateUser(ctx, dbQueries, email, password)
	if err != nil {
		t.Errorf("AuthenticateUser returned an error: %v", err)
	}
	if user == nil {
		t.Errorf("AuthenticateUser should have returned a non-nil user")
	}
}

func TestAuthenticateUser_InvalidPassword(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	defer func() {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	}()

	email := "test@example.com"
	password := "WrongPassword1!"

	hashed, err := auth.HashPassword("NotThePassword1!")
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, created_at, updated_at, email, first_name, last_name, password FROM users WHERE email = $1`)).
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "first_name", "last_name", "password"}).
			AddRow(uuid.New(), time.Now(), time.Now(), email, "John", "Doe", hashed))

	user, err := AuthenticateUser(ctx, dbQueries, email, password)
	if err == nil {
		t.Errorf("AuthenticateUser should have returned an error for invalid password")
	}
	if user != nil {
		t.Errorf("AuthenticateUser should have returned a nil user for invalid password, got: %v", user)
	}
}

// --- UpdateUser ---
func TestUpdateUserName_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	defer func() {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	}()

	input := UpdateUserNameInput{
		UserID:    uuid.New(),
		FirstName: "Jane",
		LastName:  "Smith",
	}

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE users`)).
		WithArgs(input.UserID, input.FirstName, input.LastName).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, created_at, updated_at, email, first_name, last_name, password FROM users WHERE id = $1`)).
		WithArgs(input.UserID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "first_name", "last_name", "password"}).
			AddRow(input.UserID, time.Now(), time.Now(), "test@example.com", input.FirstName, input.LastName, "hashedpassword"))

	user, err := UpdateUserName(ctx, dbQueries, input)
	if err != nil {
		t.Errorf("UpdateUserName returned an error: %v", err)
	}
	if user == nil {
		t.Errorf("UpdateUserName should have returned a non-nil user")
	}
	if user.FirstName != input.FirstName || user.LastName != input.LastName {
		t.Errorf("expected user first name '%s', last name '%s', got first name '%s', last name '%s'",
			input.FirstName, input.LastName, user.FirstName, user.LastName)
	}
	if user.ID != input.UserID {
		t.Errorf("expected user ID '%s', got '%s'", input.UserID, user.ID)
	}
}

func TestUpdateUser_InvalidInput(t *testing.T) {
	dbQueries, _ := newTestDB(t)
	ctx := newTestContext()

	input := UpdateUserNameInput{UserID: uuid.New()}
	_, err := UpdateUserName(ctx, dbQueries, input)
	if err == nil {
		t.Errorf("UpdateUserName should have returned an error for invalid input")
	}
	if err != nil && err.Error() != "first name and last name are required" {
		t.Errorf("expected error 'first name and last name are required', got '%s'", err.Error())
	}
}

// --- UpdateUserPassword ---
func TestUpdateUserPassword_Success(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	defer func() {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	}()

	userID := uuid.New()
	oldPassword := "OldPassword1!"
	newPassword := "NewPassword1!"

	hashedOld, err := auth.HashPassword(oldPassword)
	if err != nil {
		t.Fatalf("failed to hash old password: %v", err)
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, created_at, updated_at, email, first_name, last_name, password FROM users WHERE id = $1`)).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "first_name", "last_name", "password"}).
			AddRow(userID, time.Now(), time.Now(), "user@example.com", "John", "Doe", hashedOld))

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET password = $2, updated_at = NOW() WHERE id = $1`)).
		WithArgs(userID, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = UpdateUserPassword(ctx, dbQueries, UpdateUserPasswordInput{
		UserID:      userID,
		OldPassword: oldPassword,
		NewPassword: newPassword,
	})
	if err != nil {
		t.Errorf("UpdateUserPassword returned an error: %v", err)
	}
}

func TestUpdateUserPassword_InvalidOldPassword(t *testing.T) {
	dbQueries, mock := newTestDB(t)
	ctx := newTestContext()
	defer func() {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	}()

	userID := uuid.New()
	oldPassword := "WrongOldPassword1!"

	hashedOld, err := auth.HashPassword("NotThePassword1!")
	if err != nil {
		t.Fatalf("failed to hash old password: %v", err)
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, created_at, updated_at, email, first_name, last_name, password FROM users WHERE id = $1`)).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "first_name", "last_name", "password"}).
			AddRow(userID, time.Now(), time.Now(), "user@example.com", "John", "Doe", hashedOld))

	err = UpdateUserPassword(ctx, dbQueries, UpdateUserPasswordInput{
		UserID:      userID,
		OldPassword: oldPassword,
		NewPassword: "irrelevant",
	})
	if err == nil {
		t.Errorf("UpdateUserPassword should have returned an error for invalid old password")
	}
	expectedError := "invalid old password: crypto/bcrypt: hashedPassword is not the hash of the given password"
	if err != nil && err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%s'", expectedError, err.Error())
	}
}
