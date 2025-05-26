package auth

import (
	"fmt"
	"net"
	"net/mail"
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// IsValidEmail performs a more robust check for a valid email address.
func IsValidEmail(email string) error {
	// Basic format check
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("invalid email format: %w", err)
	}

	// Extract domain
	parts := strings.Split(addr.Address, "@")
	if len(parts) != 2 {
		return fmt.Errorf("invalid email address structure")
	}
	domain := parts[1]

	// Check MX records
	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return fmt.Errorf("email domain does not have MX records")
	}

	return nil
}

func IsValidPassword(password string) error {
	if password == "" {
		return fmt.Errorf("password cannot be empty")
	}

	// Forbidden characters
	forbidden := " \"'\\/ "
	for _, c := range password {
		for _, f := range forbidden {
			if c == f {
				return fmt.Errorf("password contains forbidden character: '%c'", c)
			}
		}
	}

	// Check for at least one uppercase letter
	if matched, _ := regexp.MatchString(`[A-Z]`, password); !matched {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}

	// Check for at least one lowercase letter
	if matched, _ := regexp.MatchString(`[a-z]`, password); !matched {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}

	// Check for at least one digit
	if matched, _ := regexp.MatchString(`[0-9]`, password); !matched {
		return fmt.Errorf("password must contain at least one digit")
	}

	// Check for at least one special character
	if matched, _ := regexp.MatchString(`[!@#$%^&*(),.?":{}|<>]`, password); !matched {
		return fmt.Errorf("password must contain at least one special character")
	}

	return nil
}

// IsValidPhone checks if the phone number is in a valid format.
func IsValidPhone(phone string) error {
	// Basic format check using regex
	// Require area code: E.164 format with country code and at least 10 digits total (e.g., +1234567890)
	phoneRegex := `^\+?[1-9]\d{9,14}$`
	matched, err := regexp.MatchString(phoneRegex, phone)
	if err != nil {
		return fmt.Errorf("error validating phone number: %w", err)
	}
	if !matched {
		return fmt.Errorf("invalid phone number format")
	}
	return nil
}

// ValidateJWT takes a JWT token as a string and a tokenSecret as a string,
// validates the token with the given secret, and returns the UUID in the
// Subject field of the token claims. If the token is invalid or the Subject
// field is not a valid UUID, it returns an error.
func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("expected *jwt.RegisteredClaims, got %T", token.Claims)
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("expected valid uuid in Subject field, got %q", claims.Subject)
	}

	return userID, nil
}
