package model

type CommandType int

const (
	CmdNone CommandType = iota
	CmdHelp
	CmdVersion
	CmdInit
	CmdAdd
	CmdStatus
	CmdCommit
	CmdRestore
)
