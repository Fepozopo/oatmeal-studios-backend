package service

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/auth"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// --- RegisterUser ---
func TestRegisterUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	dbQueries := database.New(db)
	ctx := context.Background()

	input := RegisterUserInput{
		Email:     "test@example.com",
		Password:  "Password1!",
		FirstName: "John",
		LastName:  "Doe",
	}

	// Use real password hashing and email validation
	hashed, err := auth.HashPassword(input.Password)
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO users`)).
		WithArgs(input.Email, sqlmock.AnyArg(), input.FirstName, input.LastName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "first_name", "last_name", "password"}).
			AddRow(uuid.New(), time.Now(), time.Now(), input.Email, input.FirstName, input.LastName, hashed))

	user, err := RegisterUser(ctx, dbQueries, input)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, input.Email, user.Email)
}

func TestRegisterUser_InvalidInput(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()
	dbQueries := database.New(db)
	ctx := context.Background()

	input := RegisterUserInput{}
	user, err := RegisterUser(ctx, dbQueries, input)
	assert.Error(t, err)
	assert.Nil(t, user)
}

// --- AuthenticateUser ---
func TestAuthenticateUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	dbQueries := database.New(db)
	ctx := context.Background()

	email := "test@example.com"
	password := "Password1!"

	hashed, err := auth.HashPassword(password)
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, created_at, updated_at, email, first_name, last_name, password FROM users WHERE email = $1`)).
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "first_name", "last_name", "password"}).
			AddRow(uuid.New(), time.Now(), time.Now(), email, "John", "Doe", hashed))

	user, err := AuthenticateUser(ctx, dbQueries, email, password)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
}

func TestAuthenticateUser_InvalidPassword(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	dbQueries := database.New(db)
	ctx := context.Background()

	email := "test@example.com"
	password := "WrongPassword1!"

	// Hash a different password than the one provided
	hashed, err := auth.HashPassword("NotThePassword1!")
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, created_at, updated_at, email, first_name, last_name, password FROM users WHERE email = $1`)).
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "first_name", "last_name", "password"}).
			AddRow(uuid.New(), time.Now(), time.Now(), email, "John", "Doe", hashed))

	user, err := AuthenticateUser(ctx, dbQueries, email, password)
	assert.Error(t, err)
	assert.Nil(t, user)
}

// --- UpdateUser ---
func TestUpdateUserName_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	dbQueries := database.New(db)
	ctx := context.Background()

	input := UpdateUserNameInput{
		UserID:    uuid.New(),
		FirstName: "Jane",
		LastName:  "Smith",
	}

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE users`)).
		WithArgs(input.UserID, input.FirstName, input.LastName).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = UpdateUserName(ctx, dbQueries, input)
	assert.NoError(t, err)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateUser_InvalidInput(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()
	dbQueries := database.New(db)
	ctx := context.Background()

	input := UpdateUserNameInput{UserID: uuid.New()}
	err := UpdateUserName(ctx, dbQueries, input)
	assert.Error(t, err)
	assert.Equal(t, "first name and last name are required", err.Error())
}

// --- UpdateUserPassword ---
func TestUpdateUserPassword_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	dbQueries := database.New(db)
	ctx := context.Background()

	userID := uuid.New()
	oldPassword := "OldPassword1!"
	newPassword := "NewPassword1!"

	hashedOld, err := auth.HashPassword(oldPassword)
	assert.NoError(t, err)

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
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUpdateUserPassword_InvalidOldPassword(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	dbQueries := database.New(db)
	ctx := context.Background()

	userID := uuid.New()
	oldPassword := "WrongOldPassword1!"

	// Hash a different password than the one provided
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
	assert.Error(t, err)
	assert.Equal(t, "invalid old password: crypto/bcrypt: hashedPassword is not the hash of the given password", err.Error())
}
