package main

import (
	"vote-for-a-language/database"
	"vote-for-a-language/extensions/events"
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

	client.Gateway().BotReady(events.BotReady(client))
	client.Gateway().InteractionCreate(events.InteractionCreate)
}
