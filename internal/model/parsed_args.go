package model

// ParsedArgs manages the argument parsing
type ParsedArgs struct {
	Command CommandType
	Verbose bool
	Other   string
}
