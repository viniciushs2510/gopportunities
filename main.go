package main

import (
	"github.com/viniciushs2510/gopportunities/config"
	"github.com/viniciushs2510/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	router.Initialize()
}
