package main

import (
	// "github.com/parinyapt/StreamySnap_AuthService/config"
	// "os"

	"github.com/parinyapt/StreamySnap_AuthService/config"
	"github.com/parinyapt/StreamySnap_AuthService/logger"
)

func main() {
	logger.InitializeLogger("production")

	config.InitializeConfig()
}
