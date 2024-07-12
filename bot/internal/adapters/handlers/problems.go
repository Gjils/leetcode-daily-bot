package handlers

import tele "gopkg.in/telebot.v3"

func (h Handlers) startProblems() {
	h.bot.Handle("/daily", func(c tele.Context) error {
		res, err := h.services.ProblemsService.GetDailyInfo()
		if err != nil {
			return err
		}
		return c.Send(res, getSendOptions(c.Message().ThreadID))
	})
}