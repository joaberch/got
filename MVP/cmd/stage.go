package cmd

import (
	"MVP/utils"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// HandleStageCommand adds the provided files or directories to the staging area and logs success or error messages.
func HandleStageCommand(paths []string) {
	//Check if the file are already tracked
	untrackedPaths, alreadyTracked, err := CheckIsAlreadyTracked(paths)
	if err != nil {
		fmt.Errorf("error while checking if the files are already tracked : %v\n", err)
		return
	}
	if alreadyTracked && len(untrackedPaths) == 0 {
		fmt.Errorf("all the files are already tracked")
	}

	err = utils.AddEntryToStaging(paths)
	if err != nil {
		fmt.Errorf("error while adding in stage : %v\n", err)
		return
	}
	fmt.Println(strings.Join(paths, ", ") + " successfully added to the staging area!")
}

func CheckIsAlreadyTracked(paths []string) ([]string, bool, error) {
	file, err := os.Open(trackingFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return paths, false, nil
		}
		return nil, false, fmt.Errorf("error while opening the tracking file : %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Errorf("error while closing the tracking file : %v", err)
		}
	}(file)

	trackedPaths := make(map[string]bool)
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, false, fmt.Errorf("error while reading the tracking file : %v", err)
	}

	for _, record := range records {
		if len(record) > 0 {
			trackedPaths[record[0]] = true
		}
	}

	var untrackedPaths []string
	alreadyTracked := false
	for _, path := range paths {
		if _, isTracked := trackedPaths[path]; isTracked {
			alreadyTracked = true
		} else {
			untrackedPaths = append(untrackedPaths, path)
		}
	}
	return untrackedPaths, alreadyTracked, nil
}
