package updateDB

import (
	groupsApi "leetcodebot/api/groupsService"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var groupsService = groupsApi.GetApi()


func Hanlde(update tgbotapi.Update) error {
	chat := update.FromChat()
	if chat == nil {
		return nil
	}
	if chat.Type == "group" || chat.Type == "supergroup" {
		inDB, err := groupsService.Check(int(chat.ID))
		if err != nil {
			return err
		}
		if !inDB {
			if err := groupsService.Add(int(chat.ID)); err != nil {
				return err
			}
		}
	}
	return nil
}