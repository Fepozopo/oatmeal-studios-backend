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
