package main

import (
	"fmt"
	"log"
	"net/http"

	"api-golang/config"
	db "api-golang/database"
	httpd "api-golang/internal/delivery/http"
	"api-golang/middleware"
)

func main() {
	// Load konfigurasi dari .env
	cfg := config.LoadConfig()

	// Koneksi ke database
	pool, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}
	defer pool.Close()
	fmt.Println("✅ Database Terkoneksi via pgxpool!")

	// Setup routes menggunakan polan Centralized Router Clean Architecture
	mux := http.NewServeMux()
	httpd.SetupRouter(mux, pool)

	// Wrap dengan middleware (Recovery -> Logger -> Mux)
	handler := middleware.Recovery(middleware.Logger(mux))

	// Jalankan server
	port := ":8080"
	fmt.Printf("🚀 Server berjalan di http://localhost%s\n", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
