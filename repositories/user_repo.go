package repositories

import (
	"api-golang/models"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserRepository adalah interface untuk operasi database user
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetAll(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id int) (*models.User, error)
	Delete(ctx context.Context, id int) error
}

type UserRepo struct {
	DB *pgxpool.Pool
}

// Create menambahkan user baru ke database
func (r *UserRepo) Create(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	err := r.DB.QueryRow(ctx, query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetAll mengambil semua user dari database
func (r *UserRepo) GetAll(ctx context.Context) ([]models.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	// Cek error setelah iterasi
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetByID mengambil user berdasarkan ID
func (r *UserRepo) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"
	var user models.User
	err := r.DB.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user tidak ditemukan")
		}
		return nil, err
	}
	return &user, nil
}

// Delete menghapus user berdasarkan ID
func (r *UserRepo) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	result, err := r.DB.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	// Cek apakah ada row yang terhapus
	if result.RowsAffected() == 0 {
		return errors.New("user tidak ditemukan")
	}

	return nil
}
