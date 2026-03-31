package httpd

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"api-golang/internal/repository/postgres"
	"api-golang/internal/usecase"
)

// SetupRouter mem-wiring seluruh dependencies per layer dan me-register routes.
func SetupRouter(mux *http.ServeMux, pool *pgxpool.Pool) {
	// --- Inisialisasi Module User ---
	userRepo := postgres.NewUserRepository(pool)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := NewUserHTTPHandler(userUsecase)
	userHandler.RegisterRoutes(mux)

	// --- Inisialisasi Module Agama ---
	agamaRepo := postgres.NewAgamaRepository(pool)
	agamaUsecase := usecase.NewAgamaUsecase(agamaRepo)
	agamaHandler := NewAgamaHTTPHandler(agamaUsecase)
	agamaHandler.RegisterRoutes(mux)
	
	// Tambahkan injeksi modul lain di sini di masa depan...
}
