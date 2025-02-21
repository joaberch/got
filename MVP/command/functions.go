package command

import (
	"errors"
)

var ErrUnknownCommand = errors.New("command unknown")

var CommandsMap = map[string]Type{
	"hello": Hello,
}

func GetCommand(name string) (Type, error) {
	if cmdType, found := CommandsMap[name]; found {
		return cmdType, nil
	}
	return -1, ErrUnknownCommand
}
