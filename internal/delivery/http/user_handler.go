package httpd

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api-golang/internal/domain"
)

type UserHTTPHandler struct {
	usecase domain.UserUsecase
}

func NewUserHTTPHandler(usecase domain.UserUsecase) *UserHTTPHandler {
	return &UserHTTPHandler{usecase: usecase}
}

// RegisterRoutes mendaftarkan seluruh endpoint milik modul User
func (h *UserHTTPHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /users", h.CreateUser)
	mux.HandleFunc("GET /users", h.GetUsers)
	mux.HandleFunc("PATCH /users/{id}", h.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", h.DeleteUser)
}

// ErrorResponse adalah struktur untuk response error
type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *UserHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "format JSON tidak valid")
		return
	}

	if err := h.usecase.CreateUser(r.Context(), &user); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func (h *UserHTTPHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.usecase.GetAllUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "gagal mengambil data user")
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func (h *UserHTTPHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID harus berupa angka")
		return
	}

	var req domain.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "format JSON tidak valid")
		return
	}

	if err := h.usecase.UpdateUser(r.Context(), id, &req); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "user berhasil di-patch/update"})
}

func (h *UserHTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID harus berupa angka")
		return
	}

	if err := h.usecase.DeleteUser(r.Context(), id); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, ErrorResponse{Error: message})
}
