package main

import (
	"fmt"
)

type Commit struct {
	author       string
	date         string
	commitHash   string
	changedFiles []string
}

type FileChange struct {
	name  string
	count int
}

func main() {
	gitLogs := getGitLogs()
	commits := scanLog(gitLogs)
	sortedResults :=
		analyzeCommits(commits)

	// Output
	for _, change := range sortedResults {
		fmt.Printf("%d: %s\n", change.count, change.name)
	}
}
