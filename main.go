package main

import (
	"Code_Analyzer/gitLogScanner"
	"fmt"
	"sort"
)

func main() {
	fileInfos := gitLogScanner.ScanLog()

	for i, _ := range fileInfos {
		fileInfos[i].AnalyzeFile()
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].GetChangeAmmount() > fileInfos[j].GetChangeAmmount()
	})

	// Output
	fmt.Println("Changes: LoC --- Filename")
	for _, fileInfo := range fileInfos {
		if fileInfo.FileInfo.Removed {
			continue
		}
		fmt.Printf("%d: %d --- %s\n", fileInfo.GetChangeAmmount(), fileInfo.FileInfo.LinesOfCode, fileInfo.FileInfo.FileName)
	}
}
