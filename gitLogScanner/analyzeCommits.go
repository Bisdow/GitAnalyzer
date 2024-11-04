package gitLogScanner

import (
	"os"
	"sort"
)

var checkFileExists = os.Stat

func analyzeCommits(commits []commit) []FileChange {
	// Collect
	fileChanges := make(map[string]int)
	for _, commit := range commits {
		for _, file := range commit.changedFiles {
			fileChanges[file]++
		}
	}

	// Sorting
	sortedChanges := []FileChange{}
	for name, count := range fileChanges {
		if _, err := checkFileExists(name); err != nil {
			continue
		}
		sortedChanges = append(sortedChanges, FileChange{name, count})
	}
	sort.Slice(sortedChanges, func(i, j int) bool {
		return sortedChanges[i].Count > sortedChanges[j].Count
	})

	return sortedChanges
}
