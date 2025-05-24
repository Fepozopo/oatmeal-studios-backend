package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/auth"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
	"github.com/google/uuid"
)

// RegisterUserInput holds the registration details for a new user.
type RegisterUserInput struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// UpdateUserNameInput holds the details for updating a user's name.
type UpdateUserNameInput struct {
	UserID    uuid.UUID `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

// UpdateUserPasswordInput holds the details for changing a user's password.
type UpdateUserPasswordInput struct {
	UserID      uuid.UUID `json:"user_id"`
	OldPassword string    `json:"old_password"`
	NewPassword string    `json:"new_password"`
}

// RegisterUser registers a new user after validating input and hashing the password.
func RegisterUser(ctx context.Context, db *database.Queries, input RegisterUserInput) (*database.User, error) {
	// Validate input fields
	if input.Email == "" || input.Password == "" || input.FirstName == "" || input.LastName == "" {
		return nil, errors.New("all fields are required")
	}
	if err := auth.IsValidEmail(input.Email); err != nil {
		return nil, fmt.Errorf("invalid email format: %w", err)
	}

	// Hash the password (includes strength validation)
	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Prepare params for DB insert
	params := database.CreateUserParams{
		Email:     input.Email,
		Password:  hashedPassword,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	user, err := db.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}

// AuthenticateUser checks the user's credentials and returns the user if valid.
func AuthenticateUser(ctx context.Context, db *database.Queries, email, password string) (*database.User, error) {
	// Validate input
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	// Fetch user by email
	user, err := db.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Verify password
	if err := auth.CheckPasswordHash(password, user.Password); err != nil {
		return nil, fmt.Errorf("invalid password: %w", err)
	}

	return &user, nil
}

// UpdateUser updates the user's profile information after validating input.
func UpdateUserName(ctx context.Context, db *database.Queries, input UpdateUserNameInput) (*database.User, error) {
	// Validate input
	if input.FirstName == "" || input.LastName == "" {
		return nil, errors.New("first name and last name are required")
	}

	// Update user profile
	params := database.UpdateUserNameParams{
		ID:        input.UserID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err := db.UpdateUserName(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update user profile: %w", err)
	}

	// Fetch updated user
	user, err := db.GetUserByID(ctx, input.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated user: %w", err)
	}
	return &user, nil
}

func UpdateUserPassword(ctx context.Context, db *database.Queries, input UpdateUserPasswordInput) error {
	// Validate input
	if input.OldPassword == "" || input.NewPassword == "" {
		return fmt.Errorf("old password and new password are required")
	}

	// Fetch user by ID
	user, err := db.GetUserByID(ctx, input.UserID)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	// Verify old password
	if err := auth.CheckPasswordHash(input.OldPassword, user.Password); err != nil {
		return fmt.Errorf("invalid old password: %w", err)
	}

	// Hash the new password
	hashedNewPassword, err := auth.HashPassword(input.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to hash new password: %w", err)
	}

	// Update user's password
	params := database.UpdateUserPasswordParams{
		ID:       input.UserID,
		Password: hashedNewPassword,
	}

	if err := db.UpdateUserPassword(ctx, params); err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}

// DeleteUser deletes a user by their ID.
func DeleteUser(ctx context.Context, db *database.Queries, userID uuid.UUID) error {
	if userID == uuid.Nil {
		return errors.New("user ID is required")
	}
	if err := db.DeleteUser(ctx, userID); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// DeleteAllUsers deletes all users from the database.
func DeleteAllUsers(ctx context.Context, db *database.Queries) error {
	if err := db.DeleteAllUsers(ctx); err != nil {
		return fmt.Errorf("failed to delete all users: %w", err)
	}
	return nil
}
