package main

import (
	"fmt"
	"os"
	"os/exec"
)

func getGitLogs() string {
	cmd := exec.Command("git", "log", "--all", "--numstat", "--date=short", "--pretty=format:--%H--%ad--%aN", "--no-renames", "--since=2000-01-01")
	result, err := cmd.Output()
	if err != nil {
		fmt.Println("Fehler beim Ausführen des git log Befehls:", err)
		os.Exit(1)
	}
	return string(result)
}
