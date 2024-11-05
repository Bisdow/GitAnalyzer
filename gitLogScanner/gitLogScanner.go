package gitLogScanner

type FileChange struct {
	Name  string
	Count int
}

func ScanLog() []FileChangeInfo {
	gitLogs := getGitLogs()
	commits := scanLog(gitLogs)
	return analyze(commits)
}
