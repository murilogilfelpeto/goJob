package configuration

import (
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	logger   *Logger
)

func Init() error {

	return nil
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
