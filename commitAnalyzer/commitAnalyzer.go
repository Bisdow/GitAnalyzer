package commitAnalyzer

import (
	"Code_Analyzer/gitLog"
	"slices"
)

func CommitAnalyzer(gitLogCommits []gitLog.CommitInfo) map[string]File {
	commits := make(map[string]gitLog.CommitInfo)
	for _, commit := range gitLogCommits {
		commits[commit.CommitHash] = commit
	}

	files := make(map[string]File)

	// Sorts from oldest commit to newest
	slices.SortFunc(gitLogCommits, func(a, b gitLog.CommitInfo) int {
		return a.AuthorDateTime.Compare(b.AuthorDateTime)
	})

	for _, commit := range gitLogCommits {
		analyzeCommit(&commit, &files)
	}
	return files
}

func analyzeCommit(commit *gitLog.CommitInfo, files *map[string]File) {
	for _, gitFile := range commit.ChangedFiles {
		currentFile := getFile(gitFile.FileName, files)
		currentFile.AnalyzeContent()
		currentFile.addCommit(commit)
		if gitFile.RenamedFile != "" {
			renamedFile := getFile(gitFile.RenamedFile, files)
			renamedFile.addCommit(commit)
			renamedFile.AnalyzeContent()
			(*files)[renamedFile.FileId] = renamedFile
			currentFile.SetRenamedTo(&renamedFile)
			currentFile.addCommitsBeforeRenamed(renamedFile.touchedBeforeRenamedInCommits)
			currentFile.addCommitsBeforeRenamed(renamedFile.touchedInCommits)
		}
		(*files)[currentFile.FileId] = currentFile
	}
}

func getFile(fileId string, files *map[string]File) File {
	if _, isKnown := (*files)[fileId]; !isKnown {
		(*files)[fileId] = *NewFile(fileId)
	}
	return (*files)[fileId]
}
