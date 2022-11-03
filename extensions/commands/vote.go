package commands

import (
	"context"
	"vote-for-a-language/constants"
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
	RateLimit: 5000,
}

func VoteHandler(session disgord.Session, interaction *disgord.InteractionCreate) {
	languageName = interaction.Data.Options[0].Value.(string)
	foundLanguage := models.Language{}
	database.Client.Where("name = ?", constants.EnglishLowerCase.String(languageName)).Find(&foundLanguage)

	if foundLanguage.Name == "" {
		interaction.Reply(context.Background(), session, &disgord.CreateInteractionResponse{
			Type: disgord.InteractionCallbackChannelMessageWithSource,
			Data: &disgord.CreateInteractionResponseData{
				Flags: disgord.MessageFlagEphemeral,
				Embeds: []*disgord.Embed{
					{
						Description: "This language does not exist!",
						Color:       constants.RED_COLOR,
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
					Color:       constants.GREEN_COLOR,
				},
			},
		},
	})
}
