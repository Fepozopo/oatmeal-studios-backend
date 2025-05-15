package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a password string and returns a hashed string of the
// password and an error. The hashed string is a string of bytes generated
// by the bcrypt.GenerateFromPassword() function. The error is returned if
// there is an error generating the hash.
func HashPassword(password string) (string, error) {
	err := ValidatePassword(password)
	if err != nil {
		return "", fmt.Errorf("invalid password: %w", err)
	}
	err = ValidatePasswordStrength(password)
	if err != nil {
		return "", fmt.Errorf("weak password: %w", err)
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

func ValidatePassword(password string) error {
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

	return nil
}

// ValidatePasswordStrength checks if the password meets the strength requirements.
// It returns an error if the password is too weak.
func ValidatePasswordStrength(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	if len(password) > 128 {
		return fmt.Errorf("password must be at most 128 characters long")
	}
	if !containsUppercase(password) {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !containsLowercase(password) {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !containsDigit(password) {
		return fmt.Errorf("password must contain at least one digit")
	}
	if !containsSpecialChar(password) {
		return fmt.Errorf("password must contain at least one special character")
	}
	return nil
}
func containsUppercase(s string) bool {
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			return true
		}
	}
	return false
}
func containsLowercase(s string) bool {
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			return true
		}
	}
	return false
}
func containsDigit(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}
func containsSpecialChar(s string) bool {
	specialChars := "!@#$%^&*()-_=+[]{}|;:'\",.<>?/`~"
	for _, c := range s {
		for _, sc := range specialChars {
			if c == sc {
				return true
			}
		}
	}
	return false
}
