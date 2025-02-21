package command

import "fmt"

func GetCommand(name string) (Type, error) {
	switch name {
	case "hello":
		return Hello, nil
	default:
		return -1, fmt.Errorf("command unknown")
	}
}
