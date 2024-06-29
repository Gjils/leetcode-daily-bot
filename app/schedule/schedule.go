package schedule

import (
	"fmt"
	"leetcodebot/api/bot"
	"leetcodebot/api/groupsService"
	"leetcodebot/api/leetcodeApi"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
)

const morningMessage = `Доброе утро!
Сегодняшняя задача: %v
Сложность: %v

Удачного решения!
%v
`

const eveningMessage = `
День скоро кончится. Не забудь решить дейлик!

%v
`

func Start() {
	groupsService := groupsService.GetApi()
	c := cron.New()

	c.Start()

	var err error

	_, err = c.AddFunc("0 9 * * * ", func() {
		idList, _ := groupsService.GetAll()
		for _, elem := range idList {
			daily := leetcodeApi.GetDaily().ActiveDailyCodingChallengeQuestion
			msg := tgbotapi.NewMessage(int64(elem), fmt.Sprintf(morningMessage, daily.Question.Title, daily.Question.Difficulty, leetcodeApi.Url + daily.Link))
			if _, err := bot.Bot.Send(msg); err != nil {
				log.Print(err)
				err = groupsService.Delete(elem)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.AddFunc("0 20 * * * ", func() {
		idList, _ := groupsService.GetAll()
		for _, elem := range idList {
			daily := leetcodeApi.GetDaily().ActiveDailyCodingChallengeQuestion
			msg := tgbotapi.NewMessage(int64(elem), fmt.Sprintf(eveningMessage, leetcodeApi.Url + daily.Link))
			if _, err := bot.Bot.Send(msg); err != nil {
				log.Print(err)
				err = groupsService.Delete(elem)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
