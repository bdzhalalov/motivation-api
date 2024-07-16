package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"motivations-api/internal/handler"
	"motivations-api/internal/services"
)

func Router(logger *logrus.Logger, service *services.MotivationService) *mux.Router {
	router := mux.NewRouter()

	h := handler.New(service, logger)

	// group of routes for motivations
	group := router.PathPrefix("/motivations").Subrouter()
	group.HandleFunc("/", h.List).Methods("GET")
	group.HandleFunc("/", h.Create).Methods("POST")

	return router
}
