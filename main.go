package main

import (
	"Code_Analyzer/gitLogScanner"
	"Code_Analyzer/loc"
	"fmt"
)

func main() {
	sortedResults := gitLogScanner.ScanLog()

	fileInfos := loc.CountLinesOfCode(sortedResults)

	// Output
	fmt.Println("Changes: LoC --- Filename")
	for _, change := range fileInfos {
		fmt.Printf("%d: %d --- %s\n", change.Changes, change.Loc, change.Name)
	}
}
