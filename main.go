package main

import (
	"Code_Analyzer/gitLogScanner"
	"fmt"
	"github.com/boyter/scc/processor"
	"os"
)

func main() {
	sortedResults := gitLogScanner.ScanLog()

	fileInfos := countLinesOfCode(sortedResults)

	// Output
	fmt.Println("Changes: LoC --- Filename")
	for _, change := range fileInfos {
		fmt.Printf("%d: %d --- %s\n", change.changes, change.loc, change.name)
	}
}

type FileInfo struct {
	name    string
	changes int
	loc     int64
}

func countLinesOfCode(files []gitLogScanner.FileChange) []FileInfo {
	var fileInfos []FileInfo

	for _, file := range files {
		currentFileInfo := FileInfo{
			name:    file.Name,
			changes: file.Count,
			loc:     loc(file.Name),
		}
		fileInfos = append(fileInfos, currentFileInfo)
	}
	return fileInfos
}

func loc(fileName string) int64 {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Datei konnte nicht geöffnet werden")
		os.Exit(1)
	}

	filejob := &processor.FileJob{
		Filename: fileName,
		Content:  content,
		Bytes:    int64(len(content)),
	}

	processor.ProcessConstants()
	processor.CountStats(filejob)
	return filejob.Code
}
