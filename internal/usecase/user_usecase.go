package usecase

import (
	"context"

	"api-golang/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	// Hashing Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return u.repo.Create(ctx, user)
}

func (u *userUsecase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return u.repo.GetAll(ctx)
}

func (u *userUsecase) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *userUsecase) UpdateUser(ctx context.Context, id int, req *domain.UpdateUserRequest) error {
	existingUser, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Merge data (Partial Update)
	if req.Name != nil {
		existingUser.Name = *req.Name
	}
	if req.Email != nil {
		existingUser.Email = *req.Email
	}
	if req.KdKaryawan != nil {
		existingUser.KdKaryawan = *req.KdKaryawan
	}
	if req.IsAdmin != nil {
		existingUser.IsAdmin = *req.IsAdmin
	}
	if req.IsBanned != nil {
		existingUser.IsBanned = *req.IsBanned
	}
	if req.IsVerifikasi != nil {
		existingUser.IsVerifikasi = *req.IsVerifikasi
	}

	if err := existingUser.Validate(); err != nil {
		return err
	}

	return u.repo.Update(ctx, id, existingUser)
}

func (u *userUsecase) DeleteUser(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}
