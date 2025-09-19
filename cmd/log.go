package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"path/filepath"
	"strconv"
	"strings"
)

func Log() error {
	commitsPath := filepath.Join(".got", "commits.csv")
	contents, err := utils.GetFileContent(commitsPath)
	if err != nil {
		return fmt.Errorf("error getting file contents: %s", err)
	}

	csvReader := csv.NewReader(strings.NewReader(string(contents)))
	records, err := csvReader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading file contents: %s", err)
	}

	for i := 0; i <= len(records)-1; i++ {
		record := records[i]
		if len(record) < 5 {
			continue //Skip
		}

		time, err := strconv.ParseInt(record[4], 10, 64)
		if err != nil {
			return fmt.Errorf("error converting time from file contents: %s", err)
		}
		commitDisplay := model.CommitDisplay{
			Hash:      record[0],
			Author:    record[2],
			Message:   record[3],
			Timestamp: time,
		}
		commitDisplay.Display()
	}

	return nil
}
