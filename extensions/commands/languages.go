package commands

import (
	"context"
	"fmt"
	"math"
	"vote-for-a-language/constants"
	"vote-for-a-language/database"
	"vote-for-a-language/database/models"
	"vote-for-a-language/extensions/components"
	"vote-for-a-language/utils"

	"github.com/andersfylling/disgord"
)

var embeds []*disgord.Embed
var currentEmbed = 0
var quantityOfEmbeds int

var LanguagesData = utils.SlashCommandData{
	Name:        "languages",
	Description: "Shows all the available languages for voting",
	RateLimit:   10000,
}

func LanguagesHandler(session disgord.Session, interaction *disgord.InteractionCreate) {
	languages := []models.Language{}
	database.Client.Find(&languages)

	var biggerLanguageNameLength int
	for _, lang := range languages {
		if len(lang.Name) > biggerLanguageNameLength {
			biggerLanguageNameLength = len(lang.Name)
		}
	}

	quantityOfEmbeds = int(math.Ceil(float64(len(languages)) / 9.0))
	languagesInChunks := utils.ChunkBy(9, languages)

	for q := 0; q < quantityOfEmbeds; q++ {
		description := ""
		for _, lang := range languagesInChunks[q] {
			description += fmt.Sprintf("<:%s:%s> %s - **%d**\n", lang.Name, lang.EmojiId, constants.EnglishTitleCase.String(lang.Name), lang.Votes)
		}

		embeds = append(embeds, &disgord.Embed{
			Title:       "Available languages",
			Description: description,
			Color:       constants.GREEN_COLOR,
		})
	}

	messageComponents := []*disgord.MessageComponent{
		{
			Type:     disgord.MessageComponentButton,
			Label:    "⬅️",
			Style:    disgord.Secondary,
			CustomID: "go-back",
			Disabled: currentEmbed-1 < 0,
		},
		{
			Type:     disgord.MessageComponentButton,
			Label:    "➡️",
			Style:    disgord.Secondary,
			CustomID: "go-next",
			Disabled: currentEmbed+1 >= quantityOfEmbeds,
		},
	}

	messageButtons := []utils.Button{
		{
			Data:    *messageComponents[0],
			Handler: goBackHandler,
		},
		{
			Data:    *messageComponents[1],
			Handler: goNextHandler,
		},
	}

	components.Buttons.Add(messageButtons...)

	interaction.Reply(context.Background(), session, &disgord.CreateInteractionResponse{
		Type: disgord.InteractionCallbackChannelMessageWithSource,
		Data: &disgord.CreateInteractionResponseData{
			Embeds: []*disgord.Embed{
				embeds[currentEmbed],
			},
			Components: []*disgord.MessageComponent{
				{
					Type:       disgord.MessageComponentActionRow,
					Components: messageComponents,
				},
			},
		},
	})
}

func baseHandler(session disgord.Session, interaction *disgord.InteractionCreate, currentEmbed int) {
	messageComponents := []*disgord.MessageComponent{
		{
			Type:     disgord.MessageComponentButton,
			Label:    "⬅️",
			Style:    disgord.Secondary,
			CustomID: "go-back",
			Disabled: currentEmbed-1 < 0,
		},
		{
			Type:     disgord.MessageComponentButton,
			Label:    "➡️",
			Style:    disgord.Secondary,
			CustomID: "go-next",
			Disabled: currentEmbed+1 >= quantityOfEmbeds,
		},
	}

	interaction.Reply(context.Background(), session, &disgord.CreateInteractionResponse{
		Type: disgord.InteractionCallbackDeferredUpdateMessage,
	})

	interaction.Edit(context.Background(), session, &disgord.Message{
		Embeds: []*disgord.Embed{
			embeds[currentEmbed],
		},
		Components: []*disgord.MessageComponent{
			{
				Type:       disgord.MessageComponentActionRow,
				Components: messageComponents,
			},
		},
	})
}

func goBackHandler(session disgord.Session, interaction *disgord.InteractionCreate) {
	currentEmbed -= 1

	baseHandler(session, interaction, currentEmbed)
}

func goNextHandler(session disgord.Session, interaction *disgord.InteractionCreate) {
	currentEmbed += 1

	baseHandler(session, interaction, currentEmbed)
}
