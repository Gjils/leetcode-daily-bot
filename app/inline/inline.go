package inline

import (
	"fmt"
	bot "leetcodebot/api/bot"
	leetcodeApi "leetcodebot/api/leetcodeApi"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const dailyText = `Сегодняшний дейлик: %v
Сложность: %v
			
%v`

func Hanlde(update tgbotapi.Update) error {
	if update.InlineQuery == nil {
		return nil
	}

	daily := leetcodeApi.GetDaily().ActiveDailyCodingChallengeQuestion
	article := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, "Дейлик", fmt.Sprintf(dailyText, daily.Question.Title, daily.Question.Difficulty, leetcodeApi.Url + daily.Link))
	article.Description = "Отправляет информацию о сегодняшнем дейлике"

	inlineConf := tgbotapi.InlineConfig{
			InlineQueryID: update.InlineQuery.ID,
			Results:       []interface{}{article},
			CacheTime:     1, // кэширование на 1 секунду
	}

	if _, err := bot.Bot.Request(inlineConf); err != nil {
			return err
	}
	return nil
}