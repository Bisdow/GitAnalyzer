package commitAnalyzer

import (
	"errors"
	"github.com/boyter/scc/processor"
	"os"
	"reflect"
	"testing"
)

func TestFile_SetId(t *testing.T) {
	id := "hiho/wakawka.go"
	underTest := File{}
	underTest.SetId(id)
	if underTest.FileId != id {
		t.Errorf("File.Id\n"+
			"Expect : %s\n"+
			"Current: %s", id, underTest.FileId)
	}
}

func TestFile_GetExtension(t *testing.T) {
	id := "hiho/wakawka.go"
	underTest := NewFile(id)
	if underTest.GetExtension() != ".go" {
		t.Errorf("File.GetExtension\n"+
			"Expect : .go\n"+
			"Current: %s", underTest.GetExtension())
	}
}

func TestFile_GetFielename(t *testing.T) {
	id := "hiho/wakawka.go"
	underTest := NewFile(id)
	if underTest.GetFilename() != "wakawka.go" {
		t.Errorf("File.GetFilename\n"+
			"Expect : wakawka.go\n"+
			"Current: %s", underTest.GetFilename())
	}
}

func TestFile_GetPath(t *testing.T) {
	id := "path/subpath/wakawka.go"
	underTest := NewFile(id)
	if underTest.GetPath() != "path/subpath" {
		t.Errorf("File.GetPath\n"+
			"Expect : path/subpath\n"+
			"Current: %s", underTest.GetPath())
	}
}

// Mock-Funktion für os.ReadFile
func mockReadFile(filename string) ([]byte, error) {
	switch filename {
	case "someFile.txt":
		return []byte("package main\n" +
			"func x(){}\n" +
			"\n"), nil
	case "removedFile.del":
		return nil, errors.New("file not found")
	}
	return nil, errors.New("No Mock defined for this fileId")
}

func TestFile_SetRenamedTo(t *testing.T) {
	underTest := NewFile("someFile.txt")
	renamedTo := NewFile("newFile.txt")
	underTest.SetRenamedTo(renamedTo)
	if underTest.RenamedTo != renamedTo {
		t.Errorf("File.SetRenamedTo\n"+
			"Expect : %p\n"+
			"Current: %p", renamedTo, underTest.RenamedTo)
	}
}

func TestFile_AnalyzeContent(t *testing.T) {
	// Mocken von os.ReadFile
	readFile = mockReadFile
	defer func() { readFile = os.ReadFile }()
	// Must be initialized before calling AnalyzeContent
	processor.ProcessConstants()

	underTest := NewFile("someFile.txt")
	underTest.AnalyzeContent()
	expected := FileContent{
		LinesOfCode:     2,
		LinesOfComments: 0,
		LinesBlank:      1,
	}
	if !reflect.DeepEqual(&expected, underTest.Content) {
		t.Errorf("File.AnalyzeContent\n"+
			"Expect : %+v\n"+
			"Current: %+v", &expected, underTest.Content)
	}
}
