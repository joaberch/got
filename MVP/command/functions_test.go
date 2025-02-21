package command

import (
	"testing"
)

func TestGetCommand(t *testing.T) {
	tests := []struct {
		input    string
		expected Type
		err      error
	}{
		{"hello", Hello, nil},
		{"doesntExist", -1, ErrUnknownCommand},
		{"", -1, ErrUnknownCommand},
	}

	for _, test := range tests {
		result, err := GetCommand(test.input)
		if result != test.expected || err != test.err {
			t.Errorf("For input '%s', expected (%v, %v) but got (%v, %v)", test.input, test.expected, test.err, result, err)
		}
	}
}
