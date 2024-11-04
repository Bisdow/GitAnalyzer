package loc

import (
	"fmt"
	"github.com/boyter/scc/processor"
	"os"
)

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
