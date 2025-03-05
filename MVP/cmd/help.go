package cmd

import "fmt"

// ShowHelp prints the usage instructions and a list of available commands to the standard output.
func ShowHelp() {
	fmt.Println("Usage: got [command] [arguments]")
	fmt.Println("Commands:")
	fmt.Println("    init     		Initialize a new project .got")
	fmt.Println("    stage    		Add file or folder into the staging area")
	fmt.Println("    commit   		Commit the changes in the staging area")
	fmt.Println("    unstage  		Remove file or folder from the staging area")
	fmt.Println("    help     		Display that help message")
	fmt.Println("    version  		Display the current version of Got")
	fmt.Println("    hello    		Test command")
	fmt.Println("    diffstage		Display the staging area")
}
