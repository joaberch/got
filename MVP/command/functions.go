package command

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

	if runtime.GOOS == "windows" {
		cmd := exec.Command("attrib", "+H", ".got") //Utilise attrib pour cacher le dossier
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("impossible to hide the folder : %v", err)
		}
	}

	return nil
}
