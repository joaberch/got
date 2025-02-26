package command

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// InitProject initializes a .got repository by creating the respective directory with necessary permissions.
func InitProject() {
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		fmt.Printf("Error while initialization : %v\n", err)
		return
	}

	//Create .got folder
	if runtime.GOOS == "windows" {
		cmd := exec.Command("attrib", "+H", folder) //Use attrib to hide the folder
		err := cmd.Run()
		if err != nil {
			fmt.Printf("impossible to hide the folder : %v", err)
		}
	}

	//Create the staging file
	if _, err := os.Stat(stagingPath); err != nil {
		if _, err := os.Create(stagingPath); err != nil {
			fmt.Printf("impossible to create the staging file : %v", err)
		}
	}
	fmt.Println("Got project initialized successfully !")
}
