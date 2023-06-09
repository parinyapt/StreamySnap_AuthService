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
		"APP_NAME",
		"APP_BASE_URL",
		"DATABASE_TABLE_PREFIX",
		"DATABASE_MYSQL_HOST",
		"DATABASE_MYSQL_DBNAME",
		"DATABASE_MYSQL_USERNAME",
		"DATABASE_MYSQL_PASSWORD",
		"JWT_SIGN_KEY_AUTHSESSION",
		"JWT_SIGN_KEY_TEMPORARYTOKEN",
		"JWT_SIGN_KEY_ACCESSTOKEN",
		"JWT_SIGN_KEY_REFRESHTOKEN",
	}

	for _, v := range requireEnvVariableList {
		if len([]byte(os.Getenv(v))) == 0 {
			logger.Fatal("Environment Variable '" + v + "' is not set")
		}
	}

	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "80")
	}
}