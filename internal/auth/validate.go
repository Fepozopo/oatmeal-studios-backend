package auth

import (
	"fmt"
	"net"
	"net/mail"
	"regexp"
	"strings"
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
