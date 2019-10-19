package main

import (
	"github.com/mtmendonca/teamder-api/common/logger"
	"github.com/mtmendonca/teamder-api/gateway/api"
)

func main() {
	// Load Config
	cfg := api.LoadConfig()

	// Initialize logger
	log := logger.New()

	// Configure service
	service := &api.Service{
		Cfg: cfg,
		Log: log,
	}

	// Start api
	api.Start(service, ":"+cfg.APIPort)
}
