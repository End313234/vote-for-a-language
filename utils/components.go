package utils

import "github.com/andersfylling/disgord"

type Button struct {
	Data    disgord.MessageComponent
	Handler func(session disgord.Session, interaction *disgord.InteractionCreate)
}

type Buttons []Button

func (b *Buttons) Add(buttons ...Button) {
	for _, button := range buttons {
		*b = append(*b, button)
	}
}

func (b Buttons) FindByCustomId(customId string) Button {
	for _, button := range b {
		if button.Data.CustomID == customId {
			return button
		}
	}

	return Button{}
}
