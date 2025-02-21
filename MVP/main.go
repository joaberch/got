package main

import (
	"MVP/command"
	"fmt"
	"os"
)

func getCommand(name string) (command.Type, error) {
	switch name {
	case "hello":
		return command.Hello, nil
	default:
		return -1, fmt.Errorf("command unknown")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: got hello")
		return
	}

	userCommand, err := getCommand(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if userCommand == command.Hello {
		fmt.Println("got got got")
	}
}
