package internal

import (
	"errors"
	"github.com/sirupsen/logrus"
	"motivations-api/api"
	"motivations-api/config"
	"net/http"
)

type APIServer struct {
	config *config.Config
	logger *logrus.Logger
	server *http.Server
}

func Init(config *config.Config, logger *logrus.Logger) *APIServer {

	router := api.Router(logger)

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

	s.logger.Info("Starting API server...")

	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.WithError(err).Fatal("Failed to start API server")
	}

	return nil
}
