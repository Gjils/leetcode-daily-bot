package handlers

import (
	serviceInterfaces "leetcodebot/internal/domain/services/interfaces"

	tele "gopkg.in/telebot.v3"
)

type Handlers struct {
	services serviceInterfaces.Services
	bot *tele.Bot
}

func GetHandlers(services serviceInterfaces.Services, bot *tele.Bot) Handlers {
	return Handlers{services: services, bot: bot}
}

func (h Handlers) Start() {
	h.startBase()
	h.startManage()
	h.startProblems()
	h.StartSchedule()
}