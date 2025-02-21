package main

import (
	"MVP/command"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: got hello")
		return
	}

	userCommand, err := command.GetCommand(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if userCommand == command.Hello {
		fmt.Println("got got got")
	}

	if userCommand == command.Init {
		command.InitGot()
	}
}
