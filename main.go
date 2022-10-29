package main

import (
	"fmt"
	"vote-for-a-language/database"
	"vote-for-a-language/utils"

	"github.com/andersfylling/disgord"
)

func main() {
	client := disgord.New(disgord.Config{
		BotToken: utils.GetEnv("BOT_TOKEN"),
		Intents:  disgord.AllIntents(),
	})
	defer client.Gateway().StayConnectedUntilInterrupted()

	database.Connect()

	client.Gateway().BotReady(func() {
		fmt.Println("Bot is ready to Go!")
	})
}
