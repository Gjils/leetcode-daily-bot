package main

import (
	"leetcodebot/app/commandsHandler"
	"leetcodebot/app/inlineHandler"
	"leetcodebot/app/schedule"
	"leetcodebot/app/updateDB"

	bot "leetcodebot/api/bot"
)
func main() {

	go schedule.Start()

	for update := range bot.Updates {
		commandsHandler.Hanlde(update)
		updateDB.Hanlde(update)
		inlineHandler.Hanlde(update)
	}

	select {}
}