package commands

import (
	"fmt"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"

	"github.com/BecauseOfProg/BecauseOfProg_Bot/data"
	"github.com/BecauseOfProg/BecauseOfProg_Bot/lib"
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
		categoryID := args[0]
		category, exists := data.Categories[categoryID]
		if !exists {
			_, err := bot.Send(telegram.NewMessage(chatID, "❓ Oups, cette catégorie est inconnue. Faites /categories pour obtenir la liste!"))
			return err
		}

		page := 1
		if len(args) > 1 {
			var err error
			page, err = strconv.Atoi(args[1])
			if err != nil {
				_, err := bot.Send(telegram.NewMessage(chatID, "❌ Merci de préciser un nombre entier pour le numéro de page"))
				return err
			}
		}

		result, err := lib.GetPublicationsByCategory(categoryID, page)
		if err != nil {
			return err
		}

		if page > result.Pages || page < 1 {
			_, err := bot.Send(telegram.NewMessage(chatID, fmt.Sprintf("❌ Merci de préciser un numéro de page entre 1 et %d", result.Pages)))
			return err
		}

		var publications []interface{}
		for _, publication := range result.Data {
			photo := telegram.NewInputMediaPhoto(publication.Banner)
			_, link := publication.FormatLink()
			photo.Caption = fmt.Sprintf("%s\n%s - %s", link, publication.Author.Name, publication.Description)
			photo.ParseMode = "Markdown"
			publications = append(publications, photo)
		}

		var controlKeyboard []telegram.InlineKeyboardButton
		if page != 1 {
			controlKeyboard = append(controlKeyboard, telegram.NewInlineKeyboardButtonData("⏪ Précédent", fmt.Sprintf("/categories %s %d", category.ID, page-1)))
		}
		controlKeyboard = append(controlKeyboard, telegram.NewInlineKeyboardButtonURL("Page", "https://becauseofprog.fr/categorie/"+category.ID))
		if page != result.Pages {
			controlKeyboard = append(controlKeyboard, telegram.NewInlineKeyboardButtonData("Suivant ⏩", fmt.Sprintf("/categories %s %d", category.ID, page+1)))
		}

		msg := telegram.NewMessage(chatID, category.Name+" | Toutes les publications")
		msg.ReplyMarkup = telegram.NewInlineKeyboardMarkup(controlKeyboard)

		bot.Send(msg)
		_, err = bot.Send(telegram.NewMediaGroup(chatID, publications))

		return err
	}

	msg := telegram.NewMessage(chatID, "📰 Choisissez la catégorie que vous souhaitez consulter")
	msg.ReplyMarkup = categoriesKeyboard
	_, err := bot.Send(msg)

	return err
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
