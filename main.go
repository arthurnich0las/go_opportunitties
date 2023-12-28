package main

import (
	"github.com/arthurnich0las/go_opportunities/config"
	"github.com/arthurnich0las/go_opportunities/router"
)

func main() {
	logger := config.GetLogger("main")
	// Inicializing configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initializing router
	router.Initialize()
}
