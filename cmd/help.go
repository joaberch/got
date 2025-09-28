package cmd

import "fmt"

// ShowHelp prints the command-line help and usage information for got to standard output.
func ShowHelp() {
	fmt.Println("got - A simple version control system\n" +
		"\n" +
		"Usage:\n" +
		"  got <command> [arguments]\n" +
		"\n" +
		"Available Commands:\n" +
		"  help					Show this help message\n" +
		"  version				Display the current version of Got\n" +
		"  init					Initialize a new Got repository\n" +
		"  add <file>			Add a file to the staging area\n" +
		"  status				Show the status of the working directory\n" +
		"  commit <msg>			Commit staged changes with a message\n" +
		"  restore <id>			Restore a file from a previous commit by hash\n" +
		"  log					Display the log from the commits file\n" +
		"  diff	[-v]			Display the differences in the file from the last commit who changed that file, verbose option\n" +
		"\n" +
		"Examples:\n" +
		"  got init\n" +
		"  got add main.go\n" +
		"  got commit \"Initial commit\"\n" +
		"  got status\n" +
		"  got restore abc123\n" +
		"  got log\n" +
		"  got diff\n" +
		"\n" +
		"Source code at https://github.com/joaberch/got")
}
