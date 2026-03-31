package domain

import (
	"context"
	"errors"
	"strings"
	"time"
)

type Agama struct {
	ID        int       `json:"id"`
	KdAgama   string    `json:"kd_agama"`
	NamaAgama string    `json:"nama_agama"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateAgamaRequest struct {
	KdAgama   *string `json:"kd_agama"`
	NamaAgama *string `json:"nama_agama"`
	Active    *bool   `json:"active"`
}

func (a *Agama) Validate() error {
	if strings.TrimSpace(a.KdAgama) == "" {
		return errors.New("kode agama tidak boleh kosong")
	}
	if len(a.KdAgama) > 16 {
		return errors.New("kode agama maksimal 16 karakter")
	}

	if strings.TrimSpace(a.NamaAgama) == "" {
		return errors.New("nama agama tidak boleh kosong")
	}
	if len(a.NamaAgama) > 100 {
		return errors.New("nama agama maksimal 100 karakter")
	}
	return nil
}

type AgamaRepository interface {
	Create(ctx context.Context, agama *Agama) error
	GetAll(ctx context.Context) ([]Agama, error)
	GetByID(ctx context.Context, id int) (*Agama, error)
	Update(ctx context.Context, id int, agama *Agama) error
	Delete(ctx context.Context, id int) error
}

type AgamaUsecase interface {
	CreateAgama(ctx context.Context, agama *Agama) error
	GetAllAgamas(ctx context.Context) ([]Agama, error)
	GetAgamaByID(ctx context.Context, id int) (*Agama, error)
	UpdateAgama(ctx context.Context, id int, req *UpdateAgamaRequest) error
	DeleteAgama(ctx context.Context, id int) error
}
