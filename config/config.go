package config

import (
	"log"
	"os"

	"github.com/parinyapt/golang_utils/timezone/v1"
)

func GlobalSetup() {
	if err := PTGUtimezone.GlobalTimezoneSetup(os.Getenv("TZ")); err != nil {
		log.Fatalf(err.Error())
	}
}