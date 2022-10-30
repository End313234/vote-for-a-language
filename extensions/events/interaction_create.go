package events

import (
	"vote-for-a-language/extensions/commands"
	"vote-for-a-language/extensions/components"

	"github.com/andersfylling/disgord"
)

func InteractionCreate(session disgord.Session, interaction *disgord.InteractionCreate) {
	switch interaction.Type {
	case disgord.InteractionApplicationCommand:
		commands.SlashCommands.FindByName(interaction.Data.Name).Handler(session, interaction)
	case disgord.InteractionMessageComponent:
		components.Buttons.FindByCustomId(interaction.Data.CustomID).Handler(session, interaction)
	}
}
