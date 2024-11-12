package main

import (
	"Code_Analyzer/commitAnalyzer"
	"Code_Analyzer/gitLog"
	"Code_Analyzer/output"
	"fmt"
	"os"
)

func main() {
	commits, err := gitLog.GetCommits()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	files := commitAnalyzer.CommitAnalyzer(commits)
	output.CliOutput(files)
}
