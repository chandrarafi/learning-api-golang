package db

import (
	"api-golang/config"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse config gagal: %w", err)
	}

	poolCfg.MaxConns = 10                      // maks koneksi aktif
	poolCfg.MinConns = 2                       // min koneksi idle
	poolCfg.MaxConnLifetime = 1 * time.Hour    // umur maks tiap koneksi
	poolCfg.MaxConnIdleTime = 30 * time.Minute // idle sebelum ditutup

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("buat pool gagal: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping gagal — cek host/user/password: %w", err)
	}

	return pool, nil
}
