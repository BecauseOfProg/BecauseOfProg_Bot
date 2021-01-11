package main

import (
	"errors"
	"fmt"
	"strconv"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

// HandleInlineQuery handles an inline query from a user (here: suggest publications to send into the channel)
func HandleInlineQuery(bot *telegram.BotAPI, update telegram.Update) error {
	result, err := SearchArticles(update.InlineQuery.Query)
	if err != nil {
		return errors.New(red.Sprintf("‼ Error calling the BecauseOfProg API: %s", err))
	}

	var results []interface{}
	for _, publication := range result.Data {
		url := fmt.Sprintf("https://becauseofprog.fr/publication/%s", publication.URL)
		results = append(results, telegram.InlineQueryResultArticle{
			Type: "publication",
			InputMessageContent: telegram.InputTextMessageContent{
				Text:      fmt.Sprintf("[%s](%s)", publication.Title, url),
				ParseMode: "Markdown",
			},
			ID:          strconv.Itoa(publication.Timestamp),
			Title:       publication.Title,
			Description: fmt.Sprintf("%s - %s", publication.Author.Name, publication.Description),
			URL:         url,
			ThumbURL:    publication.Banner,
		})
	}

	_, err = bot.AnswerInlineQuery(telegram.InlineConfig{
		InlineQueryID: update.InlineQuery.ID,
		Results:       results,
	})
	if err != nil {
		err = errors.New(red.Sprintf("‼ Error sending inline query result: %s", err))
	}

	return err
}
