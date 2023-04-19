package main

import (
	// "os"

	"fmt"
	"time"

	"github.com/parinyapt/StreamySnap_AuthService/config"
)

func main() {
	// config.FlagSetup()
	// if os.Getenv("DEPLOY_MODE") == "development" {
		config.EnvironmentFileSetup()
	// }
	// if os.Getenv("DEPLOY_MODE") == "production" {
	// 	config.GinReleaseModeSetup()
	// }
	// config.EnvironmentVariableCheck()
	// config.TimezoneSetup()

	config.GlobalSetup()

	fmt.Println(time.Local)
}
