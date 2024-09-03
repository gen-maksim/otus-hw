package main

import (
	"errors"
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

	name := cmd[0]
	if !strings.Contains(name, "bash") {
		return -1
	}

	command := exec.Command(name, cmd[1:]...)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}

		return -1
	}

	return 0
}
