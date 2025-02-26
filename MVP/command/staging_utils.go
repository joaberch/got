package command

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var folder = ".got"
var stagingFile = "staging.csv"
var stagingPath = folder + "/" + stagingFile

// StagingEntry enum the json key name
type StagingEntry struct {
	Path string `json:"path"`
	Type string `json:"type"` //file or directory
}

// ReadStagingEntries reads the staging CSV file and returns a slice of StagingEntry or an empty slice if the file doesn't exist.
// Returns an error if the file cannot be opened, read, or contains invalid entries.
func ReadStagingEntries() ([]StagingEntry, error) {
	// Check if the staging file exist
	file, err := os.Open(stagingPath)
	if err != nil {
		if os.IsNotExist(err) {
			return []StagingEntry{}, nil // If it doesn't exist return empty list
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

	// Convert every record in StagingEntry (ignore the first 2 object)
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

func RemoveEntryToStaging(paths []string) error {
	file, err := os.OpenFile(stagingPath, os.O_RDONLY, 0644) // Open the file in reading
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("the staging file doesn't exist : %v", err)
		}
		return fmt.Errorf("impossible to open the staging file : %v", err)
	}
	defer file.Close()

	//Bufio to read line by line
	scanner := bufio.NewScanner(file)
	var filteredLine []string

	//Foreach line
	for scanner.Scan() {
		line := scanner.Text()
		//Get the name of the file read
		text := strings.Split(line, ",")[0]
		//If we should delete that line
		shouldRemove := false

		//Check if arg match line
		for _, path := range paths {
			absPath, _ := filepath.Abs(path) //get absolute path
			absText, _ := filepath.Abs(text) //get absolute path

			//if absPath == absText they are the same file
			//if it has the prefix we remove the whole folder
			if absPath == absText || strings.HasPrefix(absText, absPath+"\\") {
				shouldRemove = true
				break
			}
		}

		//Keep the line
		if !shouldRemove {
			filteredLine = append(filteredLine, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error while reading the staging file : %v", err)
	}

	//Open file in writing with trunc to overwrite it
	file, err = os.OpenFile(stagingPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("impossible to open the staging file : %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, entry := range filteredLine {
		_, err := writer.WriteString(entry + "\n")
		if err != nil {
			return fmt.Errorf("impossible to write in the csv file : %v", err)
		}
	}

	//Check everything has been written
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("error while flushing data in the staging file : %v", err)
	}

	return nil
}

// AddEntryToStaging adds specified files or directories to the staging area, ensuring no duplicates and validating their existence.
// Returns an error if any file or directory does not exist or cannot be processed.
func AddEntryToStaging(paths []string) error {
	existingEntries, err := ReadStagingEntries()
	if err != nil {
		return fmt.Errorf("impossible to read the staging file : %v", err)
	}

	entryMap := make(map[string]struct{})
	for _, entry := range existingEntries {
		entryMap[entry.Path] = struct{}{}
	}

	// Open file with read/write, create it if it doesn't exist
	file, err := os.OpenFile(stagingPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
				checksum := GetChecksum(path) //Get SHA-1 hash
				err := writer.Write([]string{path, checksum, "added/removed"})
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
			checksum := GetChecksum(path) //Get SHA-1 hash
			err = writer.Write([]string{path, checksum, "added/removed"})
			if err != nil {
				return fmt.Errorf("impossible to write in the csv file : %v", err)
			}
			entryMap[path] = struct{}{}
		}
	}

	return nil
}
