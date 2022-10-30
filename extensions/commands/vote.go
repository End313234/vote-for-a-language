package commands

import (
	"context"
	"vote-for-a-language/database"
	"vote-for-a-language/database/models"
	"vote-for-a-language/utils"

	"github.com/andersfylling/disgord"
)

var VoteData = utils.SlashCommandData{
	Name:        "vote",
	Description: "Votes for a language",
	Options: []*disgord.ApplicationCommandOption{
		{
			Type:        disgord.OptionTypeString,
			Name:        "language",
			Description: "The language to vote",
			Required:    true,
		},
	},
}

func VoteHandler(session disgord.Session, interaction *disgord.InteractionCreate) {
	foundLanguage := models.Language{}
	database.Client.Where("name = ?", languageName).Find(&foundLanguage)

	if foundLanguage.Name == "" {
		interaction.Reply(context.Background(), session, &disgord.CreateInteractionResponse{
			Type: disgord.InteractionCallbackChannelMessageWithSource,
			Data: &disgord.CreateInteractionResponseData{
				Flags: disgord.MessageFlagEphemeral,
				Embeds: []*disgord.Embed{
					{
						Description: "This language does not exist!",
						Color:       0xFB1D2C,
					},
				},
			},
		})

		return
	}

	foundLanguage.Votes += 1
	database.Client.Save(&foundLanguage)

	interaction.Reply(context.Background(), session, &disgord.CreateInteractionResponse{
		Type: disgord.InteractionCallbackChannelMessageWithSource,
		Data: &disgord.CreateInteractionResponseData{
			Flags: disgord.MessageFlagEphemeral,
			Embeds: []*disgord.Embed{
				{
					Description: "Your vote has been registered successfully!",
					Color:       0x40FB6F,
				},
			},
		},
	})
}
