package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

func iamgone(bot *tgbotapi.BotAPI, update tgbotapi.Update) (ok bool) {
	if strings.HasPrefix(update.Message.Text, "我") &&
		strings.HasSuffix(update.Message.Text, "了") {
		handleErr(bot.PinChatMessage(tgbotapi.PinChatMessageConfig{
			ChatID:              update.Message.Chat.ID,
			MessageID:           update.Message.MessageID,
			DisableNotification: true,
		}))
		return true
	}
	if cmd := update.Message.CommandWithAt(); cmd != "" {
		if cmd == "unpin" {
			//handleErr(bot.UnpinChatMessage(tgbotapi.UnpinChatMessageConfig{
			//	ChatID: update.Message.Chat.ID,
			//}))
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "TODO!\nhttps://github.com/go-telegram-bot-api/telegram-bot-api/issues/390"))
		}
	}
	return
}
