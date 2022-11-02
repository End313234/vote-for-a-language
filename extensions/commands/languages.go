package commands

import (
	"context"
	"fmt"
	"vote-for-a-language/database"
	"vote-for-a-language/database/models"
	"vote-for-a-language/utils"

	"github.com/andersfylling/disgord"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var LanguagesData = utils.SlashCommandData{
	Name:        "languages",
	Description: "Shows all the available languages for voting",
	RateLimit:   10000,
}

func LanguagesHandler(session disgord.Session, interaction *disgord.InteractionCreate) {
	languages := []models.Language{}
	database.Client.Find(&languages)

	embedDescription := ""
	for _, lang := range languages {
		embedDescription += fmt.Sprintf("**%s**\n", cases.Title(language.AmericanEnglish).String(lang.Name))
	}

	interaction.Reply(context.Background(), session, &disgord.CreateInteractionResponse{
		Type: disgord.InteractionCallbackChannelMessageWithSource,
		Data: &disgord.CreateInteractionResponseData{
			Embeds: []*disgord.Embed{
				{
					Title:       "Available languages",
					Description: embedDescription,
					Color:       0x40FB6F,
				},
			},
		},
	})
}
