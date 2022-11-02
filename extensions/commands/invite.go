package commands

import (
	"context"
	"fmt"
	"vote-for-a-language/utils"

	"github.com/andersfylling/disgord"
)

var InviteData = utils.SlashCommandData{
	Name:        "invite",
	Description: "Returns the invite of the bot",
	Type:        disgord.ApplicationCommandChatInput,
	RateLimit:   5000,
}

func InviteHandler(session disgord.Session, interaction *disgord.InteractionCreate) {
	interaction.Reply(context.Background(), session, &disgord.CreateInteractionResponse{
		Type: disgord.InteractionCallbackChannelMessageWithSource,
		Data: &disgord.CreateInteractionResponseData{
			Embeds: []*disgord.Embed{
				{
					Description: fmt.Sprintf("Invite me using [this link](%s)!", utils.GetEnv("INVITE_LINK")),
					Color:       0x40FB6F,
				},
			},
		},
	})
}
