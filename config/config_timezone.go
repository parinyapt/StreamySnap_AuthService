package config

import (
	"log"
	"os"

	"github.com/parinyapt/golang_utils/timezone/v1"
)

func initializeGlobalTimezone() {
	// Global TimeZone Setup
	if err := PTGUtimezone.GlobalTimezoneSetup(os.Getenv("TZ")); err != nil {
		log.Fatalf(err.Error())
	}
}