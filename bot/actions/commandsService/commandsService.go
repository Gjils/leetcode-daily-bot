package commandsService

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Commands struct {
	update tgbotapi.Update
}

func GetCommands(update tgbotapi.Update) Commands {
	return Commands{update: update}
}