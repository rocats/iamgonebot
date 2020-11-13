package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
	"unicode"
)

const SLD = "哦噢喔耶啊哇呀哎哟阿啊呃欸哇呀也耶哟欤呕噢呦嘢吧罢呗啵的价家啦来唻嘞哩咧咯啰喽吗嘛嚜么麽哪呢呐否呵哈不兮般则连罗给噻哉呸了"

func consecutiveHanLetterMixedSpace(str string) bool {
	runes := []rune(str)
	preSpace := false
	lastPos := -1
	for i := range runes {
		if unicode.IsSpace(runes[i]) {
			if i > 0 && !preSpace {
				preSpace = true
				lastPos = i - 1
			}
		} else if preSpace {
			preSpace = false
			if unicode.Is(unicode.Han, runes[lastPos]) &&
				unicode.Is(unicode.Han, runes[i]) {
				return true
			} else if (unicode.Is(unicode.Han, runes[lastPos]) ||
				unicode.Is(unicode.Han, runes[i])) &&
				(unicode.Is(unicode.Letter, runes[lastPos]) ||
					unicode.Is(unicode.Letter, runes[i])) {
				return true
			}
		}
	}
	return false
}

func trim(str string) string {
	ban := strings.ReplaceAll(SLD, "了", "")
	str = strings.TrimRightFunc(str, func(r rune) bool {
		return !unicode.Is(unicode.Han, r) || strings.ContainsRune(ban, r)
	})
	str = strings.TrimLeftFunc(str, func(r rune) bool {
		return !unicode.Is(unicode.Han, r) || strings.ContainsRune(ban, r)
	})
	return str
}

func shouldPin(str string) bool {
	t := trim(str)
	return strings.HasPrefix(t, "我") &&
		strings.HasSuffix(t, "了") &&
		!strings.ContainsAny(t, "，,") &&
		!consecutiveHanLetterMixedSpace(t)
}

func iamgone(bot *tgbotapi.BotAPI, update tgbotapi.Update) (ok bool) {
	if shouldPin(update.Message.Text) {
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
