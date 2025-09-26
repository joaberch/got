package model

type RemoteType int

const (
	Local RemoteType = iota
	Remote
	Git
)
