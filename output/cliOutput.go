package output

import (
	"Code_Analyzer/commitAnalyzer"
	"fmt"
	"slices"
)

var NormalizedAmountLimit float64 = 0.1
var RecentCommitScoreLimitPercentage float64 = 0.1

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

	maxChanges := files[0].ChangeAmount(true)
	maxRecentCommitScore := maxRecentCommitScore(files)

	fmt.Println("======================================================================================")
	fmt.Printf("Change-Quote (absolute Changes) === RecentCommitScore === ComplexityByIndent === Lines of Code  === Filename\n")
	fmt.Println("======================================================================================")
	for _, file := range files {
		currentAmmount := file.ChangeAmount(true)
		normalizedAmount := float64(currentAmmount) / float64(maxChanges)
		recentCommitScore := file.GetRecentCommitScore()

		if normalizedAmount < NormalizedAmountLimit && float64(recentCommitScore) < float64(maxRecentCommitScore)*RecentCommitScoreLimitPercentage {
			break
		}

		fmt.Printf("Change-Quote %1.2f ( %5d ) === RCS %4d === CBI %4d === LoC %6d === %s\n", normalizedAmount, currentAmmount, recentCommitScore, file.Content.ComplexityByIndent, file.Content.LinesOfCode, file.FileId)

	}
}

func maxRecentCommitScore(files []commitAnalyzer.File) int {
	maxScore := 0

	for _, file := range files {
		if file.Removed {
			continue
		}
		if file.GetRecentCommitScore() > maxScore {
			maxScore = file.GetRecentCommitScore()
		}
	}
	return maxScore
}
