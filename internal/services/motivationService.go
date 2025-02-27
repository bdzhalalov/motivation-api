package services

import (
	"github.com/sirupsen/logrus"
	"motivations-api/internal/database"
	"motivations-api/internal/errors"
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

func (s MotivationService) GetMotivations() ([]*motivations.Motivation, *errors.BaseError) {

	s.logger.Debug("Start getting list of motivations")
	list, err := s.repo.GetMotivations()

	return list, err
}

func (s MotivationService) CreateMotivation(motivation *motivations.Motivation) (*motivations.Motivation, *errors.BaseError) {

	s.logger.Debug("Start creating motivation")

	res, err := s.repo.CreateMotivation(motivation)

	return res, err
}

func (s MotivationService) GetMotivationById(id string) (*motivations.Motivation, *errors.BaseError) {
	s.logger.Debug("Start getting motivation by id")

	res, err := s.repo.GetMotivationById(id)
	return res, err
}

func (s MotivationService) UpdateMotivationById(id string, motivation *motivations.Motivation) (*motivations.Motivation, *errors.BaseError) {
	s.logger.Debug("Start updating motivation")

	res, err := s.repo.UpdateMotivationById(id, motivation.Motivation)

	return res, err
}

func (s MotivationService) DeleteMotivationById(id string) *errors.BaseError {
	s.logger.Debug("Start deleting motivation by id")

	err := s.repo.DeleteMotivationById(id)

	return err
}

func (s MotivationService) GetRandomMotivation() (*motivations.Motivation, *errors.BaseError) {
	s.logger.Debug("Start getting random motivation")

	res, err := s.repo.GetRandomMotivation()

	return res, err
}
