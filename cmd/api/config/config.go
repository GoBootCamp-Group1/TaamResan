package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Server Server `json:"server"`
	DB     DB     `json:"db"`
}

type Server struct {
	HttpPort               int    `json:"http_port"`
	Host                   string `json:"host"`
	TokenExpMinutes        uint   `json:"token_exp_minutes"`
	RefreshTokenExpMinutes uint   `json:"refresh_token_exp_minute"`
	TokenSecret            string `json:"token_secret"`
}

type DB struct {
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	DBName string `json:"db_name"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = cleanenv.ReadConfig(dir+"/config.json", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
