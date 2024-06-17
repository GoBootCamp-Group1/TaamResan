package main

import (
	"TaamResan/config"
	"TaamResan/service"
	"flag"
	"log"
	"os"
	//http_server "TaamResan/api/http"
)

var configPath = flag.String("config", "", "configuration path")

func main() {
	cfg := readConfig()

	_, err := service.NewAppContainer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	//http_server.Run(cfg.Server, app)
}

func readConfig() config.Config {
	flag.Parse()

	if cfgPathEnv := os.Getenv("APP_CONFIG_PATH"); len(cfgPathEnv) > 0 {
		*configPath = cfgPathEnv
	}

	if len(*configPath) == 0 {
		log.Fatal("configuration file not found")
	}

	cfg, err := config.ReadStandard(*configPath)

	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
