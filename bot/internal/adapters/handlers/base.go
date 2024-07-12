package handlers

import tele "gopkg.in/telebot.v3"

const startMessage = `Привет! Этот бот рассылает дейлики литкода
Добавь его в группу, напиши /leetcodebotInit и каждый день в 9 утра и 8 вечера он будет напоминать о задаче`

const helpMessage = `Доступные команды:
/start - приветственное сообщение
/help - список доступных команд

/daily - информация о сегодняшнем дейлике

/leetcodebotInit - начать рассылку в этом чате
/leetcodebotStatus - статус чата
/leetcodebotStop - прекратить рассылку`

func (h Handlers) startBase() {
	h.bot.Handle("/start", func(c tele.Context) error {
		return c.Send(startMessage)
	})
	h.bot.Handle("/help", func(c tele.Context) error {
		return c.Send(helpMessage)
	})
}