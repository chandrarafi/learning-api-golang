-- Migration file untuk membuat tabel users
-- Jalankan dengan: psql -U postgres -d golangdb -f internal/database/migration.sql

-- Drop table jika sudah ada (hati-hati di production!)
DROP TABLE IF EXISTS users;

-- Buat tabel users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    kd_karyawan VARCHAR(50) NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE,
    is_banned BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    email_verified_at TIMESTAMP NULL,
    remember_token VARCHAR(255) NULL,
    is_verifikasi BOOLEAN DEFAULT FALSE
);

-- Buat index untuk email (untuk pencarian lebih cepat)
CREATE INDEX idx_users_email ON users(email);

-- Insert sample data untuk testing
INSERT INTO users (name, email, password, kd_karyawan, is_admin) VALUES
    ('John Doe', 'john@example.com', 'hasashed123', 'KRY001', TRUE),
    ('Jane Smith', 'jane@example.com', 'hashed456', 'KRY002', FALSE);

-- Drop table agamas jika sudah ada (hati-hati di production!)
DROP TABLE IF EXISTS agamas;

-- Buat tabel agamas
CREATE TABLE agamas (
    id SERIAL PRIMARY KEY,
    kd_agama VARCHAR(16) NOT NULL UNIQUE,
    nama_agama VARCHAR(100) NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert sample data untuk testing
INSERT INTO agamas (kd_agama, nama_agama, active) VALUES
    ('ISL', 'Islam', TRUE),
    ('KRS', 'Kristen', TRUE),
    ('KTL', 'Katolik', TRUE),
    ('HND', 'Hindu', TRUE),
    ('BDH', 'Buddha', TRUE),
    ('KHC', 'Konghucu', TRUE);
