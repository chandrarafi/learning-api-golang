package models

import (
	"testing"
)

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid user",
			user:    User{Name: "John Doe", Email: "john@example.com"},
			wantErr: false,
		},
		{
			name:    "empty name",
			user:    User{Name: "", Email: "john@example.com"},
			wantErr: true,
			errMsg:  "name tidak boleh kosong",
		},
		{
			name:    "name too short",
			user:    User{Name: "Jo", Email: "john@example.com"},
			wantErr: true,
			errMsg:  "name minimal 3 karakter",
		},
		{
			name:    "empty email",
			user:    User{Name: "John Doe", Email: ""},
			wantErr: true,
			errMsg:  "email tidak boleh kosong",
		},
		{
			name:    "invalid email format",
			user:    User{Name: "John Doe", Email: "invalid-email"},
			wantErr: true,
			errMsg:  "format email tidak valid",
		},
		{
			name:    "invalid email without domain",
			user:    User{Name: "John Doe", Email: "john@"},
			wantErr: true,
			errMsg:  "format email tidak valid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("User.Validate() error message = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}
