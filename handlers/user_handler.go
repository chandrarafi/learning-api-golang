package handlers

import (
	"api-golang/models"
	"api-golang/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Service services.UserService
}

// ErrorResponse adalah struktur untuk response error
type ErrorResponse struct {
	Error string `json:"error"`
}

// CreateUser menangani request POST untuk membuat user baru
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "format JSON tidak valid")
		return
	}

	ctx := r.Context()

	// Panggil service untuk create user (sudah termasuk validasi)
	if err := h.Service.CreateUser(ctx, &user); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Response sukses
	respondWithJSON(w, http.StatusCreated, user)
}

// GetUsers menangani request GET untuk mengambil semua user
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.Service.GetAllUsers(ctx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "gagal mengambil data users")
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

// DeleteUser menangani request DELETE untuk menghapus user
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari path parameter
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID harus berupa angka")
		return
	}

	ctx := r.Context()

	// Panggil service untuk delete user
	if err := h.Service.DeleteUser(ctx, id); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	// Response sukses (204 No Content)
	w.WriteHeader(http.StatusNoContent)
}

// respondWithJSON adalah helper function untuk mengirim response JSON
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// respondWithError adalah helper function untuk mengirim error response
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, ErrorResponse{Error: message})
}
