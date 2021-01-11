# BecauseOfProg_Bot

**Telegram bot to interact with the BecauseOfProg API**

[t.me/BecauseOfProg_Bot](https://t.me/BecauseOfProg_Bot)

## 🌈 Features

A simple goal: search for publications using an inline query.

## 💻 Development

Make sure you have Git and Go 1.15 installed on your machine.

Clone the repository locally:

```bash
git clone https://github.com:BecauseOfProg/BecauseOfProg_Bot.git  # HTTPS
git clone git@github.com:BecauseOfProg/BecauseOfProg_Bot          # SSH
```

Install the dependencies using go modules with the `go get` command. You can also use the dependency manager of your choice. Every dependency is listed in the [go.mod](./go.mod) file.

Initialize these environment variables:

- `TELEGRAM_APITOKEN` - the token of your bot (get one with [@BotFather](https://t.me/BotFather))
- `BOT_ENV` - Set it to "development" to get detailed debug logs

You can also create a `.env` file at the root of the directory.

Finally, start the bot with `go run` or build the project with `go build`.

## 📜 Credits

- Maintainer: [Théo Vidal](https://github.com/theovidal)
- Library: [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api)

## 🔐 License

This code is protected under the [GNU GPL v3](./LICENSE) license.
