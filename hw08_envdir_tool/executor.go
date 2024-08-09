package main

import (
	"os"
	"os/exec"
	"strings"
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	for s, value := range env {
		if value.NeedRemove {
			os.Unsetenv(s)
		}
		os.Setenv(s, value.Value)
	}

	if strings.Contains(cmd[0], "env") || strings.Contains(cmd[1], "bash") {
		return -1
	}

	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdout = os.Stdout
	err := command.Run()
	if err != nil {
		return -1
	}

	return
}
