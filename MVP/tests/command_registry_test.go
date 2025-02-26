package tests

import (
	"MVP/command"
	"testing"
)

func TestGetCommand(t *testing.T) {
	tests := []struct {
		input    string
		expected command.Type
		err      error
	}{
		{"hello", command.Hello, nil},
		{"init", command.Init, nil},
		{"help", command.Help, nil},
		{"version", command.Version, nil},
		{"stage", command.Stage, nil},
		{"unstage", command.Unstage, nil},
		{"doesntExist", -1, command.ErrUnknownCommand},
		{"", -1, command.ErrUnknownCommand},
	}

	for _, test := range tests {
		result, err := command.GetCommand(test.input)
		if result != test.expected || err != test.err {
			t.Errorf("For input '%s', expected (%v, %v) but got (%v, %v)", test.input, test.expected, test.err, result, err)
		}
	}
}
