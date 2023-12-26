package service

import (
	"github.com/murilogilfelpeto/goJob/configuration"
	"gorm.io/gorm"
)

var (
	logger   *configuration.Logger
	database *gorm.DB
)

func init() {
	database = configuration.GetDatabase()
	logger = configuration.GetLogger("handler")
}
