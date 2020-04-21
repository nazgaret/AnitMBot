package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	MWords = []string{
		"манчкин",
		"мачкин",
		"манчикин",
		"мачикин",
		"монополия",
	}
	MWordsReplacer = NewMWordReplacer()
)

func main() {
	bot, err := tgbotapi.NewBotAPI("mykey")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	//addwebhook
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		var needRemove bool
		for _, word := range MWords {
			if strings.Contains(strings.ToLower(update.Message.Text), word) {
				needRemove = true
				break
			}
		}
		if needRemove {
			editedMsg := fmt.Sprintf("[%s] says:\n %s", update.Message.From.UserName, MWordsReplacer.Replace(strings.ToLower(update.Message.Text)))
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			del := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, editedMsg)
			bot.Send(del)
			bot.Send(msg)
		}
	}
}
func NewMWordReplacer() *strings.Replacer {
	newolds := make([]string, len(MWords)*2)
	for _, word := range MWords {
		newolds = append(newolds, word, "М-хуйня")
	}
	return strings.NewReplacer(newolds...)
}
