package handlers

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func (h Handlers) StartSchedule() {
	var err error

	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal(err)
	}

	c := cron.New(cron.WithLocation(location))

	c.Start()
	
	_, err = c.AddFunc("0 9 * * * ", func() {
		daily, err := h.services.ProblemsService.GetMorningInfo()
		if err != nil {
			log.Println(err)
			return
		}
		groupList, err := h.services.ManageService.GetAllGroups()
		if err != nil {
			log.Println(err)
			return
		}
		err = h.sendToAll(&groupList, daily)
		if err != nil {
			log.Println(err)
			return
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.AddFunc("0 20 * * * ", func() {
		daily, err := h.services.ProblemsService.GetEveningInfo()
		if err != nil {
			log.Println(err)
			return
		}
		groupList, err := h.services.ManageService.GetAllGroups()
		if err != nil {
			log.Println(err)
			return
		}
		err = h.sendToAll(&groupList, daily)
		if err != nil {
			log.Println(err)
			return
		}
	})
	if err != nil {
		log.Fatal(err)
	}

}
