package gitLog

import (
	"bufio"
	"fmt"
	"strings"
)

const headerSeperator = "<|>"

func extractCommits(log string) ([]CommitInfo, error) {
	scanner := bufio.NewScanner(strings.NewReader(log))

	var commits []CommitInfo
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, headerSeperator) {
			// commit header
			commit, err := readCommitHeader(line)
			if err != nil {
				return commits, fmt.Errorf("Headerline of Commit is not processable: \n %s \n because or Error \n %s", line, err)
			}
			commits = append(commits, commit)
			continue
		} else {
			// commit changed file information
			if err := lastOf(commits).AddChangedFile(line); err != nil {
				return commits, fmt.Errorf("ChangedFile Information of Commit is not processable \n"+
					"CommitHash: %s\n"+
					"ChangedFile line: \n%s\n"+
					"Error:\n"+
					"%s",
					lastOf(commits).CommitHash, line, err)
			}
			continue
		}
	}
	return commits, nil
}

func lastOf(list []CommitInfo) *CommitInfo {
	return &list[len(list)-1]
}

func readCommitHeader(line string) (CommitInfo, error) {
	// Format
	// <|>{{CommitHash}}<|>{{AuthorDateTime}}<|>{{AuthorName}}<|>{{CommiterDateTime}}<|>{{CommitterName}}<|>{{CommitMessage}}<|>
	commit := CommitInfo{}
	commitParts := strings.Split(line, headerSeperator)

	commit.SetCommitHash(commitParts[1])
	if err := commit.SetAuthorDate(commitParts[2]); err != nil {
		return commit, err
	}
	commit.SetAuthor(commitParts[3])
	if err := commit.SetCommitterDate(commitParts[4]); err != nil {
		return commit, err
	}
	commit.SetCommitter(commitParts[5])
	return commit, nil
}
