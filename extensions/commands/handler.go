package commands

import (
	"vote-for-a-language/utils"
)

var SlashCommands = utils.SlashCommands{
	{
		Data:    RequestNewLanguageData,
		Handler: RequestNewLanguageHandler,
	},
}
