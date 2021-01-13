package commands

import (
	"fmt"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

var StartKeyboard = telegram.NewInlineKeyboardMarkup(
	telegram.NewInlineKeyboardRow(
		telegram.NewInlineKeyboardButtonData("🗂 Catégories", "/categories"),
		// telegram.NewInlineKeyboardButtonData("🔎 Rechercher", "/search"),
	),
)

func StartCommand(bot *telegram.BotAPI, update telegram.Update, _ []string) (err error) {
	fmt.Println(update.Message.From.LanguageCode)
	msg := telegram.NewMessage(update.Message.Chat.ID, "🤖 Bienvenue sur le bot BecauseOfProg! Choisissez une action pour démarrer :")
	msg.ReplyMarkup = StartKeyboard

	_, err = bot.Send(msg)
	return
}
