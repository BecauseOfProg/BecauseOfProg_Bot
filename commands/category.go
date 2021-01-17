package commands

import (
	"fmt"
	"strconv"
	"strings"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/BecauseOfProg/BecauseOfProg_Bot/data"
	"github.com/BecauseOfProg/BecauseOfProg_Bot/lib"
)

func Category(bot *telegram.BotAPI, update *telegram.Update, chatID int64, args []string) error {
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
	fmt.Println(page)

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
		publications = append(publications, parsePublication(publication))
	}

	var controlKeyboard []telegram.InlineKeyboardButton
	if page != 1 {
		controlKeyboard = append(
			controlKeyboard,
			telegram.NewInlineKeyboardButtonData("⏪ Précédent", parseCallback(category.ID, page-1)),
		)
	}
	controlKeyboard = append(
		controlKeyboard,
		telegram.NewInlineKeyboardButtonURL("Voir sur le blog", "https://becauseofprog.fr/categorie/"+category.ID),
	)
	if page != result.Pages {
		controlKeyboard = append(
			controlKeyboard,
			telegram.NewInlineKeyboardButtonData("Suivant ⏩", parseCallback(category.ID, page+1)),
		)
	}
	messageContent := fmt.Sprintf("%s | Page %d/%d", category.Name, page, result.Pages)
	markup := telegram.NewInlineKeyboardMarkup(controlKeyboard)

	if update.CallbackQuery != nil {
		if _, err = bot.AnswerCallbackQuery(telegram.NewCallback(update.CallbackQuery.ID, "")); err != nil {
			return err
		}
	}

	// To use with messages that can be edited (so, not media groups)
	/* if len(args) > 2 {
		messageID, err := strconv.Atoi(args[2])
		if err != nil {
			_, err := bot.Send(telegram.NewMessage(chatID, "❌ Merci de préciser un ID de message valide"))
			return err
		}
		edit := telegram.NewEditMessageText(chatID, messageID, messageContent)
		edit.ReplyMarkup = &markup
		_, err = bot.Send(edit)
		return err
	} */

	msg := telegram.NewMessage(chatID, messageContent)
	msg.ReplyMarkup = markup

	_, _ = bot.Send(msg)
	_, err = bot.Send(telegram.NewMediaGroup(chatID, publications))

	return err
}

func parsePublication(publication lib.Publication) (photo telegram.InputMediaPhoto) {
	banner := "https://i.cdn.becauseofprog.fr/" + strings.TrimPrefix(publication.Banner, "https://") + "?resize=1024,576"
	photo = telegram.NewInputMediaPhoto(banner)
	_, link := publication.FormatLink()
	photo.Caption = fmt.Sprintf("%s\n%s - %s", link, publication.Author.Name, publication.Description)
	photo.ParseMode = "Markdown"
	return
}
