package service

import (
	"context"
	"fmt"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
	"github.com/google/uuid"
)

// CreateRefreshToken creates a new refresh token.
func CreateRefreshToken(ctx context.Context, db *database.Queries, token string, id uuid.UUID) error {
	if token == "" || id == uuid.Nil {
		return fmt.Errorf("invalid parameters: token and user ID must not be empty")
	}
	params := database.CreateRefreshTokenParams{
		Token:  token,
		UserID: id,
	}

	err := db.CreateRefreshToken(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to create refresh token: %w", err)
	}
	return nil
}

// RevokeRefreshToken revokes a refresh token by its token string.
func RevokeRefreshToken(ctx context.Context, db *database.Queries, token string) error {
	if token == "" {
		return fmt.Errorf("token must not be empty")
	}

	err := db.RevokeRefreshToken(ctx, token)
	if err != nil {
		return fmt.Errorf("failed to revoke refresh token: %w", err)
	}
	return nil
}

// GetRefreshToken retrieves a refresh token by its token string.
func GetRefreshToken(ctx context.Context, db *database.Queries, token string) (*database.RefreshToken, error) {
	if token == "" {
		return nil, fmt.Errorf("token must not be empty")
	}

	// Retrieve the refresh token from the database
	rt, err := db.GetRefreshToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("failed to get refresh token: %w", err)
	}
	return &rt, nil
}
