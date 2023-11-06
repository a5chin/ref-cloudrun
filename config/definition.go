package config

import (
	"os"
)

type Config struct {
	HOSTNAME   string `env:"hostname"`
	PORT       string `env:"port"`
	DB_USER    string `env:"dbuser"`
	DB_PWD     string `env:"dbpwd"`
	DB_NAME    string `env:"dbname"`
	DB_TCPHOST string `env:"dbtcphost"`
	DB_PORT    string `env:"dbport"`
}

func Load() *Config {
	return &Config{
		HOSTNAME:   os.Getenv("HOSTNAME"),
		PORT:       os.Getenv("PORT"),
		DB_USER:    os.Getenv("DB_USER"),
		DB_PWD:     os.Getenv("DB_PWD"),
		DB_NAME:    os.Getenv("DB_NAME"),
		DB_TCPHOST: os.Getenv("DB_TCPHOST"),
		DB_PORT:    os.Getenv("DB_PORT"),
	}
}
