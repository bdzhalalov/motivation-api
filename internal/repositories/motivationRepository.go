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

func (m MotivationRepository) UpdateMotivationById(id string, motivation string) (*motivations.Motivation, *customErrors.BaseError) {
	res, err := m.GetMotivationById(id)
	if err != nil {
		return nil, err
	}

	if err := m.db.Connection.Model(&res).
		Where("id = ?", id).Update("motivation", motivation).Error; err != nil {
		m.logger.Errorf("Error while updating motivation: %s", err)
		var internalError customErrors.BaseAbstractError = &customErrors.InternalServerError{}
		return nil, internalError.New()
	}

	return res, nil
}

func (m MotivationRepository) DeleteMotivationById(id string) *customErrors.BaseError {
	var motivation motivations.Motivation

	if err := m.db.Connection.First(&motivation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var notFoundError customErrors.BaseAbstractError = &customErrors.NotFoundError{}
			return notFoundError.New()
		}
	}

	if err := m.db.Connection.Delete(&motivation, id).Error; err != nil {
		m.logger.Errorf("Error while deleting motivation by id: %s", err)

		var internalError customErrors.BaseAbstractError = &customErrors.InternalServerError{}
		return internalError.New()
	}

	return nil
}

func (m MotivationRepository) GetRandomMotivation() (*motivations.Motivation, *customErrors.BaseError) {
	var motivation motivations.Motivation

	// TODO: not optimized method for getting random entity
	if err := m.db.Connection.Raw("SELECT * FROM motivations ORDER BY RAND() LIMIT 1").Scan(&motivation).Error; err != nil {
		m.logger.Errorf("Error while getting random motivation: %s", err)

		var internalError customErrors.BaseAbstractError = &customErrors.InternalServerError{}
		return nil, internalError.New()
	}

	return &motivation, nil
}
