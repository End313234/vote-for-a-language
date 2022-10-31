package commands

import (
	"vote-for-a-language/utils"
)

var SlashCommands = utils.SlashCommands{
	{
		Data:    AddLanguageData,
		Handler: AddLanguageHandler,
	},
	{
		Data:    VoteData,
		Handler: VoteHandler,
	},
	{
		Data:    InviteData,
		Handler: InviteHandler,
	},
}
