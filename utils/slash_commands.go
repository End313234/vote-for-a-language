package utils

import "github.com/andersfylling/disgord"

type SlashCommandData struct {
	Name              string
	Type              disgord.ApplicationCommandType
	ApplicationID     disgord.Snowflake
	Description       string
	Options           []*disgord.ApplicationCommandOption
	DefaultPermission bool
}

type SlashCommand struct {
	Data    SlashCommandData
	Handler func(s disgord.Session, interaction *disgord.InteractionCreate)
}

type SlashCommands []SlashCommand

func (sc SlashCommands) FindByName(name string) SlashCommand {
	for _, slashCommand := range sc {
		if slashCommand.Data.Name == name {
			return slashCommand
		}
	}

	return SlashCommand{}
}
