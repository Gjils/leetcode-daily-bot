package commands

import (
	"fmt"
	bot "leetcodebot/api/bot"
	leetcodeApi "leetcodebot/api/leetcodeApi"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const dailyText = `Сегодняшний дейлик: %v
Сложность: %v
			
%v`

const startMessage = `Привет! Этот бот рассылает дейлики литкода
Добавь его в группу и каждый день в 9 утра и 8 вечера будет напоминать о задаче`

const helpMessage = `Доступные команды:
/start - приветственное сообщение
/help - список доступных команд
/daily - информация о сегодняшнем дейлике`

const defaultMessage = `Команда не распознана
/help - список доступных команд`

func Hanlde(update tgbotapi.Update) error {
	if update.Message == nil {
		return nil
	}
	if !update.Message.IsCommand() {
		return nil
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Command() {
	case "start":
		msg.Text = startMessage
	case "help":
		msg.Text = helpMessage
	case "daily":
		daily := leetcodeApi.GetDaily().ActiveDailyCodingChallengeQuestion
		msg.Text = fmt.Sprintf(dailyText, daily.Question.Title, daily.Question.Difficulty, leetcodeApi.Url + daily.Link)
	default:
		msg.Text = defaultMessage
	}
	if _, err := bot.Bot.Send(msg); err != nil {
		return err
	}
	return nil
}