package service

import (
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

// UpdateUserPasswordInput holds the details for changing a user's password.
type UpdateUserPasswordInput struct {
	UserID      uuid.UUID
	OldPassword string
	NewPassword string
}
