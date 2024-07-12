package config

import (
	"errors"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func GetTelegramClient() (*tele.Bot, error) {
	token, exists := os.LookupEnv("BOT_TOKEN")
	if !exists {
		return nil, errors.New("token not found")
	}
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return b, nil
}