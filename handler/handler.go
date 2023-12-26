package handler

import (
	"github.com/murilogilfelpeto/goJob/configuration"
	"gorm.io/gorm"
)

var (
	logger   *configuration.Logger
	database *gorm.DB
)

func InitializeHandler() {
	logger = configuration.GetLogger("handler")
	database = configuration.GetDatabase()
}
