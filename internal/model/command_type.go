package model

// CommandType enumerates every possible command
type CommandType int

const (
	CmdNone CommandType = iota
	CmdHelp
	CmdVersion
	CmdInit
	CmdAdd
	CmdCommit
	CmdRestore
	CmdLog
)
