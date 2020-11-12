package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func handleErr(resp tgbotapi.APIResponse,err error)  {
	if err != nil {
		log.Println(err)
	} else if !resp.Ok {
		log.Println(resp.Description)
	}
}
