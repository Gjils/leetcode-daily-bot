package main

import (
	commands "leetcodebot/app/commands"
	inline "leetcodebot/app/inline"
	schedule "leetcodebot/app/schedule"
	updateDB "leetcodebot/app/updateDB"

	bot "leetcodebot/api/bot"
)
func main() {

	go schedule.Start()

	for update := range bot.Updates {
		commands.Hanlde(update)
		updateDB.Hanlde(update)
		inline.Hanlde(update)
	}

	select {}
}