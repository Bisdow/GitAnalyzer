package gitLogScanner

import "testing"

func TestFileChangeInfo_addCommitInfo(t *testing.T) {
	commit := CommitInfo{
		CommitHash: "abc123",
		Author:     "John Doe",
		Date:       "2023-10-01",
	}

	fileChangeInfo := FileChangeInfo{}
	fileChangeInfo.AddCommitInfo(commit)

	if len(fileChangeInfo.Commits) != 1 {
		t.Errorf("Expected 1 commit, got %d", len(fileChangeInfo.Commits))
	}

	if fileChangeInfo.Commits[0] != commit {
		t.Errorf("Expected commit %v, got %v", commit, fileChangeInfo.Commits[0])
	}
}

func TestFileChangeInfo_GetChangeAmmount(t *testing.T) {
	fileChangeInfo := FileChangeInfo{
		Commits: []CommitInfo{
			{CommitHash: "abc123", Author: "John Doe", Date: "2023-10-01"},
			{CommitHash: "def456", Author: "Jane Smith", Date: "2023-10-02"},
		},
	}

	if got := fileChangeInfo.GetChangeAmmount(); got != 2 {
		t.Errorf("Expected 2 commits, got %d", got)
	}
}
