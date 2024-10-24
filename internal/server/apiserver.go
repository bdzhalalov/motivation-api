package server

import (
	"errors"
	"github.com/sirupsen/logrus"
	"motivations-api/api"
	"motivations-api/config"
	"motivations-api/internal/database"
	"motivations-api/internal/services"
	"net/http"
)

type APIServer struct {
	config *config.Config
	logger *logrus.Logger
	server *http.Server
}

func Init(config *config.Config, db *database.Database, logger *logrus.Logger) *APIServer {

	s := services.NewMotivationService(db, logger)

	router := api.Router(logger, s)

	return &APIServer{
		config: config,
		logger: logger,
		server: &http.Server{
			Addr:    config.Addr,
			Handler: router,
		},
	}
}

func (s *APIServer) Run() error {

	s.logger.Info("Running API server on port" + s.config.Addr)

	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.WithError(err).Fatal("Failed to start API server")
	}

	return nil
}
