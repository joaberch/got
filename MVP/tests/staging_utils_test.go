package tests

import (
	"MVP/utils"
	"os"
	"testing"
)

//TODO : F.A.S.T (Independant?)
////////////////////////////////////// Test ReadStagingEntries /////////////////////////////////////////////////

// TestReadStagingEntries_NoFile verifies that ReadStagingEntries returns no entries and no error when the staging file is absent.
func TestReadStagingEntries_NoFile(t *testing.T) {
	// Set up the test environment
	_ = os.RemoveAll(".got")

	//Try to read
	entries, err := utils.ReadStagingEntries()
	if err != nil {
		t.Errorf("Unexpected error : %v", err)
	}
	//Check if the func returned something, it should not because the file doesn't exist
	if len(entries) != 0 {
		t.Errorf("Expected no entries, got %d : %v", len(entries), entries)
	}
}

// TestReadStagingEntries_ValidLine verifies that ReadStagingEntries correctly reads valid staging entries from the CSV file.
func TestReadStagingEntries_ValidLine(t *testing.T) {
	//Set up the test environment
	_ = os.RemoveAll(".got")
	os.Mkdir(".got", os.ModePerm)
	file2, _ := os.Create("file2.txt")
	defer os.Remove("file2.txt")
	defer file2.Close()
	file, _ := os.Create(".got/staging.csv")
	file.WriteString("file.txt,fbb2e7c65e1506824da06ce6239189b5e50837e8,added\n" +
		"/test/folder,3fcceacb3a02c3b3a5b07814b68cadb65b5ae7c7, added\n")
	file.Close()

	//Try to read
	entries, err := utils.ReadStagingEntries()
	if err != nil {
		t.Errorf("Erreur inattendue : %v", err)
	}
	//Check we got the right number of entries
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(entries))
	}
}

func TestReadStagingEntries_MalformedLine(t *testing.T) {
	//Set up the test environment
	os.RemoveAll(".got")
	os.Mkdir(".got", os.ModePerm)
	defer os.RemoveAll(".got")
	file3, err := os.Create("file3.txt")
	defer os.Remove("file3.txt")
	defer file3.Close()

	file, _ := os.Create(".got/staging.csv")

	//MalformedLine
	file.WriteString("file3.txt,fbb2e7c65e1506824da06ce6239189b5e50837e8\n" +
		"/test/folder, added\n" +
		"malformedLine\n")
	file.Close()

	_, err = utils.ReadStagingEntries()
	if err == nil {
		t.Errorf("Expected error for malformed line, got none")
	}
}

func TestReadStagingEntries_EmptyFile(t *testing.T) {
	//Set up the test environment
	os.RemoveAll(".got")
	os.Mkdir(".got", os.ModePerm)
	defer os.RemoveAll(".got")
	file, _ := os.Create(".got/staging.csv")
	file.Close()

	entries, err := utils.ReadStagingEntries()
	if err != nil {
		t.Errorf("Unexpected error : %v", err)
	}
	if len(entries) != 0 {
		t.Errorf("Expected no entries, got %d", len(entries))
	}
}

func TestAddEntryToStaging_SuccessFile(t *testing.T) {
	//Set up the test environment
	_ = os.RemoveAll(".got")
	os.Mkdir(".got", os.ModePerm)
	defer os.RemoveAll(".got")
	os.Create(".got/staging.csv")

	//create a test file
	os.WriteFile("test.txt", []byte("data"), 0644)
	err := utils.AddEntryToStaging([]string{"test.txt"})
	if err != nil {
		t.Errorf("Unexpected error : %v", err)
	}

	entries, _ := utils.ReadStagingEntries()
	if len(entries) != 1 {
		t.Errorf("Expected 1 entry, got %d, %v", len(entries), entries)
	}
	if entries[0].Path != "test.txt" {
		t.Errorf("Expected path 'test.txt', got '%s'", entries[0].Path)
	}
	expectedHash, err := utils.GetChecksum(entries[0].Path)
	if err != nil {
		t.Errorf("Failed to compute checksum for path '%s': %v", entries[0].Path, err)
	}
	if entries[0].Hash != expectedHash {
		t.Errorf("Expected hash '%s', got '%s'", expectedHash, entries[0].Hash)
	}
	if entries[0].State != "added" {
		t.Errorf("Expected state 'added', got '%s'", entries[0].State)
	}
}

func TestAddEntryToStaging_SuccessFolder(t *testing.T) {
	//Set up the test environment
	_ = os.RemoveAll(".got")
	os.Mkdir(".got", os.ModePerm)
	defer os.RemoveAll(".got")
	os.Create(".got/staging.csv")
	os.Mkdir("testDir", os.ModePerm)
	defer os.RemoveAll("testDir")
	os.WriteFile("testDir/file1.txt", []byte("data"), 0644)
	os.Mkdir("testDir/subDir", os.ModePerm)
	os.WriteFile("testDir/subDir/file2.txt", []byte("data"), 0644)

	//Add to the staging area
	err := utils.AddEntryToStaging([]string{"testDir"})
	if err != nil {
		t.Errorf("Unexpected error : %v", err)
	}

	//Read the staging file
	entries, _ := utils.ReadStagingEntries()
	if len(entries) != 4 { // Folder itself + 3 file/dirs
		t.Errorf("Expected 5 entries, got %d, %v", len(entries), entries)
	}
}

func TestAddEntryToStaging_DuplicateError(t *testing.T) {
	//Set up the test environment
	_ = os.RemoveAll(".got")
	os.Mkdir(".got", os.ModePerm)
	os.Create(".got/staging.csv")
	os.WriteFile("test.txt", []byte("data"), 0644)

	//Add for the first time
	err := utils.AddEntryToStaging([]string{"test.txt"})
	if err != nil {
		t.Errorf("Unexpected error : %v", err)
	}

	//Add for the second time
	err = utils.AddEntryToStaging([]string{"test.txt"})
	if err == nil {
		t.Errorf("Expected duplicate error, got none")
	}

	//Clean test output
	defer os.Remove("test.txt")
	defer os.RemoveAll(".got")
	defer os.RemoveAll("/tests/.got")
}

func TestAddEntryToStaging_EmptyFolder(t *testing.T) {
	//Set up the test environment
	os.RemoveAll(".got")
	os.Mkdir(".got", os.ModePerm)
	os.Create(".got/staging.csv")
	os.Mkdir("emptyDir", os.ModePerm)

	err := utils.AddEntryToStaging([]string{"emptyDir"})
	if err != nil {
		t.Errorf("Unexpected error : %v", err)
	}
	entries, _ := utils.ReadStagingEntries()
	if len(entries) != 1 {
		t.Errorf("Expected 1 entry, got %d, %v", len(entries), entries)
	}
	if entries[0].Path != "emptyDir" {
		t.Errorf("Expected path 'emptyDir', got '%s'", entries[0].Path)
	}

	//Clean test output
	os.RemoveAll("emptyDir")
	os.RemoveAll(".got")
}

func TestAddEntryToStaging_NonExistentFile(t *testing.T) {
	//Set up test environment
	os.RemoveAll(".got")
	os.Mkdir(".got", os.ModePerm)
	os.Create(".got/staging.csv")

	err := utils.AddEntryToStaging([]string{"nonExistentFile.txt"})
	if err == nil {
		t.Errorf("Expected error for non-existent file, got none")
	}
}

func TestRemoveEntryToStaging_Success(t *testing.T) {
	//Set up test environment
	os.RemoveAll(".got")
	os.Mkdir(".got", os.ModePerm)
	defer os.RemoveAll(".got")
	file1, err := os.Create("file1.txt")
	defer os.Remove("file1.txt")
	defer file1.Close()
	file, _ := os.Create(".got/staging.csv")
	file.WriteString("file1.txt,fbb2e7c65e1506824da06ce6239189b5e50837e8,added\n")
	file.Close()

	err = utils.RemoveEntryToStaging([]string{"file1.txt"})
	if err != nil {
		t.Errorf("Unexpected error : %v", err)
	}
	entries, _ := utils.ReadStagingEntries()
	if len(entries) != 0 {
		t.Errorf("Expected 0 entries, got %d, %v", len(entries), entries)
	}
}

func TestRemoveEntryToStaging_NonExistent(t *testing.T) {
	os.RemoveAll(".got")
	os.Mkdir(".got", os.ModePerm)
	defer os.RemoveAll(".got")
	file, _ := os.Create(".got/staging.csv")
	file.WriteString("file.txt,fbb2e7c65e1506824da06ce6239189b5e50837e8,added\n")
	file.Close()

	err := utils.RemoveEntryToStaging([]string{"nonExistentFile.txt"})
	if err == nil {
		t.Errorf("Expected error for non-existent file, got none")
	}

	entries, _ := utils.ReadStagingEntries()
	if len(entries) != 1 {
		t.Errorf("Expected 1 entry, got %d, %v", len(entries), entries)
	}
}

func TestReadStagingEntries_LargeFile(t *testing.T) {
	os.RemoveAll(".got")
	os.Mkdir(".got", os.ModePerm)
	defer os.RemoveAll(".got")
	file, _ := os.Create(".got/staging.csv")

	for i := 0; i < 100000; i++ {
		file.WriteString("file.txt,fbb2e7c65e1506824da06ce6239189b5e50837e8,added\n")
	}
	file.Close()

	entries, err := utils.ReadStagingEntries()
	if err != nil {
		t.Errorf("Unexpected error : %v", err)
	}
	if len(entries) != 100000 {
		t.Errorf("Expected 100000 entries, got %d, %v", len(entries), entries)
	}
}
