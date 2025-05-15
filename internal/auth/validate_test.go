package auth

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "valid email",
			email:   "user@example.com",
			wantErr: false,
		},
		{
			name:    "missing @",
			email:   "userexample.com",
			wantErr: true,
		},
		{
			name:    "missing domain",
			email:   "user@",
			wantErr: true,
		},
		{
			name:    "missing username",
			email:   "@example.com",
			wantErr: true,
		},
		{
			name:    "uppercase valid",
			email:   "USER@EXAMPLE.COM",
			wantErr: false,
		},
		{
			name:    "special chars valid",
			email:   "user+test@example.com",
			wantErr: false,
		},
		{
			name:    "invalid TLD",
			email:   "user@example.c",
			wantErr: true,
		},
		{
			name:    "empty string",
			email:   "",
			wantErr: true,
		},
		{
			name:    "invalid characters",
			email:   "user@exa$mple.com",
			wantErr: true,
		},
		{
			name:    "invalid domain",
			email:   "user@invalid_domain",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidEmail(tt.email)
			if (got != nil) != tt.wantErr {
				t.Errorf("%s: IsValidEmail(%q) = %v, want %v", tt.name, tt.email, got, tt.wantErr)
			}
		})
	}
}

func TestIsValidPassword(t *testing.T) {
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
		// Forbidden character tests
		{
			name:     "password with space",
			password: "my Secret123!",
			wantErr:  true,
		},
		{
			name:     "password with double quote",
			password: "mySecret\"123!",
			wantErr:  true,
		},
		{
			name:     "password with backslash",
			password: "mySecret\\123!",
			wantErr:  true,
		},
		{
			name:     "password with single quote",
			password: "mySecret'123!",
			wantErr:  true,
		},
		{
			name:     "password with forward slash",
			password: "mySecret/123!",
			wantErr:  true,
		},
		// Invalid password length tests
		{
			name:     "empty password",
			password: "",
			wantErr:  true,
		},
		{
			name:     "password that is too short",
			password: "short",
			wantErr:  true,
		},
		{
			name:     "password that is too long",
			password: "thispasswordiswaytoolongandshouldnotbeacceptedbythehashingfunctionbecauseitexceedsthelimit",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsValidPassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s: IsValidPassword(%q) error = %v, wantErr %v", tt.name, tt.password, err, tt.wantErr)
			}
		})
	}
}
