package main

import (
	httpserver "TaamResan/api/tcp"
	"TaamResan/cmd/api/config"
	"TaamResan/service"
	"log"
	"log/slog"
)

func main() {
	cfg, errConfig := config.NewConfig()
	if errConfig != nil {
		slog.Error("failed get config", errConfig)
	}

	app, errAppContainer := service.NewAppContainer(*cfg)
	if errAppContainer != nil {
		log.Fatal(errAppContainer)
	}

	httpserver.Run(cfg.Server, app)
}
