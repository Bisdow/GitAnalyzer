package gitLogScanner

import (
	"os"
	"reflect"
	"testing"
)

func TestAnalyzeCommitsValidInput(t *testing.T) {
	// Mock os.Stat
	checkFileExists = func(name string) (os.FileInfo, error) {
		return nil, nil
	}
	defer func() { checkFileExists = os.Stat }()

	commits := []commit{
		{
			author:       "John Doe",
			date:         "2023-10-01",
			commitHash:   "abc123",
			changedFiles: []string{"file1.txt", "file2.txt"},
		},
		{
			author:       "Jane Smith",
			date:         "2023-10-02",
			commitHash:   "def456",
			changedFiles: []string{"file2.txt"},
		},
	}

	expectedOutput := []FileChange{
		{Name: "file2.txt", Count: 2},
		{Name: "file1.txt", Count: 1},
	}

	fileChanges := analyzeCommits(commits)

	if !reflect.DeepEqual(fileChanges, expectedOutput) {
		t.Errorf("analyzeCommits() = %v, want %v", fileChanges, expectedOutput)
	}
}

func TestAnalyzeCommitsEmptyInput(t *testing.T) {
	// Mock os.Stat
	checkFileExists = func(name string) (os.FileInfo, error) {
		return nil, nil
	}
	defer func() { checkFileExists = os.Stat }()

	commits := []commit{}

	expectedOutput := []FileChange{}

	fileChanges := analyzeCommits(commits)

	if !reflect.DeepEqual(fileChanges, expectedOutput) {
		t.Errorf("analyzeCommits() = %v, want %v", fileChanges, expectedOutput)
	}
}

func TestAnalyzeCommitsSingleCommit(t *testing.T) {
	// Mock os.Stat
	checkFileExists = func(name string) (os.FileInfo, error) {
		return nil, nil
	}
	defer func() { checkFileExists = os.Stat }()

	commits := []commit{
		{
			author:       "John Doe",
			date:         "2023-10-01",
			commitHash:   "abc123",
			changedFiles: []string{"file1.txt"},
		},
	}

	expectedOutput := []FileChange{
		{Name: "file1.txt", Count: 1},
	}

	fileChanges := analyzeCommits(commits)

	if !reflect.DeepEqual(fileChanges, expectedOutput) {
		t.Errorf("analyzeCommits() = %v, want %v", fileChanges, expectedOutput)
	}
}

func TestAnalyzeCommitsMultipleCommitsSameFile(t *testing.T) {
	// Mock os.Stat
	checkFileExists = func(name string) (os.FileInfo, error) {
		return nil, nil
	}
	defer func() { checkFileExists = os.Stat }()

	commits := []commit{
		{
			author:       "John Doe",
			date:         "2023-10-01",
			commitHash:   "abc123",
			changedFiles: []string{"file1.txt"},
		},
		{
			author:       "Jane Smith",
			date:         "2023-10-02",
			commitHash:   "def456",
			changedFiles: []string{"file1.txt"},
		},
	}

	expectedOutput := []FileChange{
		{Name: "file1.txt", Count: 2},
	}

	fileChanges := analyzeCommits(commits)

	if !reflect.DeepEqual(fileChanges, expectedOutput) {
		t.Errorf("analyzeCommits() = %v, want %v", fileChanges, expectedOutput)
	}
}
