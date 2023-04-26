package main

import (
	"github.com/parinyapt/StreamySnap_AuthService/config"
	"github.com/parinyapt/StreamySnap_AuthService/database"
	"github.com/parinyapt/StreamySnap_AuthService/logger"
	"github.com/parinyapt/StreamySnap_AuthService/routes"
)

func main() {
	logger.InitializeLogger("development")
	config.InitializeConfig()
	database.InitializeDatabase()
	routes.InitializeRoutes()
}
