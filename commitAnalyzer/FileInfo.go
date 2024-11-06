package commitAnalyzer

import (
	"Code_Analyzer/gitLog"
	"github.com/boyter/scc/processor"
	"os"
	"path/filepath"
)

type File struct {
	FileId           string
	Content          *FileContent
	RenamedTo        *File
	Removed          bool
	touchedInCommits []*gitLog.CommitInfo
}

type FileContent struct {
	Language        string
	LinesOfCode     int64
	LinesOfComments int64
	LinesBlank      int64
}

func NewFile(fileId string) *File {
	return &File{FileId: fileId}
}

func (f *File) SetId(fileId string) {
	f.FileId = fileId
}

func (f *File) GetExtension() string {
	return filepath.Ext(f.FileId)
}

func (f *File) GetFilename() string {
	return filepath.Base(f.FileId)
}

func (f *File) GetPath() string {
	return filepath.Dir(f.FileId)
}

func (f *File) SetRenamedTo(file *File) {
	f.RenamedTo = file
}

func (f *File) addCommit(commit *gitLog.CommitInfo) {
	f.touchedInCommits = append(f.touchedInCommits, commit)
}

var readFile = os.ReadFile

func (f *File) AnalyzeContent() {
	content, err := readFile(f.FileId)
	if err != nil {
		f.Removed = true
		return
	}

	if f.Content == nil {
		f.Content = &FileContent{}
	}

	filejob := &processor.FileJob{
		Filename: f.FileId,
		Content:  content,
		Bytes:    int64(len(content)),
	}
	processor.ProcessConstants()
	processor.CountStats(filejob)
	f.Content.LinesOfCode = filejob.Code
	f.Content.LinesOfComments = filejob.Comment
	f.Content.LinesBlank = filejob.Blank
}
