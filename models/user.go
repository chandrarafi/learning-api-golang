package models

import (
	"errors"
	"regexp"
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Validate melakukan validasi data user
func (u *User) Validate() error {
	// Validasi name
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("name tidak boleh kosong")
	}
	if len(u.Name) < 3 {
		return errors.New("name minimal 3 karakter")
	}
	if len(u.Name) > 100 {
		return errors.New("name maksimal 100 karakter")
	}

	// Validasi email
	if strings.TrimSpace(u.Email) == "" {
		return errors.New("email tidak boleh kosong")
	}
	if !isValidEmail(u.Email) {
		return errors.New("format email tidak valid")
	}

	return nil
}

// isValidEmail memeriksa format email menggunakan regex
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
