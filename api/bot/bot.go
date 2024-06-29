package bot

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)


func initBot() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	if err := godotenv.Load(".local.env", ".env"); err != nil {
		log.Print("No .env file found")	
	}
	
	token, exists := os.LookupEnv("BOT_TOKEN")
	if !exists {
		panic("token not found")
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	return bot, updates
}

var Bot, Updates = initBot()