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

	if userCommand == command.Hello {
		fmt.Println("got got got")
	}

	if userCommand == command.Init {
		command.InitGot()
	}

	if userCommand == command.Help {
		command.ShowHelp()
	}

	if userCommand == command.Version {
		command.ShowVersion()
	}

	if userCommand == command.Add {
		command.HandleAddCommand(os.Args[2:])
	}
}
