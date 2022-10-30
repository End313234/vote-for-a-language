package events

import (
	"fmt"
	"log"
	"vote-for-a-language/extensions/commands"
	"vote-for-a-language/utils"

	"github.com/andersfylling/disgord"
)

func BotReady(client *disgord.Client) func() {
	return func() {
		for id, command := range commands.SlashCommands {
			slashCommand := &disgord.CreateApplicationCommand{
				Name:              command.Data.Name,
				Description:       command.Data.Description,
				Type:              command.Data.Type,
				Options:           command.Data.Options,
				DefaultPermission: command.Data.DefaultPermission,
			}

			slashCommandId, _ := disgord.GetSnowflake(id)
			devGuildId, _ := disgord.GetSnowflake(utils.GetEnv("DEV_GUILD_ID"))
			environment := utils.GetEnv("ENVIRONMENT")

			applicationCommand := client.ApplicationCommand(slashCommandId)

			if environment == "development" {
				if err := applicationCommand.Guild(devGuildId).Create(slashCommand); err != nil {
					log.Fatal(err)
				}
			} else if environment == "production" {
				if err := applicationCommand.Global().Create(slashCommand); err != nil {
					log.Fatal(err)
				}
			}
		}

		fmt.Println("Bot is ready to Go!")
	}
}
