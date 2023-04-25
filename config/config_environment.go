package config

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/parinyapt/StreamySnap_AuthService/logger"
)

func initializeEnvironmentFile() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal("Failed to load environment file", logger.Field("error", err))
	}
}

func initializeEnvironmentVariableCheck() {
	var requireEnvVariableList = []string{
		"TZ",
		"DATABASE_MYSQL_HOST",
		"DATABASE_MYSQL_DBNAME",
		"DATABASE_MYSQL_USERNAME",
		"DATABASE_MYSQL_PASSWORD",
	}

	for _, v := range requireEnvVariableList {
		if len([]byte(os.Getenv(v))) == 0 {
			logger.Fatal("Environment Variable '" + v + "' is not set")
		}
	}
}