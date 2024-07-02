package schedule

import (
	"leetcodebot/actions/scheduleService"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func Start() {
	var err error

	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal(err)
	}

	c := cron.New(cron.WithLocation(location))

	c.Start()
	
	_, err = c.AddFunc("0 9 * * * ", scheduleService.SendMorning)
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.AddFunc("0 20 * * * ", scheduleService.SendEvening)
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
