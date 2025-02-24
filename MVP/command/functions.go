package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

var ErrUnknownCommand = errors.New("command unknown")

var CommandsMap = map[string]Type{
	"hello":   Hello,
	"init":    Init,
	"help":    Help,
	"version": Version,
	"add":     Add,
}

type StagingEntry struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

func GetCommand(name string) (Type, error) {
	if cmdType, found := CommandsMap[name]; found {
		return cmdType, nil
	}
	return -1, ErrUnknownCommand
}

//goland:noinspection GoDeprecation,GoDeprecation
func AddEntryToStaging(paths []string) error {
	var staging []StagingEntry

	if _, err := os.Stat(".got/staging"); err == nil {
		data, err := ioutil.ReadFile(".got/staging")
		if err != nil {
			return fmt.Errorf("impossible to read the staging file : %v", err)
		}
		if err := json.Unmarshal(data, &staging); err != nil {
			return fmt.Errorf("impossible to unmarshal the staging file : %v", err)
		}
	}

	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("impossible to get the info of the file : %v", err)
		}
		entryType := "file"
		if info.IsDir() {
			entryType = "directory"
		}

		staging = append(staging, StagingEntry{Path: path, Type: entryType})
	}
	data, err := json.MarshalIndent(staging, "", "  ")
	if err != nil {
		return fmt.Errorf("impossible to marshal the staging file : %v", err)
	}
	if err := ioutil.WriteFile(".got/staging", data, os.ModePerm); err != nil {
		return fmt.Errorf("impossible to write the staging file : %v", err)
	}
	fmt.Println("Staging file updated")
	return nil
}

func HandleAddCommand(args []string) {
	if len(args) < 1 {
		fmt.Println("You must specify at least one file or directory")
		fmt.Println("Usage : got add <file1|folder1> [<file2/folder2>...]")
		return
	}

	err := AddEntryToStaging(args)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
}

func ShowVersion() {
	fmt.Println("Got version 0.0.1")
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

	//Create .got folder
	if runtime.GOOS == "windows" {
		cmd := exec.Command("attrib", "+H", ".got") //Use attrib to hide the folder
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("impossible to hide the folder : %v", err)
		}
	}

	return nil
}
