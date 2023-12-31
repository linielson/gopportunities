package main

import (
	"github.com/linielson/gopportunities/config"
	"github.com/linielson/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.NewLogger("main")
	//logger = config.GetLogger("main") // it works too
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	router.Initialize()
}
