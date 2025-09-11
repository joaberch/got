package utils

import "Got/internal/model"

// ParseArgs process each argument given
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
		case "status", "s":
			parsed.Command = model.CmdStatus
		case "commit", "c":
			parsed.Command = model.CmdCommit
		case "restore", "r":
			parsed.Command = model.CmdRestore
		}
	}

	return parsed
}
