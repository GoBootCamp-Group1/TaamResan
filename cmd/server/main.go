package main

import (
	"TaamResan/config"
	"TaamResan/service"
	"flag"
	"log"
	"log/slog"
	//http_server "TaamResan/api/http"
)

var configPath = flag.String("config", "", "configuration path")

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("failed get config", err)
	}

	_, err = service.NewAppContainer(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	//http_server.Run(cfg.Server, app)
}
