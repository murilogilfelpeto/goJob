package handler

import (
	"github.com/murilogilfelpeto/goJob/configuration"
	"github.com/murilogilfelpeto/goJob/service"
)

var (
	logger *configuration.Logger
)

func InitializeHandler() {
	logger = configuration.GetLogger("handler")
	service.Initialize()
}
