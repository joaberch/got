package cmd

import "fmt"

// ShowHelp prints the command-line help and usage information for Got to standard output.
// The output includes the tool title, usage line, available commands with short forms,
// example commands, and the project source URL.
func ShowHelp() {
	fmt.Println("Got - A simple version control system\n" +
		"\n" +
		"Usage:\n" +
		"  got <command> [arguments]\n" +
		"\n" +
		"Available Commands:\n" +
		"  help, h         Show this help message\n" +
		"  version, v      Display the current version of Got\n" +
		"  init, i         Initialize a new Got repository\n" +
		"  add, a <file>   Add a file to the staging area\n" +
		"  status, s       Show the status of the working directory\n" +
		"  commit, c <msg> Commit staged changes with a message\n" +
		"  restore, r <id> Restore a file from a previous commit by hash\n" +
		"  log, l          Display the log from the commits file\n" +
		"  diff, d         Display the differences in the file from the last commit" + //TODO - get (distinct) all files from all commits and use all of them for a better diff display
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
