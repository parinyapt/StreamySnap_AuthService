package config

import (
	"flag"
	"os"

	"github.com/parinyapt/StreamySnap_AuthService/logger"
)

func initializeDeployModeFlag() {
	DeployModeFlag := flag.String("mode", "development", "deploy mode (development, production)")
	flag.Parse()

	if (*DeployModeFlag == "development") || (*DeployModeFlag == "production") {
		os.Setenv("DEPLOY_MODE", *DeployModeFlag)
		logger.Info("Deploy Mode : " + os.Getenv("DEPLOY_MODE"))
	} else {
		logger.Fatal("Please set deploy mode to 'development' or 'production'")
		return
	}
}