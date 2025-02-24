package command

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var folder = ".got"
var stagingFile = "staging.csv"

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
	Type string `json:"type"` //file or directory
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
	// Open file with read/write, create it if doesn't exist
	file, err := os.OpenFile(folder+"/"+stagingFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("impossible to open the staging file : %v", err)
	}
	defer file.Close() //Close the file at the end of the function

	writer := csv.NewWriter(file)
	defer writer.Flush() //Close the CSV writer at the end of the function

	// For each file/folder given
	for _, path := range paths {
		// Check the file exist
		info, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("the file or folder '%s' doesn't exist : %v", path, err)
		}

		//Check if is directory or file
		entryType := "file"
		if info.IsDir() {
			entryType = "directory"
		}

		// Write a line in the csv
		err = writer.Write([]string{path, entryType})
		if err != nil {
			return fmt.Errorf("impossible to write in the csv file : %v", err)
		}
	}

	fmt.Println("File added to staging.")
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
