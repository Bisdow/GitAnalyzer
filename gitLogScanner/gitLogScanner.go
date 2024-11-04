package gitLogScanner

type FileChange struct {
	Name  string
	Count int
}

func ScanLog() []FileChange {
	gitLogs := getGitLogs()
	commits := scanLog(gitLogs)
	return analyzeCommits(commits)
}
