package helper

import (
	"golang-api-contact/config"
	"log"
	"time"
)

var (
	appTimezone *time.Location
)

func init () {
	timezoneStr := config.GetEnv("APP_TIMEZONE", "Asia/Jakarta")

	var err error

	appTimezone, err = time.LoadLocation(timezoneStr)
	if err != nil {
		log.Fatalf("Failed to get timezone '%s': %v", timezoneStr, err)
	}
}

func FormatTimeHuman(t time.Time) string  {
	return t.In(appTimezone).Format("2009-08-07 09:08:07")	
}