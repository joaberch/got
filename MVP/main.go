package main

import (
	"MVP/command"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		command.ShowHelp()
		return
	}

	userCommand, err := command.GetCommand(os.Args[1])
	if err != nil {
		fmt.Println(err)
		command.ShowHelp()
		return
	}

	switch userCommand {
	case command.Hello:
		fmt.Println("got got got")
		break
	case command.Init:
		command.InitGot()
		break
	case command.Help:
		command.ShowHelp()
		break
	case command.Version:
		command.ShowVersion()
		break
	case command.Stage:
		command.HandleAddCommand(os.Args[2:]) //give everything after the second element (got stage ...)
		command.HandleStageCommand(os.Args[2:]) //give everything after the second element (got stage ...)
		break
	default:
		command.ShowHelp()
	}
}
