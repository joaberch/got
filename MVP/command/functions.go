package command

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

// ReadStagingEntries reads the staging CSV file and returns a slice of StagingEntry or an empty slice if the file doesn't exist.
// Returns an error if the file cannot be opened, read, or contains invalid entries.
func ReadStagingEntries() ([]StagingEntry, error) {
	// Check if the staging file exist
	file, err := os.Open(folder + "/" + stagingFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []StagingEntry{}, nil // If doesn't exist return empty list
		}
		return nil, fmt.Errorf("impossible d'ouvrir le fichier CSV : %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read every line of the csv
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("impossible de lire le fichier CSV : %v", err)
	}

	// COnvert every record in StagingEntry (ignore the first 2 object)
	var entries []StagingEntry
	for i, record := range records {
		if len(record) < 2 {
			return nil, fmt.Errorf("ligne %d invalide dans le fichier CSV", i+1)
		}
		entries = append(entries, StagingEntry{
			Path: record[0],
			Type: record[1],
		})
	}

	return entries, nil
}

// AddEntryToStaging adds provided file or directory paths to the staging file, creating or updating it as needed.
// It validates the paths, determines their type (file or directory), and adds entries to the staging JSON file.
// Returns an error if any operation, such as file reading, marshalling, or writing, fails.
func AddEntryToStaging(paths []string) error {
	existingEntries, err := ReadStagingEntries()
	if err != nil {
		return fmt.Errorf("impossible to read the staging file : %v", err)
	}

	entryMap := make(map[string]struct{})
	for _, entry := range existingEntries {
		entryMap[entry.Path] = struct{}{}
	}

	// Open file with read/write, create it if doesn't exist
	file, err := os.OpenFile(folder+"/"+stagingFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
		if info.IsDir() { //Recursive
			//Get children
			entries, err := os.ReadDir(path)
			if err != nil {
				return fmt.Errorf("impossible to read the directory '%s' : %v", path, err)
			}

			//Recursive call
			for _, entry := range entries {
				childPath := filepath.Join(path, entry.Name())
				err := AddEntryToStaging([]string{childPath})
				if err != nil {
					return err //If there's one error it's everywhere
				}
			}

			if _, exists := entryMap[path]; !exists {
				err := writer.Write([]string{path, "directory"})
				if err != nil {
					return fmt.Errorf("impossible to write in the csv file : %v", err)
				}
				entryMap[path] = struct{}{}
			}
		} else {

			//Prevent duplicate
			if _, exists := entryMap[path]; exists {
				return fmt.Errorf("the file or folder '%s' is already in the staging file, ignored", path)
			}

			// Write a line in the csv
			err = writer.Write([]string{path, "root"})
			if err != nil {
				return fmt.Errorf("impossible to write in the csv file : %v", err)
			}
			entryMap[path] = struct{}{}
		}
	}

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
	fmt.Println("File(s) added to staging area.")
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
