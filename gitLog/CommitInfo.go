package gitLog

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const fileLineInfoSeparator = "\t"
const fileRenameSeparator = " => "

type FileChangeInfo struct {
	FileName     string
	LinesAdded   *int
	LinesRemoved *int
	RenamedFile  string
}

type CommitInfo struct {
	CommitHash        string
	Author            string
	AuthorDateTime    time.Time
	Committer         string
	CommitterDateTime time.Time
	ChangedFiles      []FileChangeInfo
}

func (c *CommitInfo) SetCommitHash(hash string) {
	c.CommitHash = hash
}

func (c *CommitInfo) SetAuthor(authorName string) {
	c.Author = authorName
}

func (c *CommitInfo) SetAuthorDate(authorDateTime string) error {
	// Expected Format 2024-11-04T17:02:55+01:00
	result, err := parseDateTime(authorDateTime)
	c.AuthorDateTime = result
	return err
}

func (c *CommitInfo) SetCommitter(commiterName string) {
	c.Committer = commiterName
}

func (c *CommitInfo) SetCommitterDate(committerDateTime string) error {
	// Expected Format 2024-11-04T17:02:55+01:00
	result, err := parseDateTime(committerDateTime)
	c.CommitterDateTime = result
	return err
}

func (c *CommitInfo) AddChangedFile(line string) error {
	/* Expected formats
	added Lines -- removed lines -- filepath
	40      0       gitLogScanner/analyzeCommits.go
	or, if the file was moved or renamed
	9       9       analyzeCommits_test.go => gitLogScanner/analyzeCommits_test.go
	*/
	cf := FileChangeInfo{}
	elements := strings.Split(line, fileLineInfoSeparator)

	linesAdded, err := parseLineAmount(elements[0])
	if err != nil {
		return err
	}
	cf.LinesAdded = linesAdded

	linesRemoved, err := parseLineAmount(elements[1])
	if err != nil {
		return err
	}
	cf.LinesRemoved = linesRemoved

	if !strings.Contains(elements[2], fileRenameSeparator) {
		// created / deleted or modified file
		cf.FileName = elements[2]
	} else {
		// moved or renamed file
		fileNames := strings.Split(elements[2], fileRenameSeparator)
		cf.FileName = fileNames[0]
		cf.RenamedFile = fileNames[1]
	}
	c.ChangedFiles = append(c.ChangedFiles, cf)
	return nil
}

func parseDateTime(dateTime string) (time.Time, error) {
	result, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		err := errors.New("The date " + dateTime + " is not a valid Date")
		return result, err
	}
	return result, nil
}

func parseLineAmount(str string) (*int, error) {
	if str == "-" {
		return nil, nil
	}
	lines, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return nil, errors.New(str + " is not a number")
	}
	linesInt := int(lines)
	return &linesInt, nil
}
