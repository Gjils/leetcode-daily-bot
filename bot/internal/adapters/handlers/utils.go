package handlers

import (
	"fmt"
	"leetcodebot/internal/domain/entities"
	"log"

	tele "gopkg.in/telebot.v3"
)

func getSendOptions(threadID int) *tele.SendOptions {
	return &tele.SendOptions{
		ReplyTo: nil,
		DisableWebPagePreview: false,
		DisableNotification: false,
		ParseMode: tele.ModeMarkdown,
		ThreadID: threadID,
	}
}

func getGroupInfo(c tele.Context) entities.GroupInfo {
	isThread := c.Chat().Type == "supergroup"
	return entities.GroupInfo{
		Title: c.Chat().Title,
		ChatId: int(c.Chat().ID),
		IsThread: isThread,
		ThreadId: c.Message().ThreadID,
	}
}

func (h Handlers) sendToAll(list *[]entities.Group, message string) error {
	for _, elem := range *list {
		fmt.Println(elem.ChatId)
		_, err := h.bot.Send(&tele.Chat{ID: int64(elem.ChatId)}, message, getSendOptions(elem.ThreadId))
		if err != nil {
			log.Print(err)
		}
	}
	return nil
}