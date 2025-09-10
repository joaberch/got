package main

import (
	"Got/cmd"
	"Got/internal/model"
	"Got/utils"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		cmd.ShowHelp()
	}

	parsed := utils.ParseArgs(args)

	switch parsed.Command {
	case model.CmdNone:
		break
	case model.CmdHelp:
		cmd.ShowHelp()
	case model.CmdVersion:
		cmd.ShowVersion()
	case model.CmdInit:
		cmd.Init()
	case model.CmdAdd:
		if len(args) > 1 {
			cmd.Add(args[1])
		}
	case model.CmdStatus:
		cmd.Status()
	}
}
