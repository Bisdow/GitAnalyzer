package main

import (
	"bufio"
	"strings"
)

func scanLog(log string) []Commit {
	var commits []Commit
	scanner := bufio.NewScanner(strings.NewReader(log))

	var currentCommit Commit

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "--") {
			// Commit header
			currentCommit = Commit{}
			commitParts := strings.Split(line, "--")
			currentCommit.commitHash = commitParts[1]
			currentCommit.date = commitParts[2]
			currentCommit.author = commitParts[3]
			continue
		}
		if line != "" {
			// Commit changed file information
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
