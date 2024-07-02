package commandsService

import (
	"fmt"
	"leetcodebot/api/bot"
	"leetcodebot/api/groupsService"
	"leetcodebot/data/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commands *Commands) GroupsList() error {
	msg := tgbotapi.NewMessage(commands.update.Message.Chat.ID, "")
	if commands.update.Message.From.ID == 1778923867 {
		var groupsApi = groupsService.GetApi()
		res, err := groupsApi.GetAll()
		if err != nil {
			return err
		}
		msg.Text = fmt.Sprintf(messages.GroupsList, res)
	} else {
		msg.Text = "Вы не имеете прав на выполнение этой команды"
	}
	if _, err := bot.Bot.Send(msg); err != nil {
		return err
	}
	return nil
}