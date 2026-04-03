package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version   string
	Port      int
	Service   string
	JwtSecret string

	DBUser    string
	DBPass    string
	DBPort    string
	DBAddr    string
	DBName    string
	DBSslMode string

	DBSuperUser string
	DBSuperDB   string
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

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatalln("")
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatalln("")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatalln("")
	}

	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		log.Fatalln("")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatalln("")
	}

	dbAddr := os.Getenv("DB_ADDRESS")
	if dbAddr == "" {
		log.Fatalln("")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatalln("")
	}

	dbSSlmode := os.Getenv("DB_SSLMODE")
	if dbSSlmode == "" {
		log.Fatalln("")
	}

	dbSuperUser := os.Getenv("DB_SUPERUSER")
	if dbSSlmode == "" {
		log.Fatalln("")
	}

	dbSuperDb := os.Getenv("DB_SUPERDB")
	if dbSSlmode == "" {
		log.Fatalln("")
	}

	configuration = &Config{
		Version:     version,
		Port:        port,
		JwtSecret:   jwtSecret,
		Service:     serviceName,
		DBUser:      dbUser,
		DBPass:      dbPass,
		DBAddr:      dbAddr,
		DBPort:      dbPort,
		DBName:      dbName,
		DBSslMode:   dbSSlmode,
		DBSuperUser: dbSuperUser,
		DBSuperDB:   dbSuperDb,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
