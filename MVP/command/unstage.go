package command

import (
	"fmt"
)

// HandleUnStageCommand removes specified files or directories from the staging area and logs success or error messages.
func HandleUnStageCommand(paths []string) {
	err := RemoveEntryToStaging(paths)
	if err != nil {
		fmt.Printf("Error while removing '%v' from the staging area : %v\n", paths, err)
		return
	}
	fmt.Printf("'%v' successfully removed from the staging area !\n", paths)
}
