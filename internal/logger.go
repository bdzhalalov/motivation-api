package lib

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"motivations-api/config"
	"os"
)

func Logger(config *config.Config) *logrus.Logger {

	log := logrus.New()

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		fmt.Println(err.Error())
		log.Error("Can't parse log level")
		level = logrus.DebugLevel
	}

	log.SetLevel(level)

	file, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		fmt.Println("Failed to log to file, using default stderr")
	}

	return log
}
