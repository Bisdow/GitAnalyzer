package commitAnalyzer

import (
	"Code_Analyzer/gitLog"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestAnalyzeCommit(t *testing.T) {
	// Mocken von os.ReadFile
	readFile = mockReadFile
	defer func() { readFile = os.ReadFile }()

	files := make(map[string]File)

	commit := gitLog.CommitInfo{
		CommitHash:        "123",
		Author:            "author",
		AuthorDateTime:    time.Now(),
		Committer:         "committer",
		CommitterDateTime: time.Now(),
		ChangedFiles: []gitLog.FileChangeInfo{
			{
				FileName:     "oldFileName.txt",
				LinesAdded:   10,
				LinesRemoved: 5,
				RenamedFile:  "someFile.txt",
			},
		},
	}

	analyzeCommit(&commit, &files)

	expectedFiles := map[string]File{}
	renamedFile := File{
		FileId: commit.ChangedFiles[0].RenamedFile,
		Content: &FileContent{
			Language:        "",
			LinesOfCode:     2,
			LinesBlank:      1,
			LinesOfComments: 0,
		},
		RenamedTo:        nil,
		Removed:          false,
		touchedInCommits: []*gitLog.CommitInfo{&commit},
	}
	originalFile := File{
		FileId:           commit.ChangedFiles[0].FileName,
		Content:          nil,
		RenamedTo:        &renamedFile,
		Removed:          true,
		touchedInCommits: []*gitLog.CommitInfo{&commit},
	}

	expectedFiles[commit.ChangedFiles[0].RenamedFile] = renamedFile
	expectedFiles[commit.ChangedFiles[0].FileName] = originalFile

	if !reflect.DeepEqual(expectedFiles, files) {
		t.Errorf("analyzeCommit:\n"+
			"Expected: %v\n"+
			"Current : %v", expectedFiles, files)
	}
}
