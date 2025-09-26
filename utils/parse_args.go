package utils

import "github.com/joaberch/got/internal/model"

// ParseArgs parses a slice of argument tokens and returns a model.ParsedArgs
// whose Command field is set to the last recognized command token.
//
// Supported tokens (short and long forms): "help"/"h", "version"/"v",
// "init"/"i", "add"/"a", "commit"/"c", "restore"/"r". Unrecognized tokens are
// ignored; if no supported token is found, the returned ParsedArgs has
// ParseArgs parses a slice of argument tokens and returns a model.ParsedArgs
// whose Command field is set to the last recognized command token.
//
// It recognizes long forms for commands:
// - "help" -> model.CmdHelp
// - "version" -> model.CmdVersion
// - "init" -> model.CmdInit
// - "add" -> model.CmdAdd
// - "commit" -> model.CmdCommit
// - "restore" -> model.CmdRestore
// - "log" -> model.CmdLog
// - "diff" -> model.CmdDiff
//
// Unrecognized tokens are ignored; if no supported token is found the returned
// ParsedArgs.Command remains model.CmdNone.
func ParseArgs(args []string) model.ParsedArgs {
	parsed := model.ParsedArgs{
		Command: model.CmdNone,
		Verbose: false,
	}

	for _, arg := range args {
		switch arg {
		case "help":
			parsed.Command = model.CmdHelp
		case "version":
			parsed.Command = model.CmdVersion
		case "init":
			parsed.Command = model.CmdInit
		case "add":
			parsed.Command = model.CmdAdd
		case "commit":
			parsed.Command = model.CmdCommit
		case "restore":
			parsed.Command = model.CmdRestore
		case "log":
			parsed.Command = model.CmdLog
		case "diff":
			parsed.Command = model.CmdDiff
		case "set-remote", "setRemote":
			parsed.Command = model.CmdSetRemote
		case "push":
			parsed.Command = model.CmdPush
		case "-v":
			parsed.Verbose = true
		}
	}

	return parsed
}
