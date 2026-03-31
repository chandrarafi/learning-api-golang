package httpd

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api-golang/internal/domain"
)

type AgamaHTTPHandler struct {
	usecase domain.AgamaUsecase
}

func NewAgamaHTTPHandler(usecase domain.AgamaUsecase) *AgamaHTTPHandler {
	return &AgamaHTTPHandler{usecase: usecase}
}

// RegisterRoutes mendaftarkan seluruh endpoint milik modul Agama
func (h *AgamaHTTPHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /agamas", h.CreateAgama)
	mux.HandleFunc("GET /agamas", h.GetAgamas)
	mux.HandleFunc("PATCH /agamas/{id}", h.UpdateAgama)
	mux.HandleFunc("DELETE /agamas/{id}", h.DeleteAgama)
}

func (h *AgamaHTTPHandler) CreateAgama(w http.ResponseWriter, r *http.Request) {
	var a domain.Agama
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		respondWithError(w, http.StatusBadRequest, "format JSON tidak valid")
		return
	}

	if err := h.usecase.CreateAgama(r.Context(), &a); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, a)
}

func (h *AgamaHTTPHandler) GetAgamas(w http.ResponseWriter, r *http.Request) {
	agamas, err := h.usecase.GetAllAgamas(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "gagal mengambil data agama")
		return
	}
	respondWithJSON(w, http.StatusOK, agamas)
}

func (h *AgamaHTTPHandler) UpdateAgama(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID harus berupa angka")
		return
	}

	var req domain.UpdateAgamaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "format JSON tidak valid")
		return
	}

	if err := h.usecase.UpdateAgama(r.Context(), id, &req); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "agama berhasil di-patch/update"})
}

func (h *AgamaHTTPHandler) DeleteAgama(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID harus berupa angka")
		return
	}

	if err := h.usecase.DeleteAgama(r.Context(), id); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
