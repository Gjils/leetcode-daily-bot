package inlineHandler

import (
	"leetcodebot/actions/inlineService"
	bot "leetcodebot/api/bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func Hanlde(update tgbotapi.Update) error {
	if update.InlineQuery == nil {
		return nil
	}

	inline := inlineService.GetInline(update)

	inlineConf := tgbotapi.InlineConfig{
			InlineQueryID: update.InlineQuery.ID,
			Results:       []interface{}{inline.Daily()},
			CacheTime:     1,
	}

	if _, err := bot.Bot.Request(inlineConf); err != nil {
			return err
	}
	return nil
}