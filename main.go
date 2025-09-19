package main

import (
	"github.com/joaberch/got/cmd"
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"log"
	"os"
)

// main is the CLI entry point. It parses command-line arguments, dispatches the requested
// subcommand (help, version, init, add, commit, restore, log, diff) to the cmd package,
// and exits with a non-zero status if a command returns an error. Missing required
// arguments for commit and restore cause an immediate fatal exit with an explanatory message.
func main() {
	var err error
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
		err = cmd.Init()
	case model.CmdAdd:
		if len(args) > 1 {
			err = cmd.Add(args[1])
		}
	case model.CmdCommit:
		if len(args) > 1 {
			err = cmd.Commit(args[1])
		} else {
			log.Fatal("No commit message specified")
		}
	case model.CmdRestore:
		if len(args) > 1 {
			err = cmd.Restore(args[1])
		} else {
			log.Fatal("You need to specify the hash of the file you want to restore")
		}
	case model.CmdLog:
		err = cmd.Log()
	case model.CmdDiff:
		err = cmd.Diff()
	}

	if err != nil {
		log.Fatal(err)
	}
}
