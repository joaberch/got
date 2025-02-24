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

var folder = ".got"
var stagingFile = "staging.json"

var ErrUnknownCommand = errors.New("command unknown")

// CommandsMap is a mapping of command names to their corresponding command types, defining available commands for lookup.
var CommandsMap = map[string]Type{
	"hello":   Hello,
	"init":    Init,
	"help":    Help,
	"version": Version,
	"add":     Add,
}

// StagingEntry enum the json key name
type StagingEntry struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

// GetCommand retrieves the command type for a given name from CommandsMap, or returns an error if the command is unknown.
func GetCommand(name string) (Type, error) {
	if cmdType, found := CommandsMap[name]; found {
		return cmdType, nil
	}
	return -1, ErrUnknownCommand
}

// AddEntryToStaging adds provided file or directory paths to the staging file, creating or updating it as needed.
// It validates the paths, determines their type (file or directory), and adds entries to the staging JSON file.
// Returns an error if any operation, such as file reading, marshalling, or writing, fails.
func AddEntryToStaging(paths []string) error {
	var staging []StagingEntry //Get the enum of json key

	//Try reading the file, check it exist and can be processed
	if _, err := os.Stat(folder + "/" + stagingFile); err == nil {
		_, err := ioutil.ReadFile(folder + "/" + stagingFile)
		if err != nil {
			return fmt.Errorf("impossible to read the staging file : %v", err)
		}
	}

	//Foreach parameter given (each one is a file or a folder)
	for _, path := range paths {
		info, err := os.Stat(path) //Check the path
		if err != nil {
			return fmt.Errorf("impossible to get the info of the file : %v", err)
		}

		//If is a file or a folder
		entryType := "file"
		if info.IsDir() {
			entryType = "directory"
		}

		//Add the path and the type (file/folder)
		staging = append(staging, StagingEntry{Path: path, Type: entryType})
	}
	data, err := json.MarshalIndent(staging, "", "  ") //Serialize object
	if err != nil {
		return fmt.Errorf("impossible to marshal the staging file : %v", err)
	}
	if err := ioutil.WriteFile(folder+"/"+stagingFile, data, os.ModePerm); err != nil {
		return fmt.Errorf("impossible to write the staging file : %v", err)
	}
	fmt.Printf("Files added to staging")
	return nil
}

// HandleAddCommand processes the "add" command by adding specified files/directories to the staging file and validating inputs.
func HandleAddCommand(args []string) {
	if len(args) < 1 {
		fmt.Println("You must specify at least one file or directory")
		fmt.Println("Usage : got add <file1|folder1> [<file2/folder2>...]")
		return
	}

	err := AddEntryToStaging(args)
	if err != nil {
		fmt.Println("Error :", err)
	}
}

// ShowVersion prints the current version of the application to the standard output.
func ShowVersion() {
	fmt.Println("Got version 0.0.1")
}

// ShowHelp prints the available commands from the CommandsMap to the standard output.
func ShowHelp() {
	fmt.Println("Got commands:")
	for name := range CommandsMap {
		fmt.Println("got", name)
	}
}

// InitGot initializes the `.got` folder and creates a hidden folder on Windows with a staging file for version control.
func InitGot() error {
	err := os.Mkdir(folder, os.ModePerm)
	if err != nil {
		return err
	}

	//Create .got folder
	if runtime.GOOS == "windows" {
		cmd := exec.Command("attrib", "+H", folder) //Use attrib to hide the folder
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("impossible to hide the folder : %v", err)
		}
	}

	//Create the staging file
	if _, err := os.Stat(folder + "/" + stagingFile); err != nil {
		if _, err := os.Create(folder + "/" + stagingFile); err != nil {
			return fmt.Errorf("impossible to create the staging file : %v", err)
		}
	}
	return nil
}
