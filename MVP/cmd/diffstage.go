package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ShowDiffStage() {
	//Open the diff stage
	file, err := os.Open(stagingPath)
	if err != nil {
		fmt.Errorf("could not open staging file : %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Errorf("could not parse staging file : %v", err)
	}

	for _, record := range records {
		fmt.Println(string(record[0]))
	}
}
