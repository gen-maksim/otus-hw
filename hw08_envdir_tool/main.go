package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Too few arguments")
		return
	}

	envpath := args[1]
	commandArgs := args[2:]
	environments, err := ReadDir(envpath)
	if err != nil {
		println(err.Error())
		return
	}

	code := RunCmd(commandArgs, environments)
	if code != 0 {
		os.Exit(code)
	}
}
