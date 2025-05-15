package service

import (
	"context"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/auth"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
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
