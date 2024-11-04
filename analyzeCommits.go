package main

import "sort"

func analyzeCommits(commits []Commit) []FileChange {
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
		sortedChanges = append(sortedChanges, FileChange{name, count})
	}
	sort.Slice(sortedChanges, func(i, j int) bool {
		return sortedChanges[i].count > sortedChanges[j].count
	})

	return sortedChanges
}
