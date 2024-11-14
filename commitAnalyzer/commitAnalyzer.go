package commitAnalyzer

import (
	"Code_Analyzer/gitLog"
	"github.com/boyter/scc/processor"
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

	collectFiles(gitLogCommits, &files)
	analyzeFilesConcurrently(&files)
	return files
}

func analyzeFilesConcurrently(files *map[string]File) {
	processor.ProcessConstants()
	result := make(chan File)
	for _, file := range *files {
		go analyzeFile(file, result)
	}
	for range *files {
		file := <-result
		(*files)[file.FileId] = file
	}
}

func analyzeFile(file File, result chan (File)) {
	file.AnalyzeContent()
	result <- file
}

func collectFiles(gitLogCommits []gitLog.CommitInfo, files *map[string]File) {
	for _, commit := range gitLogCommits {
		for _, gitFile := range commit.ChangedFiles {
			currentFile := getFile(gitFile.FileName, files)
			currentFile.addCommit(&commit)
			if gitFile.RenamedFile != "" {
				renamedFile := getFile(gitFile.RenamedFile, files)
				renamedFile.addCommit(&commit)
				(*files)[renamedFile.FileId] = renamedFile
				currentFile.SetRenamedTo(&renamedFile)
				currentFile.addCommitsBeforeRenamed(renamedFile.touchedBeforeRenamedInCommits)
				currentFile.addCommitsBeforeRenamed(renamedFile.touchedInCommits)
			}
			(*files)[currentFile.FileId] = currentFile
		}
	}
}

func getFile(fileId string, files *map[string]File) File {
	if _, isKnown := (*files)[fileId]; !isKnown {
		(*files)[fileId] = *NewFile(fileId)
	}
	return (*files)[fileId]
}
