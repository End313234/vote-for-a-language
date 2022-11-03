package commands

import (
	"context"
	"fmt"
	"time"
	"vote-for-a-language/database"
	"vote-for-a-language/database/models"
	"vote-for-a-language/extensions/components"
	"vote-for-a-language/utils"

	"github.com/andersfylling/disgord"
)

var languageName string

var AddLanguageData = utils.SlashCommandData{
	Name:        "add_language",
	Description: "Performs a request to add a new language to the voting section",
	Options: []*disgord.ApplicationCommandOption{
		{
			Name:        "language",
			Description: "The name of the language",
			Type:        disgord.OptionTypeString,
			Required:    true,
		},
	},
	RateLimit: 10000,
}

func AddLanguageHandler(session disgord.Session, interaction *disgord.InteractionCreate) {
	languageName = interaction.Data.Options[0].Value.(string)

	foundLanguage := models.Language{}
	database.Client.Where("name = ?", languageName).Find(&foundLanguage)

	if foundLanguage.Name != "" {
		interaction.Reply(context.Background(), session, &disgord.CreateInteractionResponse{
			Type: disgord.InteractionCallbackChannelMessageWithSource,
			Data: &disgord.CreateInteractionResponseData{
				Flags: disgord.MessageFlagEphemeral,
				Embeds: []*disgord.Embed{
					{
						Description: "This language already exists!",
						Color:       0xFB1D2C,
					},
				},
			},
		})

		return
	}

	buttons := []*disgord.MessageComponent{
		{
			Type:     disgord.MessageComponentButton,
			Label:    "Add",
			Style:    disgord.Success,
			CustomID: "add-language",
		},
	}

	components.Buttons.Add(utils.Button{
		Data:    *buttons[0],
		Handler: addLanguageButtonHandler,
	})

	requestsChannelId := disgord.ParseSnowflakeString(utils.GetEnv("REQUESTS_CHANNEL_ID"))
	session.Channel(requestsChannelId).CreateMessage(&disgord.CreateMessage{
		Embeds: []*disgord.Embed{
			{
				Title: "New language request",
				Fields: []*disgord.EmbedField{
					{
						Name:  "Name",
						Value: languageName,
					},
				},
			},
		},
		Components: []*disgord.MessageComponent{
			{
				Type: disgord.MessageComponentActionRow,
				Components: []*disgord.MessageComponent{
					{
						Type:     disgord.MessageComponentButton,
						Label:    "Add",
						Style:    disgord.Success,
						CustomID: "add-language",
					},
				},
			},
		},
	})

	interaction.Reply(context.Background(), session, &disgord.CreateInteractionResponse{
		Type: disgord.InteractionCallbackChannelMessageWithSource,
		Data: &disgord.CreateInteractionResponseData{
			Flags: disgord.MessageFlagEphemeral,
			Embeds: []*disgord.Embed{
				{
					Description: "Your request has been performed successfully!",
					Color:       0x40FB6F,
				},
			},
		},
	})
}

func addLanguageButtonHandler(session disgord.Session, interaction *disgord.InteractionCreate) {
	interaction.Reply(context.Background(), session, &disgord.CreateInteractionResponse{
		Type: disgord.InteractionCallbackDeferredUpdateMessage,
	})

	interaction.Edit(context.Background(), session, &disgord.Message{
		Embeds: []*disgord.Embed{
			{
				Description: "Send the emoji ID in the chat",
				Color:       0xFBE24B,
			},
		},
		Components: []*disgord.MessageComponent{},
	})

	session.Gateway().WithCtrl(&disgord.Ctrl{
		Runs:     1,
		Duration: time.Minute,
	}).MessageCreate(func(s disgord.Session, h *disgord.MessageCreate) {
		msg := h.Message

		if msg.Author.ID != interaction.Member.UserID || msg.ChannelID != interaction.ChannelID {
			return
		}

		database.Client.Create(&models.Language{
			Name:    languageName,
			Votes:   0,
			EmojiId: msg.Content,
		})

		requestsChannelId := disgord.ParseSnowflakeString(utils.GetEnv("REQUESTS_CHANNEL_ID"))

		session.Channel(requestsChannelId).Message(msg.ID).Update(&disgord.UpdateMessage{
			Embeds: &[]*disgord.Embed{
				{
					Description: fmt.Sprintf("The language **%s** has been added successfully!", languageName),
					Color:       0x40FB6F,
				},
			},
		})
	})
}
