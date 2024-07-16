package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
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

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	response, err := h.service.GetMotivations()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(w, response)
}
