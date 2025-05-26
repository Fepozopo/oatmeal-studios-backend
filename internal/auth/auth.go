package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a password string and returns a hashed string of the
// password and an error. The hashed string is a string of bytes generated
// by the bcrypt.GenerateFromPassword() function. The error is returned if
// there is an error generating the hash.
func HashPassword(password string) (string, error) {
	err := IsValidPassword(password)
	if err != nil {
		return "", fmt.Errorf("invalid password: %w", err)
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashBytes), nil
}

// CheckPasswordHash checks if the given password matches the given hash.
// It returns an error if the password does not match the hash.
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// MakeJWT creates a JWT token containing the given userID as the subject and
// signs it with the given tokenSecret. The token is set to expire after the
// given expiresIn duration. It returns the JWT token as a string and an
// error if there is an error generating the token.
func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "Oatmeal Studios",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn).UTC()),
		Subject:   userID.String(),
	})

	return token.SignedString([]byte(tokenSecret))
}

// GetBearerToken extracts the Bearer token from the Authorization header
// of the provided http.Header. It returns the token as a string and an
// error if the header does not exist or the token is invalid.
func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header does not exist")
	}
	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if tokenString == "" {
		return "", fmt.Errorf("authorization header is invalid")
	}

	return tokenString, nil
}

// MakeRefreshToken generates a new refresh token as a hexadecimal string.
// The token is created using 32 bytes of cryptographically secure random data.
// It returns the generated token and an error if there is a failure in generating the random data.
func MakeRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate random data: %w", err)
	}
	return hex.EncodeToString(b), nil
}

// GetAPIKey extracts the API key from the Authorization header
// of the provided http.Header. It returns the key as a string and an
// error if the header does not exist or the key is invalid.
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header does not exist")
	}
	apiKey := strings.TrimSpace(strings.TrimPrefix(authHeader, "ApiKey "))
	if apiKey == "" {
		return "", fmt.Errorf("authorization header is invalid")
	}
	return apiKey, nil
}
