package utils

import "Got/internal/model"

// ParseArgs parses a slice of argument tokens and returns a model.ParsedArgs
// whose Command field is set to the last recognized command token.
//
// Supported tokens (short and long forms): "help"/"h", "version"/"v",
// "init"/"i", "add"/"a", "commit"/"c", "restore"/"r". Unrecognized tokens are
// ignored; if no supported token is found, the returned ParsedArgs has
// Command == model.CmdNone.
func ParseArgs(args []string) model.ParsedArgs {
	parsed := model.ParsedArgs{
		Command: model.CmdNone,
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "help", "h":
			parsed.Command = model.CmdHelp
		case "version", "v":
			parsed.Command = model.CmdVersion
		case "init", "i":
			parsed.Command = model.CmdInit
		case "add", "a":
			parsed.Command = model.CmdAdd
		case "commit", "c":
			parsed.Command = model.CmdCommit
		case "restore", "r":
			parsed.Command = model.CmdRestore
		}
	}

	return parsed
}
