package commandsService

import (
	"fmt"
	"leetcodebot/api/bot"
	"leetcodebot/api/groupsService"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func (commands *Commands) Stop() error {
	groups := groupsService.GetApi()
	message := "Группы нет в рассылке"
	inDB, err := groups.Check(int(commands.update.FromChat().ID))
	if err != nil {
		return err
	}
	if inDB {
		err = groups.Delete(int(commands.update.FromChat().ID))
		if err != nil {
			return err
		}
		message = fmt.Sprintf("Чат %v удален из рассылки", commands.update.FromChat().ID)
	}
	msg := tgbotapi.NewMessage(commands.update.Message.Chat.ID, message)
	if _, err := bot.Bot.Send(msg); err != nil {
		return err
	}
	return nil
}