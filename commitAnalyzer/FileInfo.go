package commitAnalyzer

import (
	"Code_Analyzer/gitLog"
	"bufio"
	"github.com/boyter/scc/processor"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	FileId                        string
	Content                       *FileContent
	RenamedTo                     *File
	Removed                       bool
	touchedInCommits              []*gitLog.CommitInfo
	touchedBeforeRenamedInCommits []*gitLog.CommitInfo
}

type FileContent struct {
	Language           string
	LinesOfCode        int64
	LinesOfComments    int64
	LinesBlank         int64
	ComplexityByIndent int
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

func (f *File) addCommitsBeforeRenamed(commits []*gitLog.CommitInfo) {
	f.touchedBeforeRenamedInCommits = append(f.touchedBeforeRenamedInCommits, commits...)
}

func (f *File) ChangeAmount(withHistory bool) int {
	if !withHistory {
		return len(f.touchedInCommits)
	}
	return len(f.touchedBeforeRenamedInCommits) + len(f.touchedInCommits)
}

// GetRecentCommitScore gibt einen Wert zurück, der umso höher ist, je mehr Commits diese Datei in den letzten drei Monaten erlebt hat.
func (f *File) GetRecentCommitScore() int {
	score := 0
	threeMonthsAgo := time.Now().AddDate(0, -3, 0)
	for _, commit := range f.touchedInCommits {
		if commit.AuthorDateTime.After(threeMonthsAgo) {
			score++
		}
	}
	return score
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
	// To work concurrently "processor.ProcessConstants()" must be called centrally beforehand
	processor.CountStats(filejob)
	f.Content.LinesOfCode = filejob.Code
	f.Content.LinesOfComments = filejob.Comment
	f.Content.LinesBlank = filejob.Blank

	complexity, err := f.calcComplexityByIndent()
	if err != nil {
		return
	}
	f.Content.ComplexityByIndent = complexity
}

func (f *File) calcComplexityByIndent() (int, error) {
	file, err := os.Open(f.FileId)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maxIndent := 0

	for scanner.Scan() {
		line := scanner.Text()
		indent := 0
		for _, char := range line {
			if char == ' ' || char == '\t' {
				indent++
			} else {
				break
			}
		}
		if indent > maxIndent {
			maxIndent = indent
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return maxIndent, nil
}
