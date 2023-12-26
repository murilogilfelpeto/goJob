package configuration

import (
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	logger   *Logger
)

func Init() error {
	var err error
	database, err = InitializeMySql()
	if err != nil {
		logger.Errorf("Error initializing MySQL connection: %v", err)
		return err
	}

	return nil
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}

func GetDatabase() *gorm.DB {
	return database
}
