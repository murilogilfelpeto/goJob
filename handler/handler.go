package handler

import (
	"github.com/murilogilfelpeto/goJob/configuration"
)

var (
	logger *configuration.Logger
)

func InitializeHandler() {
	logger = configuration.GetLogger("handler")
}
