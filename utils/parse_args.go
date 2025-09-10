package utils

import "Got/internal/model"

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
		}
	}

	return parsed
}
