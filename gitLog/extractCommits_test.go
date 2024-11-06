package gitLog

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestExtractCommits(t *testing.T) {
	hash := "0123456789"
	author := "John Doe"
	authorDate := "2020-01-01T00:00:00+01:00"
	committer := "Jane Doe"
	committerDate := "2000-12-31T23:59:59+01:00"
	commitMessage := "And the world is round"
	changedFileModified := []string{"12", "0", "myFolder/Added.go"}
	changedFileMoved := []string{"1", "2", "myFolder/moved.go", "=>", "newFolder/newName"}

	log := strings.Join([]string{"", hash, authorDate, author, committerDate, committer, commitMessage, ""}, headerSeperator)
	log = log + "\n" + strings.Join(changedFileModified, " ")
	log = log + "\n" + strings.Join(changedFileMoved, " ")

	commits, err := extractCommits(log)
	if err != nil {
		t.Errorf("Error on extractCommits: \n%s", err)
	}

	authorDateTime, _ := time.Parse(time.RFC3339, authorDate)
	committerDateTime, _ := time.Parse(time.RFC3339, committerDate)
	expected := []CommitInfo{
		{
			CommitHash:        hash,
			Author:            author,
			AuthorDateTime:    authorDateTime,
			Committer:         committer,
			CommitterDateTime: committerDateTime,
			ChangedFiles: []FileChangeInfo{
				{
					FileName:     changedFileModified[2],
					LinesAdded:   12,
					LinesRemoved: 0,
					RenamedFile:  "",
				},
				{
					FileName:     changedFileMoved[2],
					LinesAdded:   1,
					LinesRemoved: 2,
					RenamedFile:  changedFileMoved[4],
				},
			},
		},
	}

	if !reflect.DeepEqual(commits, expected) {
		t.Errorf("Not as expected:\n"+
			"Expected: %v\n"+
			"Current : %v", expected, commits)
	}
}
