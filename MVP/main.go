package main

import (
	"MVP/cmd"
	"MVP/command"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		cmd.ShowHelp()
		return
	}

	userCommand, err := command.GetCommand(os.Args[1])
	if err != nil {
		fmt.Println(err)
		cmd.ShowHelp()
		return
	}

	switch userCommand {
	case command.Hello:
		fmt.Println("got got got")
		break
	case command.Init:
		cmd.InitProject()
		break
	case command.Help:
		cmd.ShowHelp()
		break
	case command.Version:
		cmd.ShowVersion()
		break
	case command.Stage:
		cmd.HandleStageCommand(os.Args[2:]) //give everything after the second element (got stage ...)
		break
	case command.Unstage:
		cmd.HandleUnStageCommand(os.Args[2:]) //give every argument except the first two (got unstage ...)
		break
	default:
		cmd.ShowHelp()
	}
}
