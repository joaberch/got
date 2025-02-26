package command

import (
	"fmt"
	"strings"
)

// HandleStageCommand adds the provided files or directories to the staging area and logs success or error messages.
func HandleStageCommand(paths []string) {
	err := AddEntryToStaging(paths)
	if err != nil {
		fmt.Printf("Error while adding in stage : %v\n", err)
		return
	}
	fmt.Println(strings.Join(paths, ", ") + " successfully added to the staging area!")

}
