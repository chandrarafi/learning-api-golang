package services

import (
	"api-golang/models"
	"api-golang/repositories"
	"context"
)

// UserService adalah interface untuk business logic user
type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userService struct {
	repo repositories.UserRepository
}

// NewUserService membuat instance baru dari UserService
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser membuat user baru dengan validasi
func (s *userService) CreateUser(ctx context.Context, user *models.User) error {
	// Validasi input
	if err := user.Validate(); err != nil {
		return err
	}

	// Simpan ke database
	return s.repo.Create(ctx, user)
}

// GetAllUsers mengambil semua user
func (s *userService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAll(ctx)
}

// GetUserByID mengambil user berdasarkan ID
func (s *userService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}

// DeleteUser menghapus user berdasarkan ID
func (s *userService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
