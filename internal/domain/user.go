package domain

import (
	"context"
	"errors"
	"net/mail"
	"strings"
	"time"
)

type User struct {
	ID              int       `json:"id"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	Password        string    `json:"password"`
	KdKaryawan      string    `json:"kd_karyawan"`
	IsAdmin         bool      `json:"is_admin"`
	IsBanned        bool      `json:"is_banned"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	RememberToken   string    `json:"remember_token"`
	IsVerifikasi    bool      `json:"is_verifikasi"`
}

type UpdateUserRequest struct {
	Email        *string `json:"email"`
	Name         *string `json:"name"`
	KdKaryawan   *string `json:"kd_karyawan"`
	IsAdmin      *bool   `json:"is_admin"`
	IsBanned     *bool   `json:"is_banned"`
	IsVerifikasi *bool   `json:"is_verifikasi"`
}

func (u *User) Validate() error {
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("nama tidak boleh kosong")
	}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("format email tidak valid")
	}
	if len(u.Password) > 0 && len(u.Password) < 6 {
		return errors.New("password minimal 6 karakter")
	}
	return nil
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id int) (*User, error)
	Update(ctx context.Context, id int, user *User) error
	Delete(ctx context.Context, id int) error
}

type UserUsecase interface {
	CreateUser(ctx context.Context, user *User) error
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
	UpdateUser(ctx context.Context, id int, req *UpdateUserRequest) error
	DeleteUser(ctx context.Context, id int) error
}
