package scheduleService

import (
	"fmt"
	"leetcodebot/api/bot"
	"leetcodebot/api/groupsService"
	"leetcodebot/api/leetcodeApi"
	"leetcodebot/data/messages"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMorning() {
	groups := groupsService.GetApi()

	daily := leetcodeApi.GetDaily().ActiveDailyCodingChallengeQuestion

	groupsList, _ := groups.GetAll()
	for _, elem := range groupsList {
		if !elem.Enabled {
			continue
		}
		msg := tgbotapi.NewMessage(int64(elem.Id),fmt.Sprintf(messages.Morning, daily.Question.Title, daily.Question.Difficulty, leetcodeApi.Url + daily.Link))
		if _, err := bot.Bot.Send(msg); err != nil {
			log.Print(err)
			err = groups.Delete(elem.Id)
			if err != nil {
				log.Print(err)
			}
		}
	}
}