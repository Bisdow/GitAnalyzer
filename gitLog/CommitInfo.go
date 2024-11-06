package gitLog

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type FileChangeInfo struct {
	FileName     string
	LinesAdded   int
	LinesRemoved int
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
	elements := strings.Split(line, " ")
	linesAdded, err := strconv.ParseInt(elements[0], 10, 64)
	if err != nil {
		return errors.New(elements[0] + " is not a number")
	}
	cf.LinesAdded = int(linesAdded)

	linesRemoved, err := strconv.ParseInt(elements[1], 10, 64)
	if err != nil {
		return errors.New(elements[1] + " is not a number")
	}
	cf.LinesRemoved = int(linesRemoved)
	if len(elements) == 3 {
		// created / deleted or modified file
		cf.FileName = elements[2]
	} else if len(elements) == 5 {
		// moved or renamed file
		cf.FileName = elements[2]
		cf.RenamedFile = elements[4]
	} else {
		return errors.New("Unsupported Format of: " + line)
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
