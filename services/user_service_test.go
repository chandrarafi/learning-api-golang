package services

import (
	"api-golang/models"
	"context"
	"errors"
	"testing"
)

// mockUserRepo adalah mock repository untuk testing
type mockUserRepo struct {
	createFunc  func(ctx context.Context, user *models.User) error
	getAllFunc  func(ctx context.Context) ([]models.User, error)
	getByIDFunc func(ctx context.Context, id int) (*models.User, error)
	deleteFunc  func(ctx context.Context, id int) error
}

func (m *mockUserRepo) Create(ctx context.Context, user *models.User) error {
	if m.createFunc != nil {
		return m.createFunc(ctx, user)
	}
	return nil
}

func (m *mockUserRepo) GetAll(ctx context.Context) ([]models.User, error) {
	if m.getAllFunc != nil {
		return m.getAllFunc(ctx)
	}
	return []models.User{}, nil
}

func (m *mockUserRepo) GetByID(ctx context.Context, id int) (*models.User, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *mockUserRepo) Delete(ctx context.Context, id int) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(ctx, id)
	}
	return nil
}

func TestUserService_CreateUser(t *testing.T) {
	ctx := context.Background()

	t.Run("success create user", func(t *testing.T) {
		mockRepo := &mockUserRepo{
			createFunc: func(ctx context.Context, user *models.User) error {
				user.ID = 1
				return nil
			},
		}

		service := NewUserService(mockRepo)

		user := &models.User{
			Name:  "John Doe",
			Email: "john@example.com",
		}

		err := service.CreateUser(ctx, user)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if user.ID != 1 {
			t.Errorf("expected ID to be 1, got %d", user.ID)
		}
	})

	t.Run("validation error - empty name", func(t *testing.T) {
		mockRepo := &mockUserRepo{}
		service := NewUserService(mockRepo)

		user := &models.User{
			Name:  "",
			Email: "john@example.com",
		}

		err := service.CreateUser(ctx, user)
		if err == nil {
			t.Error("expected validation error, got nil")
		}
	})

	t.Run("validation error - invalid email", func(t *testing.T) {
		mockRepo := &mockUserRepo{}
		service := NewUserService(mockRepo)

		user := &models.User{
			Name:  "John Doe",
			Email: "invalid-email",
		}

		err := service.CreateUser(ctx, user)
		if err == nil {
			t.Error("expected validation error, got nil")
		}
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo := &mockUserRepo{
			createFunc: func(ctx context.Context, user *models.User) error {
				return errors.New("database error")
			},
		}
		service := NewUserService(mockRepo)

		user := &models.User{
			Name:  "John Doe",
			Email: "john@example.com",
		}

		err := service.CreateUser(ctx, user)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestUserService_GetAllUsers(t *testing.T) {
	ctx := context.Background()

	t.Run("success get all users", func(t *testing.T) {
		expectedUsers := []models.User{
			{ID: 1, Name: "John Doe", Email: "john@example.com"},
			{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
		}

		mockRepo := &mockUserRepo{
			getAllFunc: func(ctx context.Context) ([]models.User, error) {
				return expectedUsers, nil
			},
		}
		service := NewUserService(mockRepo)

		users, err := service.GetAllUsers(ctx)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(users) != 2 {
			t.Errorf("expected 2 users, got %d", len(users))
		}
	})
}

func TestUserService_DeleteUser(t *testing.T) {
	ctx := context.Background()

	t.Run("success delete user", func(t *testing.T) {
		mockRepo := &mockUserRepo{
			deleteFunc: func(ctx context.Context, id int) error {
				return nil
			},
		}
		service := NewUserService(mockRepo)

		err := service.DeleteUser(ctx, 1)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("user not found", func(t *testing.T) {
		mockRepo := &mockUserRepo{
			deleteFunc: func(ctx context.Context, id int) error {
				return errors.New("user tidak ditemukan")
			},
		}
		service := NewUserService(mockRepo)

		err := service.DeleteUser(ctx, 999)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
