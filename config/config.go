package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version string
	Port    int
}

var configuration *Config

func loadConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatalln("")
	}

	portS := os.Getenv("PORT")
	if portS == "" {
		log.Fatalln("")
	}

	port, err := strconv.Atoi(portS)
	if err != nil {
		log.Fatalln(err)
	}

	configuration = &Config{
		Version: version,
		Port:    port,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
