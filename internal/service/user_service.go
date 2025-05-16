package service

import (
	"context"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/auth"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
	"github.com/google/uuid"
)

// RegisterUserInput holds the registration details for a new user.
type RegisterUserInput struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

// RegisterUserResponse holds the result of a registration attempt.
type RegisterUserResponse struct {
	Success bool
	User    *database.User
	Error   string
}

// UpdateUserInput holds the details for updating a user's profile.
type UpdateUserInput struct {
	UserID    uuid.UUID
	FirstName string
	LastName  string
}

// RegisterUser registers a new user after validating input and hashing the password.
func RegisterUser(ctx context.Context, db *database.Queries, input RegisterUserInput) RegisterUserResponse {
	// Validate input fields
	if input.Email == "" || input.Password == "" || input.FirstName == "" || input.LastName == "" {
		return RegisterUserResponse{Success: false, Error: "all fields are required"}
	}
	if err := auth.IsValidEmail(input.Email); err != nil {
		return RegisterUserResponse{Success: false, Error: err.Error()}
	}

	// Hash the password (includes strength validation)
	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		return RegisterUserResponse{Success: false, Error: err.Error()}
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
		return RegisterUserResponse{Success: false, Error: fmt.Sprintf("failed to create user: %v", err)}
	}

	return RegisterUserResponse{Success: true, User: &user}
}

// AuthenticateUser checks the user's credentials and returns the user if valid.
func AuthenticateUser(ctx context.Context, db *database.Queries, email, password string) (*database.User, error) {
	// Validate input
	if email == "" || password == "" {
		return nil, fmt.Errorf("email and password are required")
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

func UpdateUser(ctx context.Context, db *database.Queries, input UpdateUserInput) (*database.User, error) {
	// Validate input
	if input.FirstName == "" || input.LastName == "" {
		return nil, fmt.Errorf("first name and last name are required")
	}

	// Update user profile
	params := database.UpdateUserParams{
		ID:        input.UserID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	user, err := db.UpdateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update user profile: %w", err)
	}

	return &user, nil
}
