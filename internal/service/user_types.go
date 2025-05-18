package service

import (
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
