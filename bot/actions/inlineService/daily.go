package inlineService

import (
	"fmt"
	"leetcodebot/api/leetcodeApi"
	"leetcodebot/data/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (inline Inline) Daily() interface{} {
	daily := leetcodeApi.GetDaily().ActiveDailyCodingChallengeQuestion
	article := tgbotapi.NewInlineQueryResultArticle(inline.update.InlineQuery.ID, "Дейлик", fmt.Sprintf(messages.Daily, daily.Question.Title, daily.Question.Difficulty, leetcodeApi.Url + daily.Link))
	article.Description = "Отправляет информацию о сегодняшнем дейлике"

	return article
}