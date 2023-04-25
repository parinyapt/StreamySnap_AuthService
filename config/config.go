package config

import (
	"os"

	"github.com/parinyapt/StreamySnap_AuthService/logger"
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
	
	logger.Info("Initialize Config Success")
}