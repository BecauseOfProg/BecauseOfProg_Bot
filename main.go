package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"

	"github.com/BecauseOfProg/BecauseOfProg_Bot/lib"
)

func main() {
	fmt.Println(color.HiRedString(" ________  _______   ________  ________  ___  ___  ________  _______   ________  ________ ________  ________  ________  ________     \n|\\   __  \\|\\  ___ \\ |\\   ____\\|\\   __  \\|\\  \\|\\  \\|\\   ____\\|\\  ___ \\ |\\   __  \\|\\  _____\\\\   __  \\|\\   __  \\|\\   __  \\|\\   ____\\    \n\\ \\  \\|\\ /\\ \\   __/|\\ \\  \\___|\\ \\  \\|\\  \\ \\  \\\\\\  \\ \\  \\___|\\ \\   __/|\\ \\  \\|\\  \\ \\  \\__/\\ \\  \\|\\  \\ \\  \\|\\  \\ \\  \\|\\  \\ \\  \\___|    \n \\ \\   __  \\ \\  \\_|/_\\ \\  \\    \\ \\   __  \\ \\  \\\\\\  \\ \\_____  \\ \\  \\_|/_\\ \\  \\\\\\  \\ \\   __\\\\ \\   ____\\ \\   _  _\\ \\  \\\\\\  \\ \\  \\  ___  \n  \\ \\  \\|\\  \\ \\  \\_|\\ \\ \\  \\____\\ \\  \\ \\  \\ \\  \\\\\\  \\|____|\\  \\ \\  \\_|\\ \\ \\  \\\\\\  \\ \\  \\_| \\ \\  \\___|\\ \\  \\\\  \\\\ \\  \\\\\\  \\ \\  \\|\\  \\ \n   \\ \\_______\\ \\_______\\ \\_______\\ \\__\\ \\__\\ \\_______\\____\\_\\  \\ \\_______\\ \\_______\\ \\__\\   \\ \\__\\    \\ \\__\\\\ _\\\\ \\_______\\ \\_______\\\n    \\|_______|\\|_______|\\|_______|\\|__|\\|__|\\|_______|\\_________\\|_______|\\|_______|\\|__|    \\|__|     \\|__|\\|__|\\|_______|\\|_______|\n                                                     \\|_________|                                                                    "))

	if err := godotenv.Load(); err != nil {
		log.Println("💾 No .env file at the root - Ignoring")
	}

	bot, err := telegram.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Println(lib.Red.Sprintf("‼ Error creating Telegram bot: %s", err))
		return
	}

	if os.Getenv("BOT_ENV") == "development" {
		bot.Debug = true
	}

	log.Println(lib.Green.Sprintf("✅ Authorized on account %s", bot.Self.UserName))

	u := telegram.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		var err error
		if update.InlineQuery != nil {
			err = HandleInlineQuery(bot, update)
		} else if update.CallbackQuery != nil {
			if strings.HasPrefix(update.CallbackQuery.Data, "/") {
				err = HandleCommand(bot, update, true)
			}
		} else if update.Message.IsCommand() {
			err = HandleCommand(bot, update, false)
		}

		if err != nil {
			log.Println(lib.Red.Sprintf("‼ Error handling an event: %s", err))
		}
	}
}
