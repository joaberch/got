package command

type Type int

const (
	Hello Type = iota
	Init
	Help
	Version
	Stage
	Unstage
)
