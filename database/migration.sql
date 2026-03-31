-- Migration file untuk membuat tabel users
-- Jalankan dengan: psql -U postgres -d golangdb -f database/migration.sql

-- Drop table jika sudah ada (hati-hati di production!)
DROP TABLE IF EXISTS users;

-- Buat tabel users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Buat index untuk email (untuk pencarian lebih cepat)
CREATE INDEX idx_users_email ON users(email);

-- Insert sample data untuk testing
INSERT INTO users (name, email) VALUES
    ('John Doe', 'john@example.com'),
    ('Jane Smith', 'jane@example.com'),
    ('Bob Johnson', 'bob@example.com');
