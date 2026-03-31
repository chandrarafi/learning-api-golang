package usecase

import (
	"context"

	"api-golang/internal/domain"
)

type agamaUsecase struct {
	repo domain.AgamaRepository
}

func NewAgamaUsecase(repo domain.AgamaRepository) domain.AgamaUsecase {
	return &agamaUsecase{repo: repo}
}

func (u *agamaUsecase) CreateAgama(ctx context.Context, a *domain.Agama) error {
	if err := a.Validate(); err != nil {
		return err
	}
	return u.repo.Create(ctx, a)
}

func (u *agamaUsecase) GetAllAgamas(ctx context.Context) ([]domain.Agama, error) {
	return u.repo.GetAll(ctx)
}

func (u *agamaUsecase) GetAgamaByID(ctx context.Context, id int) (*domain.Agama, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *agamaUsecase) UpdateAgama(ctx context.Context, id int, req *domain.UpdateAgamaRequest) error {
	existing, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if req.KdAgama != nil {
		existing.KdAgama = *req.KdAgama
	}
	if req.NamaAgama != nil {
		existing.NamaAgama = *req.NamaAgama
	}
	if req.Active != nil {
		existing.Active = *req.Active
	}

	if err := existing.Validate(); err != nil {
		return err
	}

	return u.repo.Update(ctx, id, existing)
}

func (u *agamaUsecase) DeleteAgama(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}
