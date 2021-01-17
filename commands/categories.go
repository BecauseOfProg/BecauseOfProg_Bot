package commands

import (
	"fmt"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/BecauseOfProg/BecauseOfProg_Bot/data"
)

var categoriesKeyboard = generateCategoriesKeyboard()

func CategoriesCommand(bot *telegram.BotAPI, update telegram.Update, args []string) error {
	var chatID int64
	if update.CallbackQuery == nil {
		chatID = update.Message.Chat.ID
	} else {
		chatID = update.CallbackQuery.Message.Chat.ID
	}

	if len(args) != 0 {
		return Category(bot, &update, chatID, args)
	}

	msg := telegram.NewMessage(chatID, "📰 Choisissez la catégorie que vous souhaitez consulter")
	msg.ReplyMarkup = categoriesKeyboard
	_, err := bot.Send(msg)

	return err
}

func parseCallback(category string, page int) string {
	return fmt.Sprintf("/categories %s %d", category, page)
}

func generateCategoriesKeyboard() telegram.InlineKeyboardMarkup {
	var buttons [][]telegram.InlineKeyboardButton
	for _, category := range data.Categories {
		buttons = append(
			buttons,
			[]telegram.InlineKeyboardButton{
				telegram.NewInlineKeyboardButtonData(category.Name, fmt.Sprintf("/categories %s 1", category.ID)),
			})
	}

	return telegram.NewInlineKeyboardMarkup(buttons...)
}
