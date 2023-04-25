package main

import (
	"os"

	"github.com/parinyapt/StreamySnap_AuthService/config"
	"github.com/parinyapt/StreamySnap_AuthService/database"
	"github.com/parinyapt/StreamySnap_AuthService/logger"
)

func main() {
	config.InitializeConfig()
	logger.InitializeLogger(os.Getenv("DEPLOY_MODE"))
	database.InitializeDatabase()

}
