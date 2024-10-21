package repositories

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"motivations-api/internal/database"
	customErrors "motivations-api/internal/errors"
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

func (m MotivationRepository) GetMotivations() ([]*motivations.Motivation, *customErrors.BaseError) {
	var list []*motivations.Motivation

	if err := m.db.Connection.Find(&list).Error; err != nil {
		m.logger.Errorf("Error while getting list of motivations: %s", err)

		var internalError customErrors.BaseAbstractError = &customErrors.InternalServerError{}
		return nil, internalError.New()
	}

	return list, nil
}

func (m MotivationRepository) CreateMotivation(motivation *motivations.Motivation) (*motivations.Motivation, *customErrors.BaseError) {
	if err := m.db.Connection.Create(&motivation).Error; err != nil {
		m.logger.Errorf("Error while creating motivation: %s", err)

		var internalError customErrors.BaseAbstractError = &customErrors.InternalServerError{}
		return nil, internalError.New()
	}

	return motivation, nil
}

func (m MotivationRepository) GetMotivationById(id string) (*motivations.Motivation, *customErrors.BaseError) {
	var motivation motivations.Motivation

	if err := m.db.Connection.First(&motivation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var notFoundError customErrors.BaseAbstractError = &customErrors.NotFoundError{}
			return nil, notFoundError.New()
		}
		m.logger.Errorf("Error while getting motivation by id: %s", err)

		var internalError customErrors.BaseAbstractError = &customErrors.InternalServerError{}
		return nil, internalError.New()
	}

	return &motivation, nil
}
