package main

import (
	"fmt"
	"log"
	"net/http"

	"api-golang/config"
	db "api-golang/database"
	"api-golang/handlers"
	"api-golang/middleware"
	"api-golang/repositories"
	"api-golang/services"
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

	// Inisialisasi layers (Repository -> Service -> Handler)
	userRepo := &repositories.UserRepo{DB: pool}
	userService := services.NewUserService(userRepo)
	userHandler := &handlers.UserHandler{Service: userService}

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("GET /users", userHandler.GetUsers)
	mux.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)

	// Wrap dengan middleware (Recovery -> Logger -> Mux)
	handler := middleware.Recovery(middleware.Logger(mux))

	// Jalankan server
	port := ":8080"
	fmt.Printf("🚀 Server berjalan di http://localhost%s\n", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
