package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "valid password",
			password: "mySecret123!",
			wantErr:  false,
		},
		{
			name:     "long password",
			password: "aVeryLongPasswordThatShouldStillWorkJustFine1234567890!#$%^&*",
			wantErr:  false,
		},
		{
			name:     "password with special characters",
			password: "mySecret123!@#$%^&*()",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashPassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(hash) == 0 {
					t.Errorf("HashPassword() returned empty hash")
				}
				if err := CheckPasswordHash(tt.password, hash); err != nil {
					t.Errorf("HashPassword() produced hash that does not match password: %v", err)
				}
			}
		})
	}
}

func TestMakeJWT(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "mySecret"

	token, err := MakeJWT(userID, tokenSecret, 1*time.Hour)
	if err != nil {
		t.Errorf("Failed to generate JWT: %v", err)
	}

	if len(token) == 0 {
		t.Error("Generated JWT is empty")
	}
}

func TestValidateJWT(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "mySecret"

	token, err := MakeJWT(userID, tokenSecret, 1*time.Hour)
	if err != nil {
		t.Errorf("Failed to generate JWT: %v", err)
	}

	validUserID, err := ValidateJWT(token, tokenSecret)
	if err != nil {
		t.Errorf("Failed to validate JWT: %v", err)
	}

	if validUserID != userID {
		t.Error("Validated user ID is not the same as the generated one")
	}

	invalidToken := "Invalid Token"
	_, err = ValidateJWT(invalidToken, tokenSecret)
	if err == nil {
		t.Error("ValidateJWT should have returned an error")
	}
}

func TestGetBearerToken(t *testing.T) {
	validToken := "myValidToken"
	headers := http.Header{
		"Authorization": []string{"Bearer " + validToken},
	}

	token, err := GetBearerToken(headers)
	if err != nil {
		t.Errorf("Failed to retrieve bearer token: %v", err)
	}

	if token != validToken {
		t.Error("Retrieved token is not the same as the valid one")
	}

	invalidHeaders := http.Header{}
	_, err = GetBearerToken(invalidHeaders)
	if err == nil {
		t.Error("GetBearerToken should have returned an error")
	}

	invalidHeaders = http.Header{
		"Invalid": []string{"Bearer " + validToken},
	}

	_, err = GetBearerToken(invalidHeaders)
	if err == nil {
		t.Error("GetBearerToken should have returned an error")
	}
}
