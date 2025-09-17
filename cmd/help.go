package cmd

import "fmt"

// ShowHelp prints the command-line help and usage information for got to standard output.
// The message includes available commands, short forms, usage examples, and the project source URL.
func ShowHelp() {
	fmt.Println("got - A simple version control system\n" +
		"\n" +
		"Usage:\n" +
		"  got <command> [arguments]\n" +
		"\n" +
		"Available Commands:\n" +
		"  help, h         Show this help message\n" +
		"  version, v      Display the current version of got\n" +
		"  init, i         Initialize a new got repository\n" +
		"  add, a <file>   Add a file to the staging area\n" +
		"  status, s       Show the status of the working directory\n" +
		"  commit, c <msg> Commit staged changes with a message\n" +
		"  restore, r <id> Restore a file from a previous commit by hash\n" +
		"\n" +
		"Examples:\n" +
		"  got init\n" +
		"  got add main.go\n" +
		"  got commit \"Initial commit\"\n" +
		"  got status\n" +
		"  got restore abc123\n" +
		"\n" +
		"Source code at https://github.com/joaberch/got")
}
