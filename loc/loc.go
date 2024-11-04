package loc

import (
	"Code_Analyzer/gitLogScanner"
)

type FileInfo struct {
	Name    string
	Changes int
	Loc     int64
}

func CountLinesOfCode(files []gitLogScanner.FileChange) []FileInfo {
	var fileInfos []FileInfo

	for _, file := range files {
		currentFileInfo := FileInfo{
			Name:    file.Name,
			Changes: file.Count,
			Loc:     loc(file.Name),
		}
		fileInfos = append(fileInfos, currentFileInfo)
	}
	return fileInfos
}
