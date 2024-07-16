package services

import (
	"github.com/sirupsen/logrus"
	"motivations-api/internal/database"
	"motivations-api/internal/repositories"
	"motivations-api/pkg/modules/motivations"
)

type MotivationService struct {
	repo   *repositories.MotivationRepository
	logger *logrus.Logger
}

func NewMotivationService(db *database.Database, logger *logrus.Logger) *MotivationService {
	repo := repositories.NewMotivationRepository(db, logger)

	return &MotivationService{
		repo:   repo,
		logger: logger,
	}
}

func (s MotivationService) GetMotivations() ([]*motivations.Motivation, error) {

	s.logger.Debug("Start getting list of motivations")
	list, err := s.repo.GetMotivations()

	if err != nil {
		return nil, err
	}

	return list, nil
}
