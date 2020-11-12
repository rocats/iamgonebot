package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

var commands = []tgbotapi.BotCommand{
	{
		Command:     "unpin",
		Description: "unpin all iamgone messages",
	},
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	err = bot.SetMyCommands(commands)
	if err != nil {
		panic(err)
	}
	for update := range updates {
		if update.Message == nil {
			// ignore any non-Message Updates
			continue
		}
		iamgone(bot, update)
	}
}
