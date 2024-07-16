package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"motivations-api/internal/services"
	"net/http"
)

type Handler struct {
	service *services.MotivationService
	logger  *logrus.Logger
}

func New(service *services.MotivationService, logger *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func renderJSON(w http.ResponseWriter, v interface{}, status int) {
	js, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	response, err := h.service.GetMotivations()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(w, response, http.StatusOK)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	response, err := h.service.CreateMotivation(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	renderJSON(w, response, http.StatusCreated)
}
