package main

import (
	"fmt"
	"os"
)

type commandType int

const (
	CommandHello commandType = iota
	Other
)

func getCommand(name string) (commandType, error) {
	switch name {
	case "hello":
		return CommandHello, nil
	default:
		return -1, fmt.Errorf("Command unknown")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: got hello")
		return
	}

	command, err := getCommand(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if command == CommandHello {
		fmt.Println("got got got")
	}
}
