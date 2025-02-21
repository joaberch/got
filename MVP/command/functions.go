package command

import (
	"errors"
	"os"
)

var ErrUnknownCommand = errors.New("command unknown")

var CommandsMap = map[string]Type{
	"hello": Hello,
	"init":  Init,
}

func GetCommand(name string) (Type, error) {
	if cmdType, found := CommandsMap[name]; found {
		return cmdType, nil
	}
	return -1, ErrUnknownCommand
}

func InitGot() error {
	err := os.Mkdir(".got", os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
