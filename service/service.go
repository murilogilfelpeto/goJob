package service

import (
	"github.com/murilogilfelpeto/goJob/configuration"
	"gorm.io/gorm"
)

var (
	logger   *configuration.Logger
	database *gorm.DB
)

func Initialize() {
	database = configuration.GetDatabase()
	logger = configuration.GetLogger("handler")
}
