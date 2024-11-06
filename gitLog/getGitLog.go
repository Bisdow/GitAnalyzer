package gitLog

import (
	"fmt"
	"os"
	"os/exec"
)

func getGitLogs() string {
	cmd := exec.Command("git", "log", "--all", "--numstat", "--pretty=format:<|>%H<|>%aI<|>%aN<|>cI<|>cN<|>%s<|>", "--no-renames", "--since=1900-01-01")
	result, err := cmd.Output()
	if err != nil {
		fmt.Println("Fehler beim Ausführen des git log Befehls:", err)
		os.Exit(1)
	}
	return string(result)
}
