package main

import (
	"github.com/joaberch/got/cmd"
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"log"
	"os"
)

// missing commit message or missing restore hash terminate the program via log.Fatal.
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
		cmd.Init()
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
			cmd.Restore(args[1])
		} else {
			log.Fatal("You need to specify the hash of the file you want to restore")
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
