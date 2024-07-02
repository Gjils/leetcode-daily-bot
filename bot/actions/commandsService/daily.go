package commandsService

import (
	"fmt"
	"leetcodebot/api/bot"
	"leetcodebot/api/leetcodeApi"
	"leetcodebot/data/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commands *Commands) Daily() error {
	msg := tgbotapi.NewMessage(commands.update.Message.Chat.ID, "")
	daily := leetcodeApi.GetDaily().ActiveDailyCodingChallengeQuestion
	msg.Text = fmt.Sprintf(messages.Daily, daily.Question.Title, daily.Question.Difficulty, leetcodeApi.Url + daily.Link)
	if _, err := bot.Bot.Send(msg); err != nil {
		return err
	}
	return nil
}