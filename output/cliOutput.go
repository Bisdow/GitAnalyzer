package output

import (
	"Code_Analyzer/commitAnalyzer"
	"fmt"
	"slices"
)

func CliOutput(filesMap map[string]commitAnalyzer.File) {
	var files []commitAnalyzer.File
	for _, file := range filesMap {
		if file.Removed {
			continue
		}
		files = append(files, file)
	}

	slices.SortFunc(files, func(a, b commitAnalyzer.File) int {
		return b.ChangeAmount(true) - a.ChangeAmount(true)
	})

	fmt.Printf("Changes : Lines of Code  === Filename\n")
	for _, file := range files {
		fmt.Printf("%d : %d === %s\n", file.ChangeAmount(true), file.Content.LinesOfCode, file.GetFilename())
	}
}
