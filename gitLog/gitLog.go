package gitLog

func GetCommits() ([]CommitInfo, error) {
	gitLogs := getGitLogs()
	return extractCommits(gitLogs)
}
