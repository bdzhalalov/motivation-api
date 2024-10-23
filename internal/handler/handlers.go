package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
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
		renderJSON(w, err.Message, err.Code)
		return
	}

	renderJSON(w, response, http.StatusOK)
}

// Create TODO: Валидация данных для request body
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		renderJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, e := h.service.CreateMotivation(body)
	if e != nil {
		renderJSON(w, e.Message, e.Code)
		return
	}

	renderJSON(w, response, http.StatusCreated)
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	response, err := h.service.GetMotivationById(id)
	if err != nil {
		renderJSON(w, err.Message, err.Code)
		return
	}

	renderJSON(w, response, http.StatusOK)
}

// Update TODO: валидация данных для request body
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		renderJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, e := h.service.UpdateMotivationById(id, body)
	if e != nil {
		renderJSON(w, e.Message, e.Code)
		return
	}

	renderJSON(w, response, http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := h.service.DeleteMotivationById(id)
	if err != nil {
		renderJSON(w, err.Message, err.Code)
		return
	}

	renderJSON(w, "", http.StatusNoContent)
}
