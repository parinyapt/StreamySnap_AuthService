package config

import (
	"flag"
	"log"
	"os"
)

func FlagSetup() {
	DeployModeFlag := flag.String("mode", "development", "deploy mode (development, production)")
	flag.Parse()

	if (*DeployModeFlag == "development") || (*DeployModeFlag == "production") {
		os.Setenv("DEPLOY_MODE", *DeployModeFlag)
		log.Println("Deploy Mode : " + os.Getenv("DEPLOY_MODE"))
	} else {
		log.Fatalf("[Error]->Please set deploy mode to 'development' or 'production'")
		return
	}
}