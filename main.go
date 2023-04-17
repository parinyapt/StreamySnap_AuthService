package main

import (
	"os"

	"github.com/parinyapt/StreamySnap_AuthService/config"
)

func main() {
	config.FlagSetup()
	if os.Getenv("DEPLOY_MODE") == "development" {
		config.EnvironmentFileSetup()
	}
	if os.Getenv("DEPLOY_MODE") == "production" {
		config.GinReleaseModeSetup()
	}
	config.EnvironmentVariableCheck()
	config.TimezoneSetup()
}
