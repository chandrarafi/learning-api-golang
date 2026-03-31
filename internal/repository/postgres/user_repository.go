package postgres

import (
	"context"
	"errors"

	"api-golang/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, u *domain.User) error {
	query := `INSERT INTO users (
		name, email, password, kd_karyawan, is_admin, is_banned, is_verifikasi
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7
	) RETURNING id`

	err := r.db.QueryRow(ctx, query,
		u.Name, u.Email, u.Password, u.KdKaryawan, u.IsAdmin, u.IsBanned, u.IsVerifikasi,
	).Scan(&u.ID)

	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	query := `SELECT id, name, email, kd_karyawan, is_admin, is_banned, is_verifikasi FROM users`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(
			&u.ID, &u.Name, &u.Email, &u.KdKaryawan, &u.IsAdmin, &u.IsBanned, &u.IsVerifikasi,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	query := `SELECT id, name, email, kd_karyawan, is_admin, is_banned, is_verifikasi FROM users WHERE id = $1`
	var u domain.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&u.ID, &u.Name, &u.Email, &u.KdKaryawan, &u.IsAdmin, &u.IsBanned, &u.IsVerifikasi,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user tidak ditemukan")
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) Update(ctx context.Context, id int, u *domain.User) error {
	query := `UPDATE users SET 
		name = $1, email = $2, kd_karyawan = $3, 
		is_admin = $4, is_banned = $5, is_verifikasi = $6
	WHERE id = $7`

	result, err := r.db.Exec(ctx, query,
		u.Name, u.Email, u.KdKaryawan, u.IsAdmin, u.IsBanned, u.IsVerifikasi, id,
	)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("user tidak ditemukan")
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("user tidak ditemukan")
	}
	return nil
}
