package gitLogScanner

func analyze(commits []commit) []FileChangeInfo {
	fileChangeInfosMap := make(map[string]FileChangeInfo)
	for _, commit := range commits {
		for _, file := range commit.changedFiles {
			if _, ok := fileChangeInfosMap[file]; !ok {
				fileChangeInfosMap[file] = FileChangeInfo{FileInfo{FileName: file}, []CommitInfo{}}
			}
			tmpChangeInfo := fileChangeInfosMap[file]
			tmpChangeInfo.AddCommitInfo(CommitInfo{commit.commitHash, commit.author, commit.date})
			fileChangeInfosMap[file] = tmpChangeInfo
		}
	}

	var fileChangeInfos []FileChangeInfo
	for _, info := range fileChangeInfosMap {
		fileChangeInfos = append(fileChangeInfos, info)
	}

	return fileChangeInfos
}
