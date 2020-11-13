package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

const SLD = "哦噢喔耶啊哇呀哎哟阿啊呃欸哇呀也耶哟欤呕噢呦嘢吧罢呗啵的价家啦来唻嘞哩咧咯啰喽吗嘛嚜么麽哪呢呐否呵哈不兮般则连罗给噻哉呸了"

func iamgone(bot *tgbotapi.BotAPI, update tgbotapi.Update) (ok bool) {
	if t := strings.TrimRight(update.Message.Text, strings.ReplaceAll(SLD, "了", ""));
		strings.HasPrefix(t, "我") &&
			strings.HasSuffix(t, "了") {
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
