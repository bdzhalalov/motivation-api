package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"motivations-api/pkg/modules/motivations"
)

func CreateTable(db *gorm.DB, logger *logrus.Logger) {
	if !db.HasTable(motivations.Motivation{}.TableName()) {
		if err := db.AutoMigrate(&motivations.Motivation{}); err != nil {
			logger.WithError(err.Error).Error("failed to migrate table")

			return
		}

		logger.Info("Motivations table created")
	}
}
