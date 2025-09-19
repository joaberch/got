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

// Log reads commits from the repository metadata file (.got/commits.csv) and displays each valid commit.
// It skips CSV rows with fewer than five fields. The function returns an error if the commits file cannot
// be read, if the CSV cannot be parsed, or if a commit timestamp cannot be converted to an int64. On
// success it returns nil.
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
