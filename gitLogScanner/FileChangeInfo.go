package gitLogScanner

import (
	"github.com/boyter/scc/processor"
	"os"
)

type FileInfo struct {
	FileName    string
	Removed     bool
	Language    string
	LinesOfCode int64
	BlankLines  int64
	Comment     int64
}

type CommitInfo struct {
	CommitHash string
	Author     string
	Date       string
}

type FileChangeInfo struct {
	FileInfo FileInfo
	Commits  []CommitInfo
}

func (f FileChangeInfo) GetChangeAmmount() int64 {
	return int64(len(f.Commits))
}

func (f FileChangeInfo) GetAuthorAmmount() int64 {
	authors := make(map[string]bool)
	for _, commit := range f.Commits {
		authors[commit.Author] = true
	}
	return int64(len(authors))
}

func (f *FileChangeInfo) AddCommitInfo(commit CommitInfo) {
	f.Commits = append(f.Commits, commit)
}

func (f *FileChangeInfo) AnalyzeFile() {
	content, err := os.ReadFile(f.FileInfo.FileName)
	if err != nil {
		f.FileInfo.Removed = true
		return
	}

	filejob := &processor.FileJob{
		Filename: f.FileInfo.FileName,
		Content:  content,
		Bytes:    int64(len(content)),
	}
	processor.ProcessConstants()
	processor.CountStats(filejob)
	f.FileInfo.LinesOfCode = filejob.Code
	f.FileInfo.Comment = filejob.Comment
	f.FileInfo.BlankLines = filejob.Blank
}
