package inlineService

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Inline struct {
	update tgbotapi.Update
}

func GetInline(update tgbotapi.Update) Inline {
	return Inline{update: update}
}