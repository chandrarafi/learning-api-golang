package postgres

import (
	"context"
	"errors"

	"api-golang/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type agamaRepository struct {
	db *pgxpool.Pool
}

func NewAgamaRepository(db *pgxpool.Pool) domain.AgamaRepository {
	return &agamaRepository{db: db}
}

func (r *agamaRepository) Create(ctx context.Context, a *domain.Agama) error {
	query := `INSERT INTO agamas (
		kd_agama, nama_agama, active, created_at, updated_at
	) VALUES (
		$1, $2, $3, $4, $5
	) RETURNING id`

	err := r.db.QueryRow(ctx, query,
		a.KdAgama, a.NamaAgama, a.Active, a.CreatedAt, a.UpdatedAt,
	).Scan(&a.ID)

	if err != nil {
		return err
	}
	return nil
}

func (r *agamaRepository) GetAll(ctx context.Context) ([]domain.Agama, error) {
	query := `SELECT id, kd_agama, nama_agama, active, created_at, updated_at FROM agamas`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agamas []domain.Agama
	for rows.Next() {
		var a domain.Agama
		if err := rows.Scan(
			&a.ID, &a.KdAgama, &a.NamaAgama, &a.Active, &a.CreatedAt, &a.UpdatedAt,
		); err != nil {
			return nil, err
		}
		agamas = append(agamas, a)
	}
	return agamas, rows.Err()
}

func (r *agamaRepository) GetByID(ctx context.Context, id int) (*domain.Agama, error) {
	query := `SELECT id, kd_agama, nama_agama, active, created_at, updated_at FROM agamas WHERE id = $1`
	var a domain.Agama
	err := r.db.QueryRow(ctx, query, id).Scan(
		&a.ID, &a.KdAgama, &a.NamaAgama, &a.Active, &a.CreatedAt, &a.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("agama tidak ditemukan")
		}
		return nil, err
	}
	return &a, nil
}

func (r *agamaRepository) Update(ctx context.Context, id int, a *domain.Agama) error {
	query := `UPDATE agamas SET 
		kd_agama = $1, nama_agama = $2, active = $3, updated_at = CURRENT_TIMESTAMP
	WHERE id = $4`

	result, err := r.db.Exec(ctx, query,
		a.KdAgama, a.NamaAgama, a.Active, id,
	)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("agama tidak ditemukan")
	}
	return nil
}

func (r *agamaRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM agamas WHERE id = $1"
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("agama tidak ditemukan")
	}
	return nil
}
