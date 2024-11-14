package commitAnalyzer

import (
	"Code_Analyzer/gitLog"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestCollectFiles(t *testing.T) {
	// Mocken von os.ReadFile
	readFile = mockReadFile
	defer func() { readFile = os.ReadFile }()

	files := make(map[string]File)

	addLines := int(10)
	removedLines := int(5)

	commit := gitLog.CommitInfo{
		CommitHash:        "123",
		Author:            "author",
		AuthorDateTime:    time.Now(),
		Committer:         "committer",
		CommitterDateTime: time.Now(),
		ChangedFiles: []gitLog.FileChangeInfo{
			{
				FileName:     "oldFileName.txt",
				LinesAdded:   &addLines,
				LinesRemoved: &removedLines,
				RenamedFile:  "someFile.txt",
			},
		},
	}
	commits := []gitLog.CommitInfo{commit}

	collectFiles(commits, &files)

	expectedFiles := map[string]File{}
	renamedFile := File{
		FileId:           commit.ChangedFiles[0].RenamedFile,
		Content:          nil,
		RenamedTo:        nil,
		Removed:          false,
		touchedInCommits: []*gitLog.CommitInfo{&commit},
	}
	originalFile := File{
		FileId:                        commit.ChangedFiles[0].FileName,
		Content:                       nil,
		RenamedTo:                     &renamedFile,
		Removed:                       false,
		touchedInCommits:              []*gitLog.CommitInfo{&commit},
		touchedBeforeRenamedInCommits: []*gitLog.CommitInfo{&commit},
	}
	expectedFiles[commit.ChangedFiles[0].RenamedFile] = renamedFile
	expectedFiles[commit.ChangedFiles[0].FileName] = originalFile

	if !reflect.DeepEqual(expectedFiles, files) {
		t.Errorf("analyzeCommit:\n"+
			"Expected: %v\n"+
			"Current : %v", expectedFiles, files)
	}
}
