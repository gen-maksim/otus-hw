package main

import (
	"bytes"
	"os"
	"strings"
)

type Environment map[string]EnvValue

type EnvValue struct {
	Value      string
	NeedRemove bool
}

func ReadDir(dir string) (Environment, error) {
	readDir, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment)

	for _, file := range readDir {
		if file.IsDir() {
			continue
		}

		value, err := readEnv(dir + "/" + file.Name())
		if err != nil {
			return nil, err
		}

		env[file.Name()] = value
	}

	return env, nil
}

func readEnv(envFilename string) (EnvValue, error) {
	file, err := os.ReadFile(envFilename)
	if err != nil {
		return EnvValue{}, err
	}

	env := EnvValue{}
	if len(file) == 0 {
		env.NeedRemove = true
	} else {
		file = bytes.Split(file, []byte("\n"))[0]
		file = bytes.ReplaceAll(file, []byte{0x00}, []byte("\n"))
		env.Value = strings.TrimRight(string(file), "\t\n\v\f\r "+string(rune(0x85))+string(rune(0xA0)))
	}

	return env, nil
}
