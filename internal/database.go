package internal

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"motivations-api/config"
	"motivations-api/pkg/migrations"
)

type Database struct {
	connection *gorm.DB
	logger     *logrus.Logger
}

func ConnectToDB(config *config.Config, logger *logrus.Logger) (*Database, error) {
	databaseURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DbUser,
		config.DbPass,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)

	db, err := gorm.Open("mysql", databaseURI)

	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	logger.Info("Successfully connected to database")

	migrations.CreateTable(db, logger)

	return &Database{
		connection: db,
		logger:     logger,
	}, nil
}

func (db *Database) Close() error {
	if db.connection != nil {
		return db.connection.Close()
	}

	db.logger.Info("Connection to database was closed")
	return nil
}
