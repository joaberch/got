package command

import (
	"os"
	"testing"
)

func TestGetCommand(t *testing.T) {
	tests := []struct {
		input    string
		expected Type
		err      error
	}{
		{"hello", Hello, nil},
		{"init", Init, nil},
		{"help", Help, nil},
		{"version", Version, nil},
		{"stage", Stage, nil},
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

func TestReadStagingEntries(t *testing.T) {
	// Case 1 : File doesn't exist
	_ = os.RemoveAll(".got/staging.csv") // Remove file if exist

	//Try to read
	entries, err := ReadStagingEntries()
	if err != nil {
		t.Errorf("Unexpected error : %v", err)
	}
	if len(entries) != 0 {
		t.Errorf("Expected no entries, got %d", len(entries))
	}

	// Case 2 : staging file invalid
	os.Mkdir(".got", os.ModePerm)
	file, _ := os.Create(".got/staging.csv")
	file.WriteString("/test/file.txt,file\n/test/folder,directory\n")
	file.Close()

	//Try to read
	entries, err = ReadStagingEntries()
	if err != nil {
		t.Errorf("Erreur inattendue : %v", err)
	}
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(entries))
	}

	// Case 3 : One invalid line
	file, _ = os.Create(".got/staging.csv")
	file.WriteString("Path,Type\n/test/file.txt\n") // invalid line
	file.Close()

	_, err = ReadStagingEntries()
	if err == nil {
		t.Errorf("Expected an error for invalid line, got none")
	}
}

func TestAddEntryToStaging(t *testing.T) {
	_ = os.RemoveAll(".got") // Clean before test
	os.Mkdir(".got", os.ModePerm)

	// Case 1 : Add file
	os.WriteFile("test.txt", []byte("data"), 0644)
	err := AddEntryToStaging([]string{"test.txt"})
	if err != nil {
		t.Errorf("Unexpected error : %v", err)
	}

	entries, _ := ReadStagingEntries()
	if len(entries) != 1 || entries[0].Path != "test.txt" || entries[0].Type != "root" {
		t.Errorf("Expected one file entry, got %+v", entries)
		t.Errorf("More debug : len(entries) = %d, entries[0].Path = %s, entries[0].Type = %s", len(entries), entries[0].Path, entries[0].Type)
	}

	// Case 2 : Add folder
	os.Mkdir("testDir", os.ModePerm)
	os.WriteFile("testDir/file1.txt", []byte("data"), 0644)
	os.Mkdir("testDir/subDir", os.ModePerm)
	os.WriteFile("testDir/subDir/file2.txt", []byte("data"), 0644)

	err = AddEntryToStaging([]string{"testDir"})
	if err != nil {
		t.Errorf("Unexpected error : %v", err)
	}

	entries, _ = ReadStagingEntries()
	if len(entries) != 5 { // Folder itself + 4 file/dirs
		t.Errorf("Expected 5 entries, got %d", len(entries))
		t.Errorf("More debug : %+v", entries)
	}

	// Case 3 : Add duplicate
	err = AddEntryToStaging([]string{"test.txt"})
	if err == nil {
		t.Errorf("Expected duplicate error, got none")
	}
}
