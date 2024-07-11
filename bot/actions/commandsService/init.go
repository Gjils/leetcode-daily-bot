package commandsService

import (
	"fmt"
	"leetcodebot/api/bot"
	"leetcodebot/api/groupsService"
	"leetcodebot/data/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func (commands *Commands) Init() error {
	groups := groupsService.GetApi()
	message := "Группа уже есть в рассылке"
	inDB, err := groups.Check(int(commands.update.FromChat().ID))
	if err != nil {
		return err
	}
	if !inDB {
		err = groups.Add(int(commands.update.FromChat().ID))
		if err != nil {
			return err
		}
		message = fmt.Sprintf(messages.Init, commands.update.FromChat().ID)
	}
	msg := tgbotapi.NewMessage(commands.update.Message.Chat.ID, message)
	if _, err := bot.Bot.Send(msg); err != nil {
		return err
	}
	return nil
}