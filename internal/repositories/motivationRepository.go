package repositories

import (
	"github.com/sirupsen/logrus"
	"motivations-api/internal/database"
	"motivations-api/pkg/modules/motivations"
)

type MotivationRepository struct {
	db     *database.Database
	logger *logrus.Logger
}

func NewMotivationRepository(db *database.Database, logger *logrus.Logger) *MotivationRepository {
	return &MotivationRepository{
		db:     db,
		logger: logger,
	}
}

func (m MotivationRepository) GetMotivations() ([]*motivations.Motivation, error) {
	var list []*motivations.Motivation

	if err := m.db.Connection.Find(&list).Error; err != nil {
		m.logger.Errorf("Error while getting list of motivations: %s", err)
		return nil, err
	}

	return list, nil
}
