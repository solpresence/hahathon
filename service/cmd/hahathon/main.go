package main

import (
	"hahathon/internal/config"
	"hahathon/internal/logger"
	"hahathon/internal/server/endpoints"
	"hahathon/internal/utils"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	srv := endpoints.NewRestApi(cfg, log)
	srv.CreateServer()
	defer srv.Close()

	utils.Shutdown(log)

	log.Info("Server stoped")
}
