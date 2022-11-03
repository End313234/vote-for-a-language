package events

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"vote-for-a-language/config"
	"vote-for-a-language/constants"
	"vote-for-a-language/extensions/commands"
	"vote-for-a-language/extensions/components"

	"github.com/andersfylling/disgord"
)

func InteractionCreate(session disgord.Session, interaction *disgord.InteractionCreate) {
	switch interaction.Type {
	case disgord.InteractionApplicationCommand:
		commandName := interaction.Data.Name
		userId := interaction.Member.UserID.String()
		rateLimits := commands.SlashCommands.CollectRateLimits()

		lastTimeCommandWasUsed := config.Cache.Users[userId][commandName]
		lastTimeCommandWasUsedInMilliseconds := time.UnixMilli(int64(lastTimeCommandWasUsed))
		timeSinceTheLastCommandWasUsed := time.Since(lastTimeCommandWasUsedInMilliseconds)
		rateLimitDuration := time.Millisecond * time.Duration(rateLimits[commandName])

		if time.Duration(timeSinceTheLastCommandWasUsed.Milliseconds()) <= time.Duration(rateLimitDuration.Milliseconds()) {
			secondsToUseTheCommandAgain := (rateLimitDuration - timeSinceTheLastCommandWasUsed).Seconds()
			formattedWaitLine := strconv.FormatFloat(secondsToUseTheCommandAgain, 'f', 2, 32)

			interaction.Reply(context.Background(), session, &disgord.CreateInteractionResponse{
				Type: disgord.InteractionCallbackChannelMessageWithSource,
				Data: &disgord.CreateInteractionResponseData{
					Embeds: []*disgord.Embed{
						{
							Description: fmt.Sprintf("You have just been rate limited! Wait %ss to use this command again!", formattedWaitLine),
							Color:       constants.RED_COLOR,
						},
					},
				},
			})

			return
		}

		command := commands.SlashCommands.FindByName(commandName)
		command.Handler(session, interaction)

		config.Cache.Users.Set(userId, map[string]int64{
			commandName: time.Now().UnixMilli(),
		})
	case disgord.InteractionMessageComponent:
		components.Buttons.FindByCustomId(interaction.Data.CustomID).Handler(session, interaction)
	}
}
