package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"hotstar/internal/application/greeting"
	"hotstar/internal/domain"
)

type Handler struct {
	greetings greeting.Service
}

func NewHandler(greetings greeting.Service) *Handler {
	return &Handler{greetings: greetings}
}

func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	g, err := h.greetings.Hello(name)
	if err != nil {
		if errors.Is(err, domain.ErrEmptyName) {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}
	writeJSON(w, http.StatusOK, g)
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
