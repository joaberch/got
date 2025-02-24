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
	"help":  Help,
}

func GetCommand(name string) (Type, error) {
	if cmdType, found := CommandsMap[name]; found {
		return cmdType, nil
	}
	return -1, ErrUnknownCommand
}

func ShowHelp() {
	fmt.Println("Got commands:")
	for name := range CommandsMap {
		fmt.Println("got", name)
	}
}

func InitGot() error {
	err := os.Mkdir(".got", os.ModePerm)
	if err != nil {
		return err
	}

	if runtime.GOOS == "windows" {
		cmd := exec.Command("attrib", "+H", ".got") //Use attrib to hide the folder
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("impossible to hide the folder : %v", err)
		}
	}

	return nil
}
