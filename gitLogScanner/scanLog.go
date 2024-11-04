package gitLogScanner

import (
	"bufio"
	"strings"
)

type commit struct {
	author       string
	date         string
	commitHash   string
	changedFiles []string
}

func scanLog(log string) []commit {
	var commits []commit
	scanner := bufio.NewScanner(strings.NewReader(log))

	var currentCommit commit

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "--") {
			// commit header
			currentCommit = commit{}
			commitParts := strings.Split(line, "--")
			currentCommit.commitHash = commitParts[1]
			currentCommit.date = commitParts[2]
			currentCommit.author = commitParts[3]
			continue
		}
		if line != "" {
			// commit changed file information
			changedFileParts := strings.Split(line, "\t")
			currentCommit.changedFiles = append(currentCommit.changedFiles, changedFileParts[2])
			continue
		}
		// Empty line
		if currentCommit.commitHash != "" {
			commits = append(commits, currentCommit)
		}

	}
	commits = append(commits, currentCommit)
	return commits
}
