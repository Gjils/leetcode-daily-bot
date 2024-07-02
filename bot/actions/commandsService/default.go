package commandsService

import (
	"leetcodebot/api/bot"
	"leetcodebot/data/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commands *Commands) Default() error {
	msg := tgbotapi.NewMessage(commands.update.Message.Chat.ID, messages.Default)
	if _, err := bot.Bot.Send(msg); err != nil {
		return err
	}
	return nil
}