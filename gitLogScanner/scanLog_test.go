package gitLogScanner

import (
	"reflect"
	"testing"
)

func TestScanLogsSingleCommit(t *testing.T) {
	logs := `--abc123--2023-10-01--John Doe
10	5	file1.txt
`

	expectedCommits := []commit{
		{
			author:       "John Doe",
			date:         "2023-10-01",
			commitHash:   "abc123",
			changedFiles: []string{"file1.txt"},
		},
	}

	commits := scanLog(logs)

	if !reflect.DeepEqual(commits, expectedCommits) {
		t.Errorf("scanLog() = %v, want %v", commits, expectedCommits)
	}
}

func TestScanLogs(t *testing.T) {
	logs := `--abc123--2023-10-01--John Doe
10	5	file1.txt
2	1	file2.txt

--def456--2023-10-02--Jane Smith
3	0	file3.txt
2	1	file2.txt
`

	expectedCommits := []commit{
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
			changedFiles: []string{"file3.txt", "file2.txt"},
		},
	}

	commits := scanLog(logs)

	if !reflect.DeepEqual(commits, expectedCommits) {
		t.Errorf("scanLog() = %v, want %v", commits, expectedCommits)
	}
}
