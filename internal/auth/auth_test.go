package auth

import (
	"testing"
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
			hash, err := HashPassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(hash) == 0 {
					t.Errorf("HashPassword() returned empty hash")
				}
				// Optionally, check that the hash matches the password
				if err := CheckPasswordHash(tt.password, hash); err != nil {
					t.Errorf("HashPassword() produced hash that does not match password: %v", err)
				}
			}
		})
	}
}
