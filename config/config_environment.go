package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvironmentFileSetup() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[Error]->Failed to load environment file : %s", err)
	}
}

func EnvironmentVariableCheck() {
	var requireEnvVariableList = []string{
		"TZ",
		"DATABASE_MYSQL_HOST",
		"DATABASE_MYSQL_DBNAME",
		"DATABASE_MYSQL_USERNAME",
		"DATABASE_MYSQL_PASSWORD",
	}

	for _, v := range requireEnvVariableList {
		if len([]byte(os.Getenv(v))) == 0 {
			log.Fatalf("[Error]->Environment Variable '%s' is not set", v)
		}
	}
}