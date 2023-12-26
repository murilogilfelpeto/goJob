package main

import (
	"github.com/murilogilfelpeto/goJob/configuration"
	"github.com/murilogilfelpeto/goJob/router"
)

var (
	logger *configuration.Logger
)

func main() {
	logger = configuration.GetLogger("main")
	err := configuration.Init()
	if err != nil {
		logger.Errorf("Error initializing configuration: %v", err)
		return
	}
	router.StartServer()
}
