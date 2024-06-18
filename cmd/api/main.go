package main

import (
	httpserver "TaamResan/api/tcp"
	"TaamResan/cmd/api/config"
	"TaamResan/service"
	"flag"
	"log"
	"log/slog"
)

var configPath = flag.String("config", "", "configuration path")

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
