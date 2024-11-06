package gitLog

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestSetCommitHash(t *testing.T) {
	commitHash := "1234564890"
	commit := CommitInfo{}
	commit.SetCommitHash(commitHash)
	if commit.CommitHash != commitHash {
		t.Errorf("Commit Hash not set correctly")
	}
}

func TestSetAuthor(t *testing.T) {
	authorName := "John Doe"
	commit := CommitInfo{}
	commit.SetAuthor(authorName)
	if commit.Author != authorName {
		t.Errorf("Author not set correctly")
	}
}

func TestSetAuthorDateWrong(t *testing.T) {
	commit := CommitInfo{}
	err := commit.SetAuthorDate("123456789")
	if err == nil {
		t.Errorf("Wrong Date Format not identified")
	}
}

func TestSetAuthorDate(t *testing.T) {
	dateTimeString := "2024-11-05T16:03:57+00:00"
	commit := CommitInfo{}
	err := commit.SetAuthorDate(dateTimeString)
	if err != nil {
		t.Errorf("Date Format wrongly rejected")
	}
	dateTime, _ := time.Parse(time.RFC3339, dateTimeString)
	if commit.AuthorDateTime != dateTime {
		t.Errorf("Author Date not correctly set")
	}
}

func TestSetCommitter(t *testing.T) {
	authorName := "John Doe"
	commit := CommitInfo{}
	commit.SetAuthor(authorName)
	if commit.Author != authorName {
		t.Errorf("Author not set correctly")
	}
}

func TestSetCommitterDateWrong(t *testing.T) {
	commit := CommitInfo{}
	err := commit.SetCommitterDate("123456789")
	if err == nil {
		t.Errorf("Wrong Date Format not identified")
	}
}

func TestSetCommitterDate(t *testing.T) {
	dateTimeString := "2024-11-05T16:03:57+00:00"
	commit := CommitInfo{}
	err := commit.SetCommitterDate(dateTimeString)
	if err != nil {
		t.Errorf("Date Format wrongly rejected")
	}
	dateTime, _ := time.Parse(time.RFC3339, dateTimeString)
	if commit.CommitterDateTime != dateTime {
		t.Errorf("Author Date not correctly set")
	}
}

func TestCommitInfo_AddChangedFileModified(t *testing.T) {
	addLines := 11
	removedLines := 22
	filePath := "myFiles/myFile.go"
	line := strings.Join([]string{strconv.FormatInt(int64(addLines), 10), strconv.FormatInt(int64(removedLines), 10), filePath}, " ")
	commit := CommitInfo{}
	commit.AddChangedFile(line)
	if len(commit.ChangedFiles) != 1 {
		t.Errorf("Expected %d files, but found %d", 1, len(commit.ChangedFiles))
	}
	expected := FileChangeInfo{
		FileName:     filePath,
		LinesAdded:   addLines,
		LinesRemoved: removedLines,
		RenamedFile:  "",
	}
	if !reflect.DeepEqual(commit.ChangedFiles[0], expected) {
		t.Errorf("Expected %v, but got %v", expected, commit.ChangedFiles[0])
	}
}

func TestCommitInfo_AddChangedFileMoved(t *testing.T) {
	addLines := 11
	removedLines := 22
	oldFileName := "myFiles/myFile.go"
	newFileName := "newFolder/otherName.go"
	filePath := strings.Join([]string{oldFileName, "=>", newFileName}, " ")
	line := strings.Join([]string{strconv.FormatInt(int64(addLines), 10), strconv.FormatInt(int64(removedLines), 10), filePath}, " ")
	commit := CommitInfo{}
	commit.AddChangedFile(line)
	if len(commit.ChangedFiles) != 1 {
		t.Errorf("Expected %d files, but found %d", 1, len(commit.ChangedFiles))
	}
	expected := FileChangeInfo{
		FileName:     oldFileName,
		LinesAdded:   addLines,
		LinesRemoved: removedLines,
		RenamedFile:  newFileName,
	}
	if !reflect.DeepEqual(commit.ChangedFiles[0], expected) {
		t.Errorf("Expected %v, but got %v", expected, commit.ChangedFiles[0])
	}
}
