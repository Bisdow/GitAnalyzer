package main

import (
	"reflect"
	"testing"
)

func TestAnalyzeCommitsValidInput(t *testing.T) {
	commits := []Commit{
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
		{name: "file2.txt", count: 2},
		{name: "file1.txt", count: 1},
	}

	fileChanges := analyzeCommits(commits)

	if !reflect.DeepEqual(fileChanges, expectedOutput) {
		t.Errorf("analyzeCommits() = %v, want %v", fileChanges, expectedOutput)
	}
}

func TestAnalyzeCommitsEmptyInput(t *testing.T) {
	commits := []Commit{}

	expectedOutput := []FileChange{}

	fileChanges := analyzeCommits(commits)

	if !reflect.DeepEqual(fileChanges, expectedOutput) {
		t.Errorf("analyzeCommits() = %v, want %v", fileChanges, expectedOutput)
	}
}

func TestAnalyzeCommitsSingleCommit(t *testing.T) {
	commits := []Commit{
		{
			author:       "John Doe",
			date:         "2023-10-01",
			commitHash:   "abc123",
			changedFiles: []string{"file1.txt"},
		},
	}

	expectedOutput := []FileChange{
		{name: "file1.txt", count: 1},
	}

	fileChanges := analyzeCommits(commits)

	if !reflect.DeepEqual(fileChanges, expectedOutput) {
		t.Errorf("analyzeCommits() = %v, want %v", fileChanges, expectedOutput)
	}
}

func TestAnalyzeCommitsMultipleCommitsSameFile(t *testing.T) {
	commits := []Commit{
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
		{name: "file1.txt", count: 2},
	}

	fileChanges := analyzeCommits(commits)

	if !reflect.DeepEqual(fileChanges, expectedOutput) {
		t.Errorf("analyzeCommits() = %v, want %v", fileChanges, expectedOutput)
	}
}
