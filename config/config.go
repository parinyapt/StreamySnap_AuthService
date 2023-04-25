package config

import (
	// "log"
	"os"
)

func InitializeConfig() {
	initializeDeployModeFlag()

	if os.Getenv("DEPLOY_MODE") == "development" {
		initializeEnvironmentFile()
	}
	if os.Getenv("DEPLOY_MODE") == "production" {
		initializeSetGinReleaseMode()
	}
	initializeEnvironmentVariableCheck()

	initializeGlobalTimezone()
	
}