package commandsHandler

import (
	"leetcodebot/actions/commandsService"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Hanlde(update tgbotapi.Update) error {
	if update.Message == nil {
		return nil
	}
	if !update.Message.IsCommand() {
		return nil
	}

	commands := commandsService.GetCommands(update)

	switch update.Message.Command() {
	case "start": commands.Start()
	case "help": commands.Help()
	case "daily": commands.Daily()
	case "groups_list": commands.GroupsList()
	case "leetcodeDailyInit": commands.Init()
	case "leetcodeDailyStop": commands.Stop()
	default: commands.Default()
	}
	return nil
}