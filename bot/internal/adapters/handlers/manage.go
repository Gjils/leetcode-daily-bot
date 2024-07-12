package handlers

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

func (h Handlers) startManage() {
	h.bot.Handle("/leetcodebotInit", func(c tele.Context) error {
		fmt.Println(c.Message().Sender.ID)
		res, err := h.services.ManageService.AddGroup(getGroupInfo(c))
		if err != nil {
			return err
		}
		return c.Send(res, getSendOptions(c.Message().ThreadID))
	})
	h.bot.Handle("/leetcodebotStatus", func(c tele.Context) error {
		res, err := h.services.ManageService.GetGroupStatus(getGroupInfo(c))
		if err != nil {
			return err
		}
		return c.Send(res, getSendOptions(c.Message().ThreadID))
	})
	h.bot.Handle("/leetcodebotStop", func(c tele.Context) error {
		res, err := h.services.ManageService.RemoveGroup(getGroupInfo(c))
		if err != nil {
			return err
		}
		return c.Send(res, getSendOptions(c.Message().ThreadID))
	})
	h.bot.Handle("/groupInfo", func(c tele.Context) error {
		return c.Send(fmt.Sprintf("%+v", getGroupInfo(c)), getSendOptions(c.Message().ThreadID))
	})
	h.bot.Handle("/sendAll", func(c tele.Context) error {
		daily, err := h.services.ProblemsService.GetDailyInfo()
		if err != nil {
			return err
		}
		groupList, err := h.services.ManageService.GetAllGroups()
		if err != nil {
			return err
		}
		return h.sendToAll(&groupList, daily)
	})
}