package command

import (
	"errors"
)

var ErrUnknownCommand = errors.New("command unknown")

// CommandsMap is a mapping of command names to their corresponding command types, defining available commands for lookup.
var CommandsMap = map[string]Type{
	"hello":   Hello,
	"init":    Init,
	"help":    Help,
	"version": Version,
	"stage":   Stage,
	"unstage": Unstage,
	"commit":  Commit,
}

// GetCommand retrieves the command type for a given name from CommandsMap, or returns an error if the command is unknown.
func GetCommand(name string) (Type, error) {
	if cmdType, found := CommandsMap[name]; found {
		return cmdType, nil
	}
	return -1, ErrUnknownCommand
}
